CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    coordinates VARCHAR(50) NOT NULL 
);

CREATE TABLE weather (
    location_id INT PRIMARY KEY,
    temperature DECIMAL(5, 2) NOT NULL, 
    humidity INT NOT NULL,
    wind_speed DECIMAL(5, 2) NOT NULL,
    pressure INT NOT NULL,
    condition VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW()
);