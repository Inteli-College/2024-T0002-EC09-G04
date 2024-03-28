package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/usecase"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/kafka"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/repository"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/web"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permitir todas as origens
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	//TODO: this is the best way to do this? need to refactor or find another way to start the server
	router := chi.NewRouter()
	router.Use(corsOptions.Handler) // Use o middleware CORS
	router.Get("/sensors", sensorHandlers.CreateSensorHandler)
	router.Get("/alerts", alertHandlers.CreateAlertHandler)
	router.Post("/alerts", alertHandlers.CreateAlertHandler)
	router.Post("/sensors", sensorHandlers.CreateSensorHandler)

	go http.ListenAndServe(":8000", router)

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
