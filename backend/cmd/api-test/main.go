package main

import (
	"fmt"
	"encoding/json"
	"database/sql"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/gas"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/rad_lum"
	_ "github.com/lib/pq" // Import do driver PostgreSQL
)

const (
	postgreSQLInfo = "host=database-stations.cxic0so62a43.us-east-1.rds.amazonaws.com user=postgres password=admin1234 dbname=postgres sslmode=disable"
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received: %s from: %s\n", msg.Payload(), msg.Topic())
	// Enviar a mensagem para o RabbitMQ
	sendToRabbitMQ(msg.Payload())
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendToRabbitMQ(message []byte) {

	// Configurar a conexão com o RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/") // Alterado para rabbitmq
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", err)
		return
	}
	defer conn.Close()

	// Criar um canal
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel:", err)
		return
	}
	defer ch.Close()

	// Declarar uma fila
	q, err := ch.QueueDeclare(
		"stations_queue", // Nome da fila
		true,            // Durable: true
		false,           // Delete when unused
		false,           // Exclusive
		false,           // No-wait
		nil,             // Arguments
	)
	
	if err != nil {
		fmt.Println("Failed to declare a queue:", err)
		return
	}

	// Publicar a mensagem na fila
	err = ch.Publish(
		"",     // Exchange
		q.Name, // Routing key
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		fmt.Println("Failed to publish a message:", err)
		return
	}

	fmt.Println("Message sent to RabbitMQ:", string(message))

	msgs, err := ch.Consume(
		"stations_queue",
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

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://broker:1883")
	opts.SetClientID("subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("/stations", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	select {}
}
