DROP TABLE IF EXISTS alerts;

CREATE TABLE alerts (
                        id SERIAL PRIMARY KEY,
                        truck_id VARCHAR(50) NOT NULL,
                        speed DOUBLE PRECISION NOT NULL,
                        lat DOUBLE PRECISION,
                        lng DOUBLE PRECISION,
                        timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);