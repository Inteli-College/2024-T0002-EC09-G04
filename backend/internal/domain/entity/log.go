package entity

import (
	"time"
)

type LogRepository interface {
	Publish(data *Log) error
}

type Log struct {
	Sensor_ID string                 `json:"sensor_id"`
	Data      map[string]interface{} `json:"data"`
	Timestamp time.Time              `json:"timestamp"`
}

func NewLog(id string, data map[string]interface{}, timestamp time.Time) *Log {
	return &Log{Sensor_ID: id, Data: data, Timestamp: timestamp}
}
