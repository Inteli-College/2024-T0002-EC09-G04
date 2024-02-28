package station

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/gas"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/rad_lum"
	"math"
	"math/rand"
	"time"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Station struct {
	Location  string `json:"location"`
	Gas       string `json:"gas"`
	RadLum    string `json:"rad_lum"`
	TimeStamp string `json:"timestamp"`
}

type Interval struct {
	Minimum float64
	Maximum float64
}

var Area = map[string]Interval{
	"latitude":  {0, 39},
	"longitude": {0, 39},
}

func LocationEntropy(key string) float64 {
	rand.NewSource(time.Now().UnixNano())
	max := Area[key].Maximum
	min := Area[key].Minimum
	value := rand.Float64()*(max-min) + min
	return math.Round(value)
}

func GenerateLocation() Location {
	data := Location{
		Latitude:  LocationEntropy("latitude"),
		Longitude: LocationEntropy("longitude"),
	}
	return data
}

func ConnectMQTT(seed rand.Source, url string) MQTT.Client {
	opts := MQTT.NewClientOptions().AddBroker(url)
	opts.SetClientID(fmt.Sprintf("station-%d", seed.Int63()))
	client := MQTT.NewClient(opts)
	if session := client.Connect(); session.Wait() && session.Error() != nil {
		panic(session.Error())
	}
	return client
}

func Start(url string) {
	client := ConnectMQTT(rand.NewSource(time.Now().UnixNano()), url)
	location := GenerateLocation()
	gasData := gas.Generate()
	radLumData := rad_lum.Generate()

	data := Station{
		Location: fmt.Sprintf(`{"latitude":%f,"longitude":%f}`, location.Latitude, location.Longitude),
		Gas: fmt.Sprintf(`{"CO2":%f,"CO":%f,"NO2":%f,"MP10":%f,"MP25":%f}`, gasData.CO2, gasData.CO, gasData.NO2, gasData.MP10, gasData.MP25),
		RadLum: fmt.Sprintf(`{"ET":%f,"LI":%f,"SR":%f}`, radLumData.ET, radLumData.LI, radLumData.SR),
		TimeStamp: time.Now().String(),
	}

	for {
		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			return
		}
		token := client.Publish("/stations", 0, false, string(payload))
		token.Wait()
		time.Sleep(5 * time.Second)
	}
}
