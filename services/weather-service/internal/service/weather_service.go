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

	locationPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location"
)

type WeatherService struct {
	weatherRepository *repository.WeatherRepository
	kafkaProducer     sarama.SyncProducer
	kafkaConsumer     sarama.ConsumerGroup
	locationClient    locationPb.LocationServiceClient
	log               *logrus.Logger
	cfg               *config.Config
}

func NewWeatherService(weatherRepo *repository.WeatherRepository, kafkaProducer sarama.SyncProducer, kafkaConsumer sarama.ConsumerGroup, locationClient locationPb.LocationServiceClient, log *logrus.Logger, cfg *config.Config) *WeatherService {
	return &WeatherService{
		weatherRepository: weatherRepo,
		kafkaConsumer:     kafkaConsumer,
		locationClient:    locationClient,
		log:               log,
		cfg:               cfg,
	}
}

func (w *WeatherService) GetWeatherById(locationId int32) (*models.WeatherData, error) {
	location, err := w.weatherRepository.GetWeatherById(locationId)
	if err != nil {
		w.log.Warnf("Ошибка получения данных по айди %v", err)
		return nil, err
	}
	w.log.Infof("получены данные из бд")
	return location, nil
}

func (w *WeatherService) StartKafkaConsumer(topik string) {
	w.log.Printf("Kafka Consumer listen to topic %s", topik)

	for {
		err := w.kafkaConsumer.Consume(context.Background(), []string{topik}, w)
		if err != nil {
			w.log.Printf("error reading message %v", err)
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
			LocationId  int32
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
		w.log.Println("LocationId: %d", location.LocationId)

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

func (w *WeatherService) SendingRequest(locationId int32, coordinates string) (*models.WeatherData, error) {

	url := fmt.Sprintf("%s?key=%s&q=%s", w.cfg.ExternalApi.BaseUrl, w.cfg.ExternalApi.ApiKey, coordinates)

	resp, err := http.Get(url)
	if err != nil {
		w.log.Warnf("error sending request %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		w.log.Warnf("error reading response body %v", err)
		return nil, err
	}

	var WeatherInformation models.WeatherInformation
	if err := json.Unmarshal(buf, &WeatherInformation); err != nil {
		w.log.Infof("Received JSON: %s", string(buf))
		w.log.Warnf("error decoding JSON %v", err)
		return nil, err
	}

	weatherData := WeatherInformation.ToWeatherData(locationId)
	return weatherData, nil
}

func (w *WeatherService) AddCurrentWeather(locationId int32, name, coordinates string) error {
	weatherData, err := w.SendingRequest(locationId, coordinates)
	if err != nil {
		w.log.Warnf("something is wrong with sending the request %v", err)
		return err
	}

	if err := w.weatherRepository.AddForcefullyWeather(weatherData); err != nil {
		w.log.Warnf("error when updating weather information from Id %d: %v", locationId, err)
		return err
	}

	return nil
}

func (w *WeatherService) UpdateCurrentWeather(locationId int32, name, coordinates string) (*models.WeatherData, error) {
	weatherData, err := w.SendingRequest(locationId, coordinates)
	if err != nil {
		w.log.Warnf("something is wrong with sending the request %v", err)
		return nil, err
	}

	if err := w.weatherRepository.UpdateCurrentWeather(weatherData); err != nil {
		w.log.Warnf("error in updating current weather data in the database %v", err)
		return nil, err
	}

	return weatherData, nil
}

func (w *WeatherService) UpdatesRegularWeather(topic string) {
	for {
		request := &locationPb.Empty{}
		locations, err := w.locationClient.GetLocations(context.Background(), request)
		if err != nil {
			w.log.Warnf("error getting locations from location-service %v", err)
		}

		for _, location := range locations.GetLocations() {
			weatherData, err := w.UpdateCurrentWeather(location.GetLocationId(), location.GetName(), location.GetCoordinates())
			if err != nil {
				w.log.Warnf("error updating current weather %v", err)
				continue
			}

			if err := w.SendMessageToKafka(weatherData, topic); err != nil {
				w.log.Warnf("failed to send message %v", err)
				continue
			}

		}
		time.Sleep(5 * time.Minute)
	}
}

func (w *WeatherService) SendMessageToKafka(weatherData *models.WeatherData, topic string) error {
	message, err := json.Marshal(weatherData)
	if err != nil {
		w.log.Warnf("JSON serialization error %v", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err = w.kafkaProducer.SendMessage(msg)
	if err != nil {
		w.log.Warnf("error sending message to kafka %v", err)
		return err
	}

	return nil
}
