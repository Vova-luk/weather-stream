package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/Vova-luk/weather-stream/services/weather-service/internal/config"
	"github.com/Vova-luk/weather-stream/services/weather-service/internal/models"
	"github.com/Vova-luk/weather-stream/services/weather-service/internal/repository"
	"github.com/sirupsen/logrus"
)

type WeatherService struct {
	weatherRepository *repository.WeatherRepository
	kafkaConsumer     sarama.ConsumerGroup
	log               *logrus.Logger
	cfg               *config.Config
}

func NewWeatherService(weatherRepo *repository.WeatherRepository, kafkaConsumer sarama.ConsumerGroup, log *logrus.Logger, cfg *config.Config) *WeatherService {
	return &WeatherService{
		weatherRepository: weatherRepo,
		kafkaConsumer:     kafkaConsumer,
		log:               log,
		cfg:               cfg,
	}
}

func (w *WeatherService) GetWeatherById(locationId int32) (*models.WeatherData, error) {
	location, err := w.weatherRepository.GetWeatherById(locationId)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (w *WeatherService) StartKafkaConsumer(topik string) {
	w.log.Printf("Kafka Consumer listen to topic %s", topik)

	for {
		err := w.kafkaConsumer.Consume(context.Background(), []string{topik}, w)
		if err != nil {
			w.log.Printf("error reading message %v", err)
			time.Sleep(5 * time.Second)
		}
	}
}

func (w *WeatherService) Cleanup(sarama.ConsumerGroupSession) error {
	w.log.Println("Kafka Consumer Group is shutting down")
	return nil
}

func (w *WeatherService) Setup(sarama.ConsumerGroupSession) error {
	w.log.Println("Kafka Consumer Group launched")
	return nil
}

func (w *WeatherService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var location struct {
			LocationId  int
			Name        string
			Coordinates string
		}

		err := json.Unmarshal(msg.Value, &location)
		if err != nil {
			w.log.Printf("error decoding JSON %v", err)
			session.MarkMessage(msg, "")
			continue
		}

		w.log.Printf("Message received from Kafka: %+v", err)

		if err := w.AddCurrentWeather(location.LocationId, location.Name, location.Coordinates); err != nil {
			w.log.Errorf("error when updating weather %v", err)
			session.MarkMessage(msg, "")
			continue
		}

		session.MarkMessage(msg, "")
		w.log.Printf("weather updated for %v", err)
	}
	return nil
}

func (w *WeatherService) AddCurrentWeather(locationId int, name, coordinates string) error {

	url := fmt.Sprintf("%s?key=%s&q=%s", w.cfg.ExternalApi.BaseUrl, w.cfg.ExternalApi.ApiKey, coordinates)

	resp, err := http.Get(url)
	if err != nil {
		w.log.Warnf("error sending request %v", err)
		return err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		w.log.Warnf("error reading response body %v", err)
		return err
	}

	var WeatherInformation models.WeatherInformation
	if err := json.Unmarshal(buf, &WeatherInformation); err != nil {
		w.log.Infof("Received JSON: %s", string(buf))
		w.log.Warnf("error decoding JSON %v", err)
		return err
	}

	weatherData := WeatherInformation.ToWeatherData(locationId)

	if err := w.weatherRepository.AddForcefullyWeather(weatherData); err != nil {
		w.log.Warnf("error when updating weather information from Id %d: %v", locationId, err)
		return err
	}

	return nil
}
