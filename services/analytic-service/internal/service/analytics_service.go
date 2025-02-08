package service

import (
	"context"
	"errors"

	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/models"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/repository"
	"github.com/sirupsen/logrus"
)

type AnalyticsService struct {
	analyticsRepository *repository.AnalyticsRepository
	kafkaConsumerGroup  sarama.ConsumerGroup
	log                 *logrus.Logger
}

func NewAnalyticsService(analyticsRepository *repository.AnalyticsRepository, kafkaConsumerGroup sarama.ConsumerGroup, log *logrus.Logger) *AnalyticsService {
	return &AnalyticsService{
		analyticsRepository: analyticsRepository,
		kafkaConsumerGroup:  kafkaConsumerGroup,
		log:                 log,
	}
}

func (a *AnalyticsService) GetAnalyticsById(locationId int32, period string) (*models.WeatherAnalytic, error) {
	var interval string
	switch period {
	case "day":
		interval = "1 day"
	case "week":
		interval = "1 week"
	case "month":
		interval = "1 month"
	default:
		return nil, errors.New("invalid period: " + period)
	}

	weatherAnalytic, err := a.analyticsRepository.GetAnalyticsById(locationId, interval)
	if err != nil {
		a.log.Warnf("error receiving analytics %v", err)
		return nil, err
	}

	return weatherAnalytic, nil

}

func (a *AnalyticsService) StartKafkaConsumer(topic string) {
	a.log.Printf("Kafka Consumer listen to topic %s", topic)

	for {
		if err := a.kafkaConsumerGroup.Consume(context.Background(), []string{topic}, a); err != nil {
			a.log.Warnf("error reading message %v", err)
		}
	}
}

func (a *AnalyticsService) Cleanup(sarama.ConsumerGroupSession) error {
	a.log.Println("Kafka Consumer Group is shutting down")
	return nil
}

func (a *AnalyticsService) Setup(sarama.ConsumerGroupSession) error {
	a.log.Println("Kafka Consumer Group launched")
	return nil
}

func (a *AnalyticsService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var weatherData models.WeatherData

		if err := json.Unmarshal(msg.Value, &weatherData); err != nil {
			a.log.Warnf("error decoding JSON %v", err)
			session.MarkMessage(msg, "")
			continue
		}

		if err := a.analyticsRepository.AddWeatherData(&weatherData); err != nil {
			a.log.Printf("error adding weather data %v", err)
			session.MarkMessage(msg, "")
			continue
		}
		session.MarkMessage(msg, "")

	}
	return nil
}
