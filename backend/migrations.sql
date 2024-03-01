CREATE TABLE Estacao (
    id SERIAL PRIMARY KEY,
    latitude FLOAT,
    longitude FLOAT
);

CREATE TABLE Gas (
    id_gas INT NOT NULL IDENTITY PRIMARY KEY,
    id_estacao INT,
    CO2 FLOAT,
    CO FLOAT,
    NO2 FLOAT,
    MP10 FLOAT,
    MP25 FLOAT,
    timestamp TIMESTAMP,
    FOREIGN KEY (id_estacao) REFERENCES Estacao(id)
);

CREATE TABLE Rad_lum (
    id_rad SERIAL PRIMARY KEY,
    id_estacao INT,
    ET FLOAT,
    LI FLOAT,
    SR FLOAT,
    timestamp TIMESTAMP,
    FOREIGN KEY (id_estacao) REFERENCES Estacao(id)
);

CREATE TABLE alertas (
    id SERIAL PRIMARY KEY,
    latitude FLOAT,
    longitude FLOAT,
    timestamp TIMESTAMP,
    opcoes TEXT,
);