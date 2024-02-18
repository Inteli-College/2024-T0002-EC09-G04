package rad_lum

import (
	"math"
	"math/rand"
	"time"
)
type Payload struct {
	ET  float64 `json:"et"`
	LI   float64 `json:"li"`
	SR  float64 `json:"sr"`
}

type Interval struct {
	Minimum float64
	Maximum float64
}

var Values = map[string] Interval {
	"et":  {1, 1000}, //change values
	"li":   {0.05, 10}, //change values
	"sr":  {10, 500}, //change values
}

func DataEntropy(key string) float64 {
	rand.NewSource(time.Now().UnixNano()).Int63()
	max := Values[key].Maximum
	min := Values[key].Minimum
	value := rand.Float64()*(max-min) + min
	return math.Round(value)
}

func Generate() Payload {
	data := Payload{
		ET:  DataEntropy("et"),
		LI:   DataEntropy("li"),
		SR:  DataEntropy("sr"),
	}
	return data
}
