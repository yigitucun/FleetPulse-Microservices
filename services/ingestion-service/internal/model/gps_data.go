package model

type GPSData struct {
	TruckID   string  `json:"truck_id" binding:"required"`
	Lat       float64 `json:"lat" binding:"required"`
	Lng       float64 `json:"lng" binding:"required"`
	Speed     float64 `json:"speed" binding:"required"`
	Timestamp int64   `json:"timestamp" binding:"required"`
}
