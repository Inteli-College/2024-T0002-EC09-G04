CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE options AS ENUM('Flood', 'Landslide', 'Fire', 'Construction', 'Accident');

CREATE TABLE sensors (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT,
    latitude FLOAT8,
    longitude FLOAT8
);

INSERT INTO sensors (name, latitude, longitude) VALUES ('MICS-6814', -23.5718, -46.708);
INSERT INTO sensors (name, latitude, longitude) VALUES ('MICS-6814', -23.5718, -46.708);
INSERT INTO sensors (name, latitude, longitude) VALUES ('SPS30', -23.5718, -46.708);
INSERT INTO sensors (name, latitude, longitude) VALUES ('RXW-LIB-900', -23.5718, -46.708);
INSERT INTO sensors (name, latitude, longitude) VALUES ('RXW-LIB-900', -23.5718, -46.708);

CREATE TABLE sensors_log (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    sensor_id UUID REFERENCES sensors(id),
    data TEXT,
    timestamp TIMESTAMP
);

CREATE TABLE alerts (
    id SERIAL PRIMARY KEY,
    latitude FLOAT8,
    longitude FLOAT8,
    option OPTIONS,
    timestamp TIMESTAMP
);