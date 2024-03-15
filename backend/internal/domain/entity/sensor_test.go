package entity

import (
	"math"
	"testing"
	"time"
)

func TestEntropy(t *testing.T) {
	entropy := Entropy([]float64{0, 100})
	if entropy <= 0 && entropy >= 100 {
		t.Errorf("Entropy should be between 0 and 100")
	}
}

func TestNewSensor(t *testing.T) {
	sensor := NewSensor("name", 0, 0, map[string]Param{"key": {Min: 0, Max: 100, Factor: 0.5}})
	if sensor.Name != "name" {
		t.Errorf("Name should be name")
	}
	if sensor.Latitude != 0 {
		t.Errorf("Latitude should be 0")
	}
	if sensor.Longitude != 0 {
		t.Errorf("Longitude should be 0")
	}
}

func TestNewSensorPayload(t *testing.T) {
	sensorPayload, _ := NewSensorPayload("id", map[string]Param{"key": {Min: 0, Max: 100, Factor: 0.5}}, time.Now())
	if sensorPayload.Sensor_ID != "id" {
		t.Errorf("Sensor_ID should be id")
	}
	if value, ok := sensorPayload.Data["key"].(float64); ok {
			if !(value <= 180 && value >= 0) {
					t.Errorf("Invalid value for Data['key'], expected outside the range %v and %v, got %v", 0, 100, value)
			}
	} else {
			t.Errorf("Invalid type for Data['key']")
	}
}

func TestNewSensorPayloadParams(t *testing.T) {
	_, err := NewSensorPayload("id", map[string]Param{"key": {Min: 0, Max: 100, Factor: 0.5}}, time.Now())
	if err != nil {
		t.Errorf("Error should be nil")
	}
}

func TestNewSensorPayloadTimestamp(t *testing.T) {
	sensorPayload, _ := NewSensorPayload("id", map[string]Param{"key": {Min: 0, Max: 100, Factor: 0.5}}, time.Now())
	if sensorPayload.Timestamp.IsZero() {
		t.Errorf("Timestamp should not be zero")
	}

	if sensorPayload.Timestamp.After(time.Now()) {
		t.Errorf("Timestamp should be before now")
	}

	if sensorPayload.Timestamp.Before(time.Now().Add(-time.Minute * 10)) {
		t.Errorf("Timestamp should be within the last 10 minutes")
	}

}

func TestNewSensorInvalidId(t *testing.T) {
	_, err := NewSensorPayload("", map[string]Param{"key": {Min: 0, Max: 100, Factor: 0.5}}, time.Now())
	if err == nil {
		t.Errorf("Error should not be nil")
	}
}

func TestNewSensorPayloadConfidenceInterval(t *testing.T) {
	params := map[string]Param{
		"key": {Min: 10, Max: 30, Factor: 1.5},
	}
	timestamp := time.Now()

	payload, err := NewSensorPayload("sensorID", params, timestamp)
	if err != nil {
		t.Fatalf("Failed to generate sensor payload: %v", err)
	}

	for key, param := range params {
		value, ok := payload.Data[key].(float64)
		if !ok {
			t.Fatalf("Generated value for %s is not a float64", key)
		}

		numValues := float64(param.Max - param.Min + 1)
		mean := float64(param.Min) + (numValues-1)/2
		stdDev := math.Sqrt(numValues*numValues - 1) / 12 
		confidenceFactor := stdDev / math.Sqrt(numValues)
		lowerBound := mean - param.Factor*confidenceFactor
		upperBound := mean + param.Factor*confidenceFactor

		if value < lowerBound || value > upperBound {
			t.Errorf("Value for %s (%v) is outside the confidence interval [%v, %v]", key, value, lowerBound, upperBound)
		}
	}
}


//TODO: add test for NewSensorPayload() with invalid params

//TODO: add test for NewSensorPayload() with invalid sensor_id

//TODO: add test for NewSensorPayload() with invalid data

//TODO: add test for NewSensorPayload() testing confidence interval

//TODO: add test for NewSensorPayload() with invalid timestamp
