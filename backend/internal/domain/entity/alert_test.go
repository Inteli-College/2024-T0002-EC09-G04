package entity

import (
	"testing"
)

func TestNewAlert(t *testing.T) {
	alert := NewAlert(0, 0, "")
	if alert.Latitude != 0 {
		t.Errorf("Latitude should be 0")
	}
	if alert.Longitude != 0 {
		t.Errorf("Longitude should be 0")
	}
	if alert.Option != "" {
		t.Errorf("Option should be empty")
	}
}
