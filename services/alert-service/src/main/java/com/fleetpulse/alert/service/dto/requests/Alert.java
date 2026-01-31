package com.fleetpulse.alert.service.dto.requests;

import java.time.LocalDateTime;

public class Alert {
    private Long id;
    private String truckId;
    private Double speed;
    private Double lat;
    private Double lng;
    private LocalDateTime timestamp;

    public Alert(Long id, String truckId, Double lat, Double lng, LocalDateTime timestamp, Double speed) {
        this.id = id;
        this.truckId = truckId;
        this.lat = lat;
        this.lng = lng;
        this.timestamp = timestamp;
        this.speed = speed;
    }

    public Alert() {
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getTruckId() {
        return truckId;
    }

    public void setTruckId(String truckId) {
        this.truckId = truckId;
    }

    public Double getSpeed() {
        return speed;
    }

    public void setSpeed(Double speed) {
        this.speed = speed;
    }

    public Double getLat() {
        return lat;
    }

    public void setLat(Double lat) {
        this.lat = lat;
    }

    public Double getLng() {
        return lng;
    }

    public void setLng(Double lng) {
        this.lng = lng;
    }

    public LocalDateTime getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(LocalDateTime timestamp) {
        this.timestamp = timestamp;
    }
}
