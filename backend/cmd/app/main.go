package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/kafka"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/repository"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/web"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/usecase"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

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

	sensorsRepository := repository.NewSensorRepositoryPostgres(db)
	createSensorLogUseCase := usecase.NewCreateSensorLogUseCase(sensorsRepository)
	createSensorUseCase := usecase.NewCreateSensorUseCase(sensorsRepository)
	sensorHandlers := web.NewSensorHandlers(createSensorUseCase)

	alertRepository := repository.NewAlertRepositoryPostgres(db)
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
		_, err = createSensorLogUseCase.Execute(dto)
		if err != nil {
			log.Fatalf("Failed to create sensor log: %v", err)
		}
	}
}
