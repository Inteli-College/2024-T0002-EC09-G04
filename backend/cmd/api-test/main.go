package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/streadway/amqp"
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received: %s from: %s\n", msg.Payload(), msg.Topic())
	// Enviar a mensagem para o RabbitMQ
	sendToRabbitMQ(msg.Payload())
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
