package com.fleetpulse.fleet_management.service.mapper;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;

@Mapper
public interface TruckMapper {
    @Select("SELECT COUNT(*) > 0 FROM trucks WHERE truck_id = #{truckId} AND status = 'ACTIVE'")
    boolean isTruckActive(String truckId);
}