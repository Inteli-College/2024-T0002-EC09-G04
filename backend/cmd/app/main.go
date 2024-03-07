package main

import (
	"context"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/infra/kafka"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/infra/repository"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/infra/web"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/usecase"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func main() {
	options := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=orbit-city", os.Getenv("MONGODB_ATLAS_USERNAME"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGODB_ATLAS_CLUSTER_HOSTNAME")))
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	msgChan := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"sasl.mechanisms":    "PLAIN",
		"security.protocol":  "SASL_SSL",
		"sasl.username":      os.Getenv("CONFLUENT_API_KEY"),
		"sasl.password":      os.Getenv("CONFLUENT_API_SECRET"),
		"session.timeout.ms": 6000,
		"group.id":           "orbit-city",
		"auto.offset.reset":  "latest",
	}

	kafkaRepository := kafka.NewKafkaConsumer(configMap, []string{os.Getenv("CONFLUENT_KAFKA_TOPIC_NAME")})
	go func() {
		if err := kafkaRepository.Consume(msgChan); err != nil {
			log.Printf("Error consuming kafka queue: %v", err)
		}
	}()

	sensorsLogRepository := repository.NewSensorLogRepositoryMongo(client, "mongodb", "sensors_log")
	createSensorLogUseCase := usecase.NewCreateSensorLogUseCase(sensorsLogRepository)
	sensorsRepository := repository.NewSensorRepositoryMongo(client, "mongodb", "sensors")
	createSensorUseCase := usecase.NewCreateSensorUseCase(sensorsRepository)
	sensorHandlers := web.NewSensorHandlers(createSensorUseCase)

	alertRepository := repository.NewAlertRepositoryMongo(client, "mongodb", "alerts")
	createAlertUseCase := usecase.NewCreateAlertUseCase(alertRepository)
	findAllAlertsUseCase := usecase.NewFindAllAlertsUseCase(alertRepository)
	alertHandlers := web.NewAlertHandlers(createAlertUseCase, findAllAlertsUseCase)

	//TODO: this is the best way to do this? need to refactor or find another way to start the server
	router := chi.NewRouter()
	router.Get("/alerts", alertHandlers.FindAllAlertsHandler)
	router.Post("/alerts", alertHandlers.CreateAlertHandler)
	router.Post("/sensors", sensorHandlers.CreateSensorHandler)
	go http.ListenAndServe(":8080", router)

	for msg := range msgChan {
		dto := usecase.CreateSensorLogInputDTO{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		}
		err = createSensorLogUseCase.Execute(dto)
		if err != nil {
			log.Fatalf("Failed to create sensor log: %v", err)
		}
	}
}
