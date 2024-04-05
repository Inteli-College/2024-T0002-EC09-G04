package test

import (
	"os"
	"testing"
	"fmt"

	godotenv "github.com/joho/godotenv"
)

func TestEnv(t *testing.T) {
	err := godotenv.Load("../config/.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	if os.Getenv("CONFLUENT_KAFKA_TOPIC_NAME") == "" {
		t.Errorf("CONFLUENT_KAFKA_TOPIC_NAME is not set")
	}
	if os.Getenv("CONFLUENT_API_KEY") == "" {
		t.Errorf("CONFLUENT_API_KEY is not set")
	}
	if os.Getenv("CONFLUENT_API_SECRET") == "" {
		t.Errorf("CONFLUENT_API_SECRET is not set")
	}
	if os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL") == "" {
		t.Errorf("CONFLUENT_BOOTSTRAP_SERVER_SASL is not set")
	}
	if os.Getenv("BROKER_TLS_URL") == "" {
		t.Errorf("BROKER_TLS_URL is not set")
	}
	if os.Getenv("BROKER_USERNAME") == "" {
		t.Errorf("BROKER_USERNAME is not set")
	}
	if os.Getenv("BROKER_PASSWORD") == "" {
		t.Errorf("BROKER_PASSWORD is not set")
	}
	if os.Getenv("BROKER_PORT") == "" {
		t.Errorf("BROKER_PORT is not set")
	}
	if os.Getenv("MONGODB_ATLAS_APP_NAME") == "" {
		t.Errorf("MONGODB_ATLAS_APP_NAME is not set")
	}
	if os.Getenv("MONGODB_ATLAS_USERNAME") == "" {
		t.Errorf("MONGODB_ATLAS_USERNAME is not set")
	}
	if os.Getenv("MONGODB_ATLAS_PASSWORD") == "" {
		t.Errorf("MONGODB_ATLAS_PASSWORD is not set")
	}
	if os.Getenv("MONGODB_ATLAS_CLUSTER_HOSTNAME") == "" {
		t.Errorf("MONGODB_ATLAS_CLUSTER_HOSTNAME is not set")
	}
	if os.Getenv("MB_DB_DBNAME") == "" {
		t.Errorf("MB_DB_DBNAME is not set")
	}
	if os.Getenv("MB_DB_TYPE") == "" {
		t.Errorf("MB_DB_TYPE is not set")
	}
	if os.Getenv("MB_DB_PORT") == "" {
		t.Errorf("MB_DB_PORT is not set")
	}
	if os.Getenv("MB_DB_USER") == "" {
		t.Errorf("MB_DB_USER is not set")
	}
	if os.Getenv("MB_DB_PASS") == "" {
		t.Errorf("MB_DB_PASS is not set")
	}
	if os.Getenv("MB_DB_HOST") == "" {
		t.Errorf("MB_DB_HOST is not set")
	}
	if os.Getenv("MB_BREAKOUT_BIN_WIDTH") == "" {
		t.Errorf("MB_BREAKOUT_BIN_WIDTH is not set")
	}
}