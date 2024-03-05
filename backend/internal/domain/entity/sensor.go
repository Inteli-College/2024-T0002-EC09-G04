package entity

import (
	"encoding/json"
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
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func Entropy(newInterval []float64) float64 {
	rand.NewSource(time.Now().UnixNano())
	return math.Round(rand.Float64()*(newInterval[0]-newInterval[1]) + newInterval[1])
}

func NewSensor(name string, latitude float64, longitude float64) *Sensor {
	return &Sensor{ID: uuid.New().String(), Name: name, Latitude: latitude, Longitude: longitude}
}

func NewSensorPayload(id string, props map[string][]float64) (string, string) {
	entropyValues := make(map[string]float64)
	for key, interval := range props {
		intervalValues := make([]float64, int(interval[1]-interval[0])+1)
		for i := range intervalValues {
			intervalValues[i] = interval[0] + float64(i)
		}
		mean, stdDev := stat.MeanStdDev(intervalValues, nil)
		factor := stdDev / math.Sqrt(float64(len(intervalValues)))
		confidenceInterval := []float64{mean - interval[2]*factor, mean + interval[2]*factor}
		entropyValues[key] = Entropy(confidenceInterval)
	}
	bytes, err := json.Marshal(entropyValues)
	if err != nil {
		return id, ""
	}
	return id, string(bytes)
}
