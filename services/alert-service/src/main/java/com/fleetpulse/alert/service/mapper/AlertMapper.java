package com.fleetpulse.alert.service.mapper;

import com.fleetpulse.alert.service.dto.requests.Alert;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Options;

@Mapper
public interface AlertMapper {
    @Insert("INSERT INTO alerts (truck_id, speed, lat, lng, timestamp) " +
            "VALUES (#{truckId}, #{speed}, #{lat}, #{lng}, #{timestamp})")
    @Options(useGeneratedKeys = true, keyProperty = "id")
    void insert(Alert alert);
}