package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type SensorRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewSensorRepositoryMongo(client *mongo.Client, dbName string, sensorsCollection string) *SensorRepositoryMongo {
	collection := client.Database(dbName).Collection(sensorsCollection)
	return &SensorRepositoryMongo{
		Collection: collection,
	}
}

func (s *SensorRepositoryMongo) CreateSensor(sensor *entity.Sensor) error {
	result, err := s.Collection.InsertOne(context.TODO(), sensor)
	log.Printf("Inserting sensor %s into the MongoDB collection: %s", result, s.Collection.Name())
	return err
}

func (s *SensorRepositoryMongo) FindAllSensors() ([]*entity.Sensor, error) {
	cur, err := s.Collection.Find(context.TODO(), bson.D{})
	log.Printf("Selecting all sensors from the MongoDB collection %s", s.Collection.Name())
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var sensors []*entity.Sensor
	for cur.Next(context.TODO()) {
		var sensor bson.M
		err := cur.Decode(&sensor)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found")
		}

		jsonSensorData, err := json.MarshalIndent(sensor, "", " ")
		if err != nil {
			return nil, err
		}

		var sensorData entity.Sensor
		err = json.Unmarshal(jsonSensorData, &sensorData)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensorData)
	}
	
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return sensors, nil
}
