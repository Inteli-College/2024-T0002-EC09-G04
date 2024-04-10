package main

import (
	"crypto"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	brokerAddress = "4384d919786d473a98b4000b9c180423.s1.eu.hivemq.cloud"
	port          = 8883
	topic         = "data/sensor1"
	username      = "riqueschilder"
	password      = "Riqueschilder123"
	privateKey    *rsa.PrivateKey
)

func init() {
	// Load RSA private key
	keyFile, err := os.Open("private_key.pem")
	if err != nil {
		log.Fatalf("Failed to open private key file: %v", err)
	}
	defer keyFile.Close()

	keyBytes, err := io.ReadAll(keyFile)
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		log.Fatal("Failed to decode PEM block containing private key")
	}

	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
}

type SensorData struct {
	CO       float64 `json:"CO"`
	NO2      float64 `json:"NO2"`
	Ethanol  float64 `json:"Ethanol"`
	Hydrogen float64 `json:"Hydrogen"`
	Ammonia  float64 `json:"Ammonia"`
}

func generateSensorData() SensorData {
	return SensorData{
		CO:       randomFloat(1, 1000),
		NO2:      randomFloat(0.05, 10),
		Ethanol:  randomFloat(10, 500),
		Hydrogen: randomFloat(1, 1000),
		Ammonia:  randomFloat(1, 500),
	}
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func signData(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(crand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func main() {
	opts := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("ssl://%s:%d", brokerAddress, port)).
		SetClientID("Publisher").
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

	go func() {
		for {
			select {
			case <-sigChan:
				log.Println("Received termination signal. Exiting...")
				return
			default:
				data := generateSensorData()
				payload, err := json.Marshal(data)
				if err != nil {
					log.Printf("Error marshalling sensor data: %v", err)
					continue
				}

				signature, err := signData(payload)
				if err != nil {
					log.Printf("Error signing data: %v", err)
					continue
				}

				// Concatena a assinatura ao payload antes de publicar
				payloadWithSignature := append(payload, signature...)

				token := client.Publish(topic, 1, false, payloadWithSignature)
				token.Wait()
				log.Printf("Published message with signature")

				time.Sleep(5 * time.Second)
			}
		}
	}()

	<-sigChan
}
