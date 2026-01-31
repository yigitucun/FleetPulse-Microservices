package com.fleetpulse.alert.service.consumer;

import com.fleetpulse.alert.service.dto.requests.Alert;
import com.fleetpulse.alert.service.dto.requests.GPSDataDTO;
import com.fleetpulse.alert.service.mapper.AlertMapper;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
public class AlertConsumer {

    private final AlertMapper alertMapper;

    public AlertConsumer(AlertMapper alertMapper) {
        this.alertMapper = alertMapper;
    }

    @RabbitListener(queues = "speed_alerts")
    public void receiveAlert(GPSDataDTO data) {
        Alert alert = new Alert();
        alert.setTruckId(data.getTruckId());
        alert.setSpeed(data.getSpeed());
        alert.setLat(data.getLat());
        alert.setLng(data.getLng());
        alert.setTimestamp(LocalDateTime.now());
        System.out.println("çalıştı");
        alertMapper.insert(alert);
    }
}