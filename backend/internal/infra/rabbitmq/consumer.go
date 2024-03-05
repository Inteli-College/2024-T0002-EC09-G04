package rabbitmq

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	QueueName string
	RabbitMQURL string
}

func NewRabbitMQConsumer(queueName string, rabbitMQURL string) *RabbitMQConsumer {
	return &RabbitMQConsumer{QueueName: queueName, RabbitMQURL: rabbitMQURL}
}

func (r *RabbitMQConsumer) Consume(msgChan chan<- *amqp.Delivery) {
	conn, err := amqp.Dial(r.RabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	messages, err := ch.Consume(
		r.QueueName, // queue
		"",        // consumer
		true,      // auto-acknowledge
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for msg := range messages {
		log.Printf("Received message: %s", msg.Body)
		msgChan <- &msg
	}
}