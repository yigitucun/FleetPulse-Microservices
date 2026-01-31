package main

import (
	"fleetpulse/ingestion/internal/client"
	"fleetpulse/ingestion/internal/discovery"
	"fleetpulse/ingestion/internal/handler"
	"fleetpulse/ingestion/internal/messaging"
	"fleetpulse/ingestion/internal/repository"
	"fleetpulse/ingestion/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	redisURL := getEnv("REDIS_URL", "localhost:6379")
	rabbitURL := getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	fleetServiceURL := getEnv("FLEET_SERVICE_URL", "http://localhost:8083")
	eurekaURL := getEnv("EUREKA_URL", "http://localhost:8761/eureka")

	redisRepo := repository.NewRedisRepository(redisURL)
	rabbitPub := messaging.NewRabbitMQPublisher(rabbitURL, "speed_alerts")
	fleetClient := client.NewFleetClient(fleetServiceURL)
	speedService := service.NewSpeedService(redisRepo, rabbitPub, fleetClient)

	discovery.Register(eurekaURL)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/ingest", func(c *gin.Context) {
			handler.IngestData(c, speedService)
		})
	}

	r.GET("/health", func(c *gin.Context) { c.Status(200) })

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}