package com.fleetpulse.fleet_management.service.controller;

import com.fleetpulse.fleet_management.service.mapper.TruckMapper;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/fleet")
public class FleetController {
    private final TruckMapper truckMapper;

    public FleetController(TruckMapper truckMapper) {
        this.truckMapper = truckMapper;
    }
    @GetMapping("/validate/{truckId}")
    public ResponseEntity<Boolean> validateTruck(@PathVariable String truckId) {
        boolean isValid = truckMapper.isTruckActive(truckId);
        return ResponseEntity.ok(isValid);
    }
}
