package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	amqp "github.com/rabbitmq/amqp091-go"
	_ "github.com/lib/pq" // Import do driver PostgreSQL
)

// Estrutura para armazenar os dados da mensagem MQTT
type GasData struct {
	IDEstacao int     `json:"id_estacao"`
	CO2       float64 `json:"CO2"`
	CO        float64 `json:"CO"`
	NO2       float64 `json:"NO2"`
	MP10      float64 `json:"MP10"`
	MP25      float64 `json:"MP25"`
}

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

func insertGasData(data GasData) error {
	// Abrir conexão com o banco de dados PostgreSQL
	connStr := "user=postgres password=admin1234 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return err // Corrigido para retornar o erro
	}
	defer db.Close()

	// Preparar a declaração SQL
	stmt, err := db.Prepare("INSERT INTO Gas (id_estacao, CO2, CO, NO2, MP10, MP25) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executar a declaração SQL com os valores fornecidos
	_, err = stmt.Exec(data.IDEstacao, data.CO2, data.CO, data.NO2, data.MP10, data.MP25)
	if err != nil {
		return err
	}

	fmt.Println("Dados inseridos com sucesso.")
	return nil
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
	queue, err := ch.QueueDeclare(
		"stations_queue", // Nome da fila
		true,             // Durable: true
		false,            // Delete when unused
		false,            // Exclusive
		false,            // No-wait
		nil,              // Arguments
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
		"",               // exchange
		"stations_queue", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

	fmt.Println("Message sent to RabbitMQ:", string(message))

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		"stations_queue",
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			// Decodificar a mensagem recebida em JSON
			var gasData GasData
			if err := json.Unmarshal(d.Body, &gasData); err != nil {
				log.Printf("Erro ao decodificar a mensagem JSON: %v", err)
				return
			}

			// Inserir os dados no banco de dados
			if err := insertGasData(gasData); err != nil {
				log.Printf("Erro ao inserir dados no banco de dados: %v", err)
			}
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://broker:1883")
	opts.SetClientID("subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		fmt.Printf("Conexão perdida: %v\n", err)
		// Implemente a lógica de reconexão aqui
		for {
			if token := client.Connect(); token.Wait() && token.Error() == nil {
				fmt.Println("Conexão reestabelecida.")
				break
			}
			time.Sleep(5 * time.Second) // Espere um pouco antes de tentar reconectar
		}
	})

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
