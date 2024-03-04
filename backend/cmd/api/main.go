package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/pkg/station"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type GasData struct {
	CO2  float64 `json:"CO2"`
	CO   float64 `json:"CO"`
	NO2  float64 `json:"NO2"`
	MP10 float64 `json:"MP10"`
	MP25 float64 `json:"MP25"`
}

type RadLumData struct {
	ET  float64 `json:"ET"`
	LI  float64 `json:"LI"`
	SR  float64 `json:"SR"`
	Timestamp string `json:"timestamp"`
}

type Data struct {
	Location  string `json:"location"`
	Gas       string `json:"gas"`
	RadLum    string `json:"rad_lum"`
	Timestamp string `json:"timestamp"`
}

func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	sendToRabbitMQ(msg.Payload())
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendToRabbitMQ(message []byte) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/") // Updated to rabbitmq
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel:", err)
		return
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"stations_queue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("Failed to declare a queue:", err)
		return
	}
	fmt.Println("Queue status:", queue)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := message

	err = ch.PublishWithContext(ctx,
		"",
		"stations_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

	fmt.Println("Message sent to RabbitMQ:", string(message))

	err = ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		"stations_queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var data Data
			if err := json.Unmarshal(d.Body, &data); err != nil {
				log.Printf("Error decoding JSON message: %v", err)
				return
			}

			var gasData GasData
			if err := json.Unmarshal([]byte(data.Gas), &gasData); err != nil {
				log.Printf("Error decoding JSON message: %v", err)
				return
			}

			var radLumData RadLumData
			if err := json.Unmarshal([]byte(data.RadLum), &radLumData); err != nil {
				log.Printf("Error decoding JSON message: %v", err)
				return
			}

			if err := insertGasAndRadLumData(gasData, radLumData); err != nil {
				log.Printf("Error inserting data into the database: %v", err)
			}
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// func insertGasData(data GasData) error {
//     db, err := sql.Open("postgres", "postgresql://postgres:admin1234@postgres:5432/postgres?sslmode=disable")
//     if err != nil {
//         fmt.Println("Erro ao conectar ao banco de dados:", err)
//         return err
//     }
//     defer db.Close()

//     stmt, err := db.Prepare("INSERT INTO Gas (co2, co, no2, mp10, mp25, timestamp) VALUES ($1, $2, $3, $4, $5, $6)")
//     if err != nil {
//         log.Fatal(err)
//     }

//     _, err = stmt.Exec(data.CO2, data.CO, data.NO2, data.MP10, data.MP25, time.Now())
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Println("Dados inseridos com sucesso.")
//     return nil
// }

func insertGasAndRadLumData(gasData GasData, radLumData RadLumData) error {
	db, err := sql.Open("postgres", "postgresql://postgres:admin1234@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Gas (co2, co, no2, mp10, mp25, timestamp) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(gasData.CO2, gasData.CO, gasData.NO2, gasData.MP10, gasData.MP25, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("INSERT INTO Rad_Lum (et, li, sr, timestamp) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(radLumData.ET, radLumData.LI, radLumData.SR, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dados de gás e radiação/luminosidade inseridos com sucesso.")
	return nil
}

func main() {
	// Crie uma nova fonte de números aleatórios
	randomSource := rand.NewSource(time.Now().UnixNano())

	// Forneça a fonte de números aleatórios e a URL do broker MQTT para a função ConnectMQTT
	client := station.ConnectMQTT(randomSource, "ssl://908447a66faf43129ef280ff434012e6.s1.eu.hivemq.cloud:8883/mqtt:1883")

	// Inscreva-se no tópico MQTT onde a estação está publicando os dados
	topic := "/stations"
	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("Inscrito no tópico: %s\n", topic)

	// Aguarde um sinal de término para sair graciosamente
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	// Desinscreva-se do tópico MQTT e desconecte o cliente
	client.Unsubscribe(topic)
	client.Disconnect(250)
	fmt.Println("Desconectado do broker MQTT")

	select {}
}
