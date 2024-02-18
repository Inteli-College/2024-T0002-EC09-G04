package station

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func TestGenerateLocation(t *testing.T) {
	data := GenerateLocation()
	if data.Latitude < Area["latitude"].Minimum || data.Latitude > Area["latitude"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "latitude", Area["latitude"].Minimum, Area["latitude"].Maximum)
	}
	if data.Longitude < Area["longitude"].Minimum || data.Longitude > Area["longitude"].Maximum {
		t.Errorf("%s: outside range [%f, %f]", "longitude", Area["longitude"].Minimum, Area["longitude"].Maximum)
	}
}

func TestConnectMQTT(t *testing.T) {
	client := ConnectMQTT(rand.NewSource(time.Now().UnixNano()), "tcp://localhost:1891")
	defer client.Disconnect(500)
	if !client.IsConnected() {
		t.Errorf("Unable to connect to MQTT broker\x1b[0m")
	}
}

func TestMessageTransmission(t *testing.T) {
	var receipts []string

	var handler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		receipts = append(receipts, fmt.Sprintf("New message on topic %s: %s", msg.Topic(), msg.Payload()))
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("test")
	opts.SetDefaultPublishHandler(handler)

	client := MQTT.NewClient(opts)
	defer client.Disconnect(500)

	if session := client.Connect(); session.Wait() && session.Error() != nil {
		panic(session.Error())
	}

	done := make(chan bool)

	go func() {
		Start("tcp://localhost:1891")
		done <- true
	}()

	go func() {
		if token := client.Subscribe("/stations", 1, nil); token.Wait() && token.Error() != nil {
			t.Logf("Error subscribing: %s", token.Error())
			return
		}
	}()

	time.Sleep(2 * time.Second)

	if len(receipts) == 0 {
		t.Errorf("No messages received")
	} else {
		for _, receipt := range receipts {
			t.Log(receipt)
		}
	}
}