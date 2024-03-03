package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received: %s from: %s\n", msg.Payload(), msg.Topic())
}

func main() {
	opts := MQTT.NewClientOptions().AddBroker("ssl://908447a66faf43129ef280ff434012e6.s1.eu.hivemq.cloud:8883/mqtt")
	opts.SetClientID("subscriber")
	opts.SetUsername("inteli")
	opts.SetPassword("@Inteli123")
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