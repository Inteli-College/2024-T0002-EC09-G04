package gas

import (
	"math"
	"math/rand"
	"time"
)

type Payload struct {
	CO2  float64 `json:"co2"`
	CO   float64 `json:"co"`
	NO2  float64 `json:"no2"`
	MP10 float64 `json:"mp10"`
	MP25 float64 `json:"mp25"`
}

type Interval struct {
	Minimum float64
	Maximum float64
}

var Data = map[string] Interval {
	"co2":  {1, 1000},
	"co":   {0.05, 10},
	"no2":  {10, 500},
	"mp10": {1, 1000},
	"mp25": {1, 500},
}

func DataEntropy(key string) float64 {
	rand.NewSource(time.Now().UnixNano()).Int63()
	max := Data[key].Maximum
	min := Data[key].Minimum
	value := rand.Float64()*(max-min) + min
	return math.Round(value)
}

func Generate() Payload {
	data := Payload{
		CO2:  DataEntropy("co2"),
		CO:   DataEntropy("co"),
		NO2:  DataEntropy("no2"),
		MP10: DataEntropy("mp10"),
		MP25: DataEntropy("mp25"),
	}
	return data
}
