package main

import (
	"encoding/json"
	"database/sql"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/gas"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/rad_lum"
)

const (
	amqpURI        = "amqp://guest:guest@rabbitmq:5672/"
	postgreSQLInfo = "host=database-stations.cxic0so62a43.us-east-1.rds.amazonaws.com user=postgres password=admin1234 dbname=postgres sslmode=disable"
	queueName      = "stations_queue"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Conectar ao RabbitMQ
	conn, err := amqp.Dial(amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false, //autoAck=false para confirmação manual
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	// Conectar ao PostgreSQL
	db, err := sql.Open("postgres", postgreSQLInfo)
	failOnError(err, "Failed to open a DB connection")
	defer db.Close()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// Tente deserializar para Gas
			var gasData gas.Payload
			if err := json.Unmarshal(d.Body, &gasData); err == nil {
				// Insira os dados na tabela Gas
				_, err := db.Exec("INSERT INTO gas (CO2, CO, NO2, MP10, MP25) VALUES ($1, $2, $3, $4, $5)", gasData.CO2, gasData.CO, gasData.NO2, gasData.MP10, gasData.MP25)
				if err != nil {
					log.Printf("Error inserting Gas into DB: %v", err)
					continue
				}
				log.Printf("Gas data inserted successfully: %v", gasData)
			}

			// Tente deserializar para RadLum
			var radLumData rad_lum.Payload
			if err := json.Unmarshal(d.Body, &radLumData); err == nil {
				// Insira os dados na tabela Rad_lum
				_, err := db.Exec("INSERT INTO rad_lum (ET, LI, SR) VALUES ($1, $2, $3)", radLumData.ET, radLumData.LI, radLumData.SR)
				if err != nil {
					log.Printf("Error inserting RadLum into DB: %v", err)
					continue
				}
				log.Printf("RadLum data inserted successfully: %v", radLumData)
			}
		}
	}()

	log.Printf(" [*] Aguardando mensagens. Para sair, pressione CTRL+C")
	<-forever
}