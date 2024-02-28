package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/pkg/station"
)

func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func main() {
	// Configure o cliente MQTT
	client := station.ConnectMQTT()

	// Inscreva-se no tópico MQTT onde a estação está publicando os dados
	topic := "/stations"
	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)

	// Aguarde um sinal de interrupção para sair graciosamente
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	// Desinscreva-se do tópico MQTT e desconecte o cliente
	client.Unsubscribe(topic)
	client.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker")
}
