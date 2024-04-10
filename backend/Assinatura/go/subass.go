package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	brokerAddress = "4384d919786d473a98b4000b9c180423.s1.eu.hivemq.cloud"
	port          = 8883
	topic         = "data/sensor1"
	username      = "riqueschilder"
	password      = "Riqueschilder123"
)

func getPublicKey() *rsa.PublicKey {
	// Load RSA public key
	keyFile, err := os.Open("public_key.pem")
	if err != nil {
		log.Fatalf("Failed to open public key file: %v", err)
	}
	defer keyFile.Close()

	keyBytes, err := io.ReadAll(keyFile)
	if err != nil {
		log.Fatalf("Failed to read public key file: %v", err)
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		log.Fatal("Failed to decode PEM block containing public key")
	}

	pubkeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatalf("Failed to parse public key: %v", err)
	}

	publicKey, ok := pubkeyInterface.(*rsa.PublicKey)
	if !ok {
		log.Fatal("Failed to convert public key to RSA public key")
	}

	return publicKey
}

type SensorData struct {
	CO       float64 `json:"CO"`
	NO2      float64 `json:"NO2"`
	Ethanol  float64 `json:"Ethanol"`
	Hydrogen float64 `json:"Hydrogen"`
	Ammonia  float64 `json:"Ammonia"`
}

func main() {
	opts := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("ssl://%s:%d", brokerAddress, port)).
		SetClientID("Subscriber").
		SetUsername(username).
		SetPassword(password).
		SetTLSConfig(&tls.Config{InsecureSkipVerify: true})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer client.Disconnect(250)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		payload := msg.Payload()
		receivedData := payload[:len(payload)-getSignatureSize()]

		signature := payload[len(payload)-getSignatureSize():]
		valid, err := verifySignature(receivedData, signature)
		if err != nil {
			log.Printf("Error verifying signature: %v", err)
			return
		}

		if valid {
			var data SensorData
			if err := json.Unmarshal(receivedData, &data); err != nil {
				log.Printf("Error unmarshalling sensor data: %v", err)
				return
			}

			log.Printf("Received valid sensor data: %+v", data)
			// Aqui você pode adicionar código para manipular os dados recebidos
		} else {
			log.Println("Received invalid sensor data. Signature verification failed.")
		}
	})

	<-sigChan
}

func verifySignature(data []byte, signature []byte) (bool, error) {
	hash := sha256.Sum256(data)
	publicKey := getPublicKey()
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getSignatureSize() int {
	publicKey := getPublicKey()
	return publicKey.Size()
}
