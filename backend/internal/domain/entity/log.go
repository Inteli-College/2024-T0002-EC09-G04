package entity

import "time"

type LogRepository interface {
	Transmit(data *Log) error
}

type Log struct {
	ID        string    `json:"id"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func NewLog(id string, data string) *Log {
	return &Log{ID: id, Data: data, Timestamp: time.Now()}
}
