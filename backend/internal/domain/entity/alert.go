package entity

import (
	"github.com/google/uuid"
	"time"
)

type AlertRepository interface {
	CreateAlert(data *Alert) error
	FindAllAlerts() ([]*Alert, error)
}

type Alert struct {
	ID        string    `json:"id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Option    string    `json:"option"`
	Timestamp time.Time `json:"timestamp"`
}

func NewAlert(latitude float64, longitude float64, option string) *Alert {
	return &Alert{ID: uuid.New().String(), Latitude: latitude, Longitude: longitude, Option: option, Timestamp: time.Now()}
}
