package entity

import (
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


//TODO: add test for NewSensorPayload() with invalid params

//TODO: add test for NewSensorPayload() with invalid sensor_id

//TODO: add test for NewSensorPayload() with invalid data

//TODO: add test for NewSensorPayload() testing confidence interval

//TODO: add test for NewSensorPayload() with invalid timestamp
