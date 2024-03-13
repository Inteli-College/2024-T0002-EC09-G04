package entity

import (
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlertRepository interface {
	CreateAlert(data *Alert) (*mongo.InsertOneResult, error)
	FindAllAlerts() ([]*Alert, error)
}

type Alert struct {
	ID        string    `json:"_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Option    string    `json:"option"`
	Timestamp time.Time `json:"timestamp"`
}

func NewAlert(latitude float64, longitude float64, option string) *Alert {
	return &Alert{Latitude: latitude, Longitude: longitude, Option: option, Timestamp: time.Now()}
}
