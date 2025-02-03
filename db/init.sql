CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    coordinates VARCHAR(50) NOT NULL 
);

CREATE TABLE weathers (
    location_id INT PRIMARY KEY,
    temperature DECIMAL(5, 2) NOT NULL, 
    humidity INT NOT NULL,
    wind_speed DECIMAL(5, 2) NOT NULL,
    pressure DECIMAL(5, 2) NOT NULL,
    precip DECIMAL(5, 2) NOT NULL,
    cloud INT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW()
);