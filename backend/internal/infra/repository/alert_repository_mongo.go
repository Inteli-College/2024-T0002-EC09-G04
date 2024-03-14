package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlertRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewAlertRepositoryMongo(client *mongo.Client, dbName string, collectionName string) *AlertRepositoryMongo {
	collection := client.Database(dbName).Collection(collectionName)
	return &AlertRepositoryMongo{
		Collection: collection,
	}
}

func (a *AlertRepositoryMongo) CreateAlert(alert *entity.Alert) (*mongo.InsertOneResult, error) {
	result, err := a.Collection.InsertOne(context.TODO(), alert)
	log.Printf("Inserting alert into the MongoDB collection")
	return result, err
}

func (a *AlertRepositoryMongo) FindAllAlerts() ([]*entity.Alert, error) {
	cur, err := a.Collection.Find(context.TODO(), bson.D{})
	log.Printf("Selecting all alerts from the MongoDB collection %s", a.Collection.Name())
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var alerts []*entity.Alert
	for cur.Next(context.TODO()) {
		var alert bson.M
		err := cur.Decode(&alert)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found")
			continue
		} else if err != nil {
			return nil, err
		}

		jsonAlertData, err := json.MarshalIndent(alert, "", " ")
		if err != nil {
			return nil, err
		}

		var alertData entity.Alert
		err = json.Unmarshal(jsonAlertData, &alertData)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, &alertData)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return alerts, nil
}
