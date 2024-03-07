package repository

import (
	"context"
	"log"

	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
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

func (a *AlertRepositoryMongo) CreateAlert(alert *entity.Alert) error {
	_, err := a.Collection.InsertOne(context.TODO(), alert)
	log.Printf("Inserting alert into the MongoDB collection")
	return err
}

func (a *AlertRepositoryMongo) FindAllAlerts() ([]*entity.Alert, error) {
	cur, err := a.Collection.Find(context.TODO(), bson.D{})
	log.Printf("Selecting all alerts from the MongoDB collection")
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var alerts []*entity.Alert
	for cur.Next(context.TODO()) {
		var alert entity.Alert
		err := cur.Decode(&alert)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, &alert)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return alerts, nil
}
