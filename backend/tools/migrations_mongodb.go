// init-mongo.go
package main

import (
	"os"
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	options := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s",
			os.Getenv("MONGODB_ATLAS_USERNAME"),
			os.Getenv("MONGODB_ATLAS_PASSWORD"),
			os.Getenv("MONGODB_ATLAS_CLUSTER_HOSTNAME"),
			os.Getenv("MONGODB_ATLAS_APP_NAME")))
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established successfully")
	db := client.Database("mongodb")
	collection := db.Collection("sensors")
	documents := []interface{}{
		map[string]interface{}{
			"name":      "MICS-6814",
			"latitude":  -23.562387,
			"longitude": -46.711777,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.564137,
			"longitude": -46.711639,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "RXW-LIB-900",
			"latitude":  -23.565203,
			"longitude": -46.709176,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "MICS-6814",
			"latitude":  -23.567714,
			"longitude": -46.708461,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.564491,
			"longitude": -46.716086,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.564978,
			"longitude": -46.713122,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.573328,
			"longitude": -46.706395,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.572475,
			"longitude": -46.708614,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.570419,
			"longitude": -46.717631,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.578249,
			"longitude": -46.709088,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.571875,
			"longitude": -46.705869,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.580818,
			"longitude": -46.706969,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.581581,
			"longitude": -46.713595,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.576305,
			"longitude": -46.711314,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.570144,
			"longitude": -46.711090,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
	}
	insertResult, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Documents inserted. IDs: %v\n", insertResult.InsertedIDs)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
