package handler

import (
	"fleetpulse/ingestion/internal/model"
	"fleetpulse/ingestion/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IngestData(c *gin.Context, speedService *service.SpeedService) {
	var data model.GPSData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri formatı"})
		return
	}

	speedService.ProcessGPSData(data)

	c.JSON(http.StatusOK, gin.H{"status": "Veri işleniyor", "truckId": data.TruckID})
}