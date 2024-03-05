package repository

import (
	"log"
	"database/sql"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	_ "github.com/lib/pq"
)

type SensorRepositoryPostgres struct {
	DB *sql.DB
}

func NewSensorRepositoryPostgres(db *sql.DB) *SensorRepositoryPostgres {
	return &SensorRepositoryPostgres{
		DB: db,
	}
}

func (s *SensorRepositoryPostgres) CreateSensor(sensor *entity.Sensor) error {
	_, err := s.DB.Exec("INSERT INTO sensors (id, name, latitude, longitude) VALUES ($1, $2, $3, $4)", sensor.ID, sensor.Name, sensor.Latitude, sensor.Longitude)
	log.Printf("Inserting sensor with ID: %s, Name: %s, Latitude: %f, Longitude: %f into the database", sensor.ID, sensor.Name, sensor.Latitude, sensor.Longitude)
	if err != nil {
		return err
	}
	return nil
}

func (s *SensorRepositoryPostgres) CreateSensorLog(sensorLog *entity.Log) error {
	_, err := s.DB.Exec("INSERT INTO sensors_log (sensor_id, data, timestamp) VALUES ($1, $2, $3)", sensorLog.ID, sensorLog.Data, sensorLog.Timestamp)
	log.Printf("Inserting log  with ID: %s, Data: %s, Timestamp: %s into the database", sensorLog.ID, sensorLog.Data, sensorLog.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (s *SensorRepositoryPostgres) FindAllSensors() ([]*entity.Sensor, error) {
	rows, err := s.DB.Query("SELECT id, name, latitude, longitude FROM sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sensors []*entity.Sensor
	for rows.Next() {
		var sensor entity.Sensor
		if err := rows.Scan(&sensor.ID, &sensor.Name, &sensor.Latitude, &sensor.Longitude); err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensor)
	}
	return sensors, nil
}