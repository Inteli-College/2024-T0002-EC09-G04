package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/repository"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/usecase"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mutex sync.Mutex

func main() {
	options := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s",
			os.Getenv("MONGODB_ATLAS_USERNAME"),
			os.Getenv("MONGODB_ATLAS_PASSWORD"),
			os.Getenv("MONGODB_ATLAS_CLUSTER_HOSTNAME"),
			os.Getenv("MONGODB_ATLAS_APP_NAME")))
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewSensorRepositoryMongo(client, "mongodb", "sensors")
	findAllSensorsUseCase := usecase.NewFindAllSensorsUseCase(repository)

	knownSensors := make(map[string]bool)

	var wg sync.WaitGroup

	sensors, err := findAllSensorsUseCase.Execute()
	if err != nil {
		log.Fatalf("Failed to find all sensors: %v", err)
	}

	for _, sensor := range sensors {
		knownSensors[sensor.ID] = true
		wg.Add(1)
		log.Printf("Starting initial sensor simulation: %v", sensor)
		go simulateSensor(sensor, &wg)
	}

	wg.Add(1)
	go checkForNewSensors(findAllSensorsUseCase, knownSensors, &wg)

	wg.Wait()
}

func checkForNewSensors(useCase *usecase.FindAllSensorsUseCase, knownSensors map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		sensors, err := useCase.Execute()
		if err != nil {
			log.Printf("Failed to find all sensors: %v", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		mutex.Lock()
		for _, sensor := range sensors {
			if _, exists := knownSensors[sensor.ID]; !exists {
				wg.Add(1)
				go simulateSensor(sensor, wg)
				knownSensors[sensor.ID] = true
			}
		}
		mutex.Unlock()

		time.Sleep(10 * time.Minute)
	}
}

func simulateSensor(sensor usecase.FindAllSensorsOutputDTO, wg *sync.WaitGroup) {
	defer wg.Done()

	opts := MQTT.NewClientOptions().AddBroker(
		fmt.Sprintf("ssl://%s:%s",
			os.Getenv("BROKER_TLS_URL"),
			os.Getenv("BROKER_PORT"))).SetUsername(
		os.Getenv("BROKER_USERNAME")).SetPassword(
		os.Getenv("BROKER_PASSWORD")).SetClientID(sensor.ID)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker for sensor %s: %v", sensor.ID, token.Error())
	}

	for {
		payload, err := entity.NewSensorPayload(
			sensor.ID,
			sensor.Params,
			time.Now(),
		)
		if err != nil {
			log.Fatalf("Failed to create payload for sensor %s: %v", sensor.ID, err)
		}

		jsonBytesPayload, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Error converting payload to JSON for sensor %s: %v", sensor.ID, err)
			continue
		}

		token := client.Publish("sensors", 1, false, string(jsonBytesPayload))
		log.Printf("Published for sensor %s, payload: %s", sensor.ID, string(jsonBytesPayload))
		token.Wait()

		time.Sleep(120 * time.Second)
	}
}