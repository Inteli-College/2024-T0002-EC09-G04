package entity

import (
	"testing"
)

func TestEntropy(t *testing.T) {
	entropy := Entropy([]float64{0, 100})
	if entropy <= 0 && entropy >= 100 {
		t.Errorf("Entropy should be between 0 and 100")
	}
}

func TestNewSensor(t *testing.T) {
	sensor := NewSensor("name", 0, 0)
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

//TODO: add test for NewSensorPayload()
