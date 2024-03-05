package main

import (
	"database/sql"
	"fmt"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/mqtt"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/repository"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/usecase"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"sync"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		log.Fatalf("Failed to connect to Database: %v", err)
	}
	defer db.Close()

	repository := repository.NewSensorRepositoryPostgres(db)
	findAllSensorsUseCase := usecase.NewFindAllSensorsUseCase(repository)

	sensors, err := findAllSensorsUseCase.Execute()
	if err != nil {
		log.Fatalf("Failed to find all sensors: %v", err)
	}

	var wg sync.WaitGroup
	for _, sensor := range sensors {
		wg.Add(1)
		go func(sensor usecase.FindAllSensorsOutputDTO) {
			defer wg.Done()
			opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("ssl://%s:8883", os.Getenv("BROKER_TLS_URL"))).SetUsername(os.Getenv("BROKER_USERNAME")).SetPassword(os.Getenv("BROKER_PASSWORD")).SetClientID(sensor.ID)
			client := MQTT.NewClient(opts)
			if session := client.Connect(); session.Wait() && session.Error() != nil {
				log.Fatalf("Failed to connect to MQTT broker: %v", session.Error())
			}
			stationRepository := mqtt.NewPublisherMQTTRepository(client)
			id, value := entity.NewSensorPayload(
				sensor.ID,
				map[string][]float64{"co2": {0, 1000, 1.96}, "co": {0, 15, 1.96}, "no2": {0, 1130, 1.96}, "mp10": {0, 250, 1.96}, "mp25": {0, 125, 1.96}, "rad": {1, 1280, 1.96}},
			)
			log := entity.NewLog(id, value)
			stationRepository.Publish(log)
		}(sensor)
	}
	wg.Wait()
}
