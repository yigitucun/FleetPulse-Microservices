package service

import (
	"fleetpulse/ingestion/internal/client" // Yeni ekledik
	"fleetpulse/ingestion/internal/messaging"
	"fleetpulse/ingestion/internal/model"
	"fleetpulse/ingestion/internal/repository"
	"log"
)

type SpeedService struct {
	repo        *repository.RedisRepository
	publisher   *messaging.RabbitMQPublisher
	fleetClient *client.FleetClient // Yeni ekledik
}

func NewSpeedService(repo *repository.RedisRepository, pub *messaging.RabbitMQPublisher, fleet *client.FleetClient) *SpeedService {
	return &SpeedService{repo: repo, publisher: pub, fleetClient: fleet}
}

func (s *SpeedService) ProcessGPSData(data model.GPSData) {
	isValid, err := s.fleetClient.ValidateTruck(data.TruckID)
	if err != nil || !isValid {
		log.Printf("Kayıtsız veya Pasif Tır -> %s", data.TruckID)
		return
	}

	if data.Speed > 120 {
		log.Printf("Hız Aşımı %s", data.TruckID)
	}

}