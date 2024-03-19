---
title: Simulador MQTT
sidebar_position: 5
slug: /mqtt-simulator
---

O Simulador MQTT cumpre um papel fundamental no desenvolvimento da solução aqui apresentada, já que permite a simulação de uma grande massa de dados sendo transmitida através de tópicos MQTT, sem a necessidade de dispositivos físicos.

### Componentes:
Este recorte do sistema é responsável por hidratar entidades do tipo retratado abaixo passando os seus respectivos parâmetros. Nesse sentido, temos, para cada sensor, um payload criado a partir de uma relação que calcula o [intervalo de confiaça](https://en.wikipedia.org/wiki/Confidence_interval) entre o intervalo do dados fornecido a partir do [z-crítico](https://pt.wikipedia.org/wiki/Testes_de_hip%C3%B3teses) também definido nos parâmetros da função "NewSensorPayload".

```golang
package main

import (
	"context"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/henriquemarlon/hipercongo/internal/domain/entity"
	"github.com/henriquemarlon/hipercongo/internal/infra/repository"
	"github.com/henriquemarlon/hipercongo/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
	"time"
)

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

	sensors, err := findAllSensorsUseCase.Execute()
	if err != nil {
		log.Fatalf("Failed to find all sensors: %v", err)
	}

	var wg sync.WaitGroup
	for _, sensor := range sensors {
		wg.Add(1)
		log.Printf("Starting sensor: %v", sensor)
		go func(sensor usecase.FindAllSensorsOutputDTO) {
			defer wg.Done()
			opts := MQTT.NewClientOptions().AddBroker(
				fmt.Sprintf("ssl://%s:%s",
					os.Getenv("BROKER_TLS_URL"),
					os.Getenv("BROKER_PORT"))).SetUsername(
				os.Getenv("BROKER_USERNAME")).SetPassword(
				os.Getenv("BROKER_PASSWORD")).SetClientID(sensor.ID)
			client := MQTT.NewClient(opts)
			if session := client.Connect(); session.Wait() && session.Error() != nil {
				log.Fatalf("Failed to connect to MQTT broker: %v", session.Error())
			}
			for {
				payload, err := entity.NewSensorPayload(
					sensor.ID,
					sensor.Params,
					time.Now(),
				)
				if err != nil {
					log.Fatalf("Failed to create payload: %v", err)
				}

				jsonBytesPayload, err := json.Marshal(payload)
				if err != nil {
					log.Println("Error converting to JSON:", err)
				}

				token := client.Publish("sensors", 1, false, string(jsonBytesPayload))
				log.Printf("Published: %s, on topic: %s", string(jsonBytesPayload), "sensors")
				token.Wait()
				time.Sleep(120 * time.Second)
			}
		}(sensor)
	}
	wg.Wait()
}
```

