---
title: Banco de dados
sidebar_position: 5
slug: /nosql-db
---

O MongoDB oferece uma solução flexível e escalável para projetos que demandam agilidade na manipulação de dados. Sua estrutura baseada em documentos JSON permite adaptar o esquema conforme necessário, enquanto sua capacidade de escala horizontal e desempenho rápido o tornam ideal para lidar com grandes volumes de dados e cargas de trabalho intensivas. Além disso, sua ampla gama de recursos, suporte a consultas complexas e robusto ecossistema de ferramentas e comunidade contribuem para sua popularidade e eficácia em diversos cenários de aplicativos que buscam alta escala e disponibilidade de dados. Com base nisso embasamos a nossa escolha por essa estratégia em detrimento da centrada em um banco de dados relacional.

### Componentes:

Este recorte implementa um repositório para operações CRUD relacionadas a alertas em um banco de dados MongoDB. Ele utiliza a biblioteca oficial do MongoDB para Go, define uma estrutura AlertRepositoryMongo com métodos para criar alertas e recuperar todos os alertas armazenados. A função CreateAlert insere um alerta no MongoDB, enquanto a função FindAllAlerts recupera todos os alertas da coleção, decodificando os documentos BSON para a estrutura de dados apropriada. Em resumo, o código oferece uma abstração para interagir eficientemente com o MongoDB no contexto do gerenciamento de alertas:

```golang
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
```
Este outro recorte para interação com o banco de dados, implementa um repositório para logs de sensores no MongoDB. Ele fornece métodos para criar e armazenar logs no banco de dados, simplificando a interação entre o programa e o MongoDB:

```golang
package repository

import (
	"context"
	"log"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type SensorLogRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewSensorLogRepositoryMongo(client *mongo.Client, dbName string, collection string) *SensorRepositoryMongo {
	sensorsColl := client.Database(dbName).Collection(collection)
	return &SensorRepositoryMongo{
		Collection: sensorsColl,
	}
}

func (s *SensorRepositoryMongo) CreateSensorLog(sensorLog *entity.Log) error {
	result, err := s.Collection.InsertOne(context.TODO(), sensorLog)
	log.Printf("Inserting log into the MongoDB collection with id: %s", result)
	return err
}
```

Este, por sua vez, implementa um repositório para interagir com um banco de dados MongoDB, permitindo criar sensores e recuperar informações sobre todos os sensores armazenados. Ele encapsula operações de criação e leitura, facilitando a interação entre o programa e o MongoDB:

```golang
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

type SensorRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewSensorRepositoryMongo(client *mongo.Client, dbName string, sensorsCollection string) *SensorRepositoryMongo {
	collection := client.Database(dbName).Collection(sensorsCollection)
	return &SensorRepositoryMongo{
		Collection: collection,
	}
}

func (s *SensorRepositoryMongo) CreateSensor(sensor *entity.Sensor) (*mongo.InsertOneResult, error) {
	result, err := s.Collection.InsertOne(context.TODO(), sensor)
	log.Printf("Inserting sensor %s into the MongoDB collection: %s", result, s.Collection.Name())
	return result, err
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
		} else if err != nil {
			return nil, err
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
```
