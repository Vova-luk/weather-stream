package service

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/config"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/models"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/repository"
	locationPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location"
	"github.com/sirupsen/logrus"
)

type LocationService struct {
	locationRepository *repository.LocationRepository
	kafkaProducer      sarama.SyncProducer
	log                *logrus.Logger
	cfg                *config.Config
}

func NewLocationService(locationRepository *repository.LocationRepository, kafkaProducer sarama.SyncProducer, log *logrus.Logger, cfg *config.Config) *LocationService {
	return &LocationService{
		locationRepository: locationRepository,
		kafkaProducer:      kafkaProducer,
		log:                log,
		cfg:                cfg}
}

func (l *LocationService) CreateLocationService(location *models.Location) (int, error) {
	locationId, err := l.locationRepository.CreateLocation(location)
	if err != nil {
		l.log.Errorf("error when creating location: %v", err.Error())
		return 0, err
	}
	go func() {
		err := l.SendMessageToKafka(*location, l.cfg.Kafka.LocationTopic)
		if err != nil {
			l.log.Errorf("error when sending message to kafka: %s", err)
		} else {
			l.log.Errorf("successful sending of message to kafka with locationId: %d", locationId)
		}
	}()

	l.log.Infof("Location saved successfully, ID : %d", locationId)
	return locationId, nil
}

func (l *LocationService) GetLocationsService() ([]*locationPb.Location, error) {
	locations, err := l.locationRepository.GetLocations()
	if err != nil {
		l.log.Errorf("error when getting locations: %v", err.Error())
		return nil, err
	}
	locationsItems := make([]*locationPb.Location, len(locations))
	for ind, location := range locations {
		locationsItems[ind] = &locationPb.Location{
			LocationId:  int32(location.ID),
			Name:        location.Name,
			Coordinates: location.Coordinates,
		}
	}
	l.log.Info("Locations successfully received")
	return locationsItems, nil
}

func (l *LocationService) SendMessageToKafka(location models.Location, topic string) error {

	message, err := json.Marshal(location)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err = l.kafkaProducer.SendMessage(msg)
	return err

}
