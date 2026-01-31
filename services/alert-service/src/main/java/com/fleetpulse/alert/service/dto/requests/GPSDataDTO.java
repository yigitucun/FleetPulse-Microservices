package com.fleetpulse.alert.service.dto.requests;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.io.Serializable;

public class GPSDataDTO implements Serializable {
    @JsonProperty("truck_id")
    private String truckId;
    private Double lat;
    private Double lng;
    private Double speed;
    private Long timestamp;

    public GPSDataDTO() {
    }

    public GPSDataDTO(String truckId, Double lat, Double lng, Double speed, Long timestamp) {
        this.truckId = truckId;
        this.lat = lat;
        this.lng = lng;
        this.speed = speed;
        this.timestamp = timestamp;
    }

    public String getTruckId() {
        return truckId;
    }

    public void setTruckId(String truckId) {
        this.truckId = truckId;
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

    public Double getSpeed() {
        return speed;
    }

    public void setSpeed(Double speed) {
        this.speed = speed;
    }

    public Long getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(Long timestamp) {
        this.timestamp = timestamp;
    }
}
