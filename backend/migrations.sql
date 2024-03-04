CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE options AS ENUM('Flood', 'Landslide', 'Fire', 'Construction', 'Accident');

INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);
INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);
INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);
INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);
INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);

CREATE TABLE estacao (
    id_estacao UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT,
    latitude FLOAT8,
    longitude FLOAT8,
    timestamp TIMESTAMP
);

CREATE TABLE Gas (
    gas_pkey SERIAL,
    id_estacao UUID REFERENCES estacao(id),
    co FLOAT8,
    co2 FLOAT8,
    no2 FLOAT8,
    mp10 FLOAT8,
    mp25 FLOAT8,
    timestamp TIMESTAMP
);

CREATE TABLE Rad_Lum (
    rad_lum_pkey SERIAL,
    id_estacao UUID REFERENCES estacao(id),
    et FLOAT8, 
    li FLOAT8, 
    sr FLOAT8, 
    timestamp TIMESTAMP
)

CREATE TABLE alerts (
    alert_pkey SERIAL,
    latitude FLOAT8,
    longitude FLOAT8,
    option OPTIONS,
    timestamp TIMESTAMP
);