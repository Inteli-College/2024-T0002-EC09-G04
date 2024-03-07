package entity

import (
	"github.com/google/uuid"
	"gonum.org/v1/gonum/stat"
	"math"
	"math/rand"
	"time"
)

type SensorRepository interface {
	CreateSensor(sensor *Sensor) error
	CreateSensorLog(log *Log) error
	FindAllSensors() ([]*Sensor, error)
}

type Sensor struct {
	ID        string             `json:"_id"`
	Name      string             `json:"name"`
	Latitude  float64            `json:"latitude"`
	Longitude float64            `json:"longitude"`
	Params    map[string]Param   `json:"params"`
}

type Param struct {
	Min    int     `json:"min"`
	Max    int     `json:"max"`
	Factor float64 `json:"z"`
}

type SensorPayload struct {
	Sensor_ID string                 `json:"sensor_id"`
	Data      map[string]interface{} `json:"data"`
	Timestamp time.Time              `json:"timestamp"`
}

func Entropy(newInterval []float64) float64 {
	rand.NewSource(time.Now().UnixNano())
	return math.Round(rand.Float64()*(newInterval[0]-newInterval[1]) + newInterval[1])
}

func NewSensor(name string, latitude float64, longitude float64) *Sensor {
	return &Sensor{ID: uuid.New().String(), Name: name, Latitude: latitude, Longitude: longitude}
}

func NewSensorPayload(id string, params map[string]Param, timestamp time.Time) (*SensorPayload, error) {
	entropyValues := make(map[string]interface{})
	for key, interval := range params {
		intervalValues := make([]float64, int(interval.Max-interval.Min)+1)
		for i := range intervalValues {
			intervalValues[i] = float64(interval.Min) + float64(i)
		}
		mean, stdDev := stat.MeanStdDev(intervalValues, nil)
		factor := stdDev / math.Sqrt(float64(len(intervalValues)))
		confidenceInterval := []float64{mean - interval.Factor*factor, mean + interval.Factor*factor}
		entropyValues[key] = Entropy(confidenceInterval)
	}
	return &SensorPayload{Sensor_ID: id, Data: entropyValues, Timestamp: timestamp}, nil
}
