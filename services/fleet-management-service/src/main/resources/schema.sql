DROP TABLE IF EXISTS trucks;

CREATE TABLE  trucks (
    id SERIAL PRIMARY KEY,
    truck_id VARCHAR(50) UNIQUE NOT NULL,
    plate_number VARCHAR(20),
    model VARCHAR(50),
    status VARCHAR(20) DEFAULT 'ACTIVE'
);

INSERT INTO trucks (truck_id, plate_number, model) VALUES ('TR001', '34 ABC 123', 'Scania R500');
INSERT INTO trucks (truck_id, plate_number, model) VALUES ('TR002', '06 DEF 456', 'Volvo FH16');