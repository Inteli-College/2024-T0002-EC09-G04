package repository

import (
	"log"
	"database/sql"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	_ "github.com/lib/pq"
)

type AlertRepositoryPostgres struct {
	DB *sql.DB
}

func NewAlertRepositoryPostgres(db *sql.DB) *AlertRepositoryPostgres {
	return &AlertRepositoryPostgres{
		DB: db,
	}
}

func (a *AlertRepositoryPostgres) CreateAlert(alert *entity.Alert) error {
	_, err := a.DB.Exec("INSERT INTO alerts (latitude, longitude, option, timestamp) VALUES ($1, $2, $3, $4)", alert.Latitude, alert.Longitude, alert.Option, alert.Timestamp)
	log.Printf("Inserting alert with Latitude: %f, Longitude: %f, Option: %s, Timestamp: %s into the database", alert.Latitude, alert.Longitude, alert.Option, alert.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (a *AlertRepositoryPostgres) FindAllAlerts() ([]*entity.Alert, error) {
	rows, err := a.DB.Query("SELECT latitude, longitude, timestamp, option FROM alerts")
	log.Printf("Selecting all alerts from the database")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var alerts []*entity.Alert
	for rows.Next() {
		var alert entity.Alert
		if err := rows.Scan(&alert.Latitude, &alert.Longitude, &alert.Timestamp, &alert.Option); err != nil {
			return nil, err
		}
		alerts = append(alerts, &alert)
	}
	return alerts, nil
}
