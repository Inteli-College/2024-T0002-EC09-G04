package usecase

import (
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateSensorUseCase struct {
	SensorRepository entity.SensorRepository
}

//TODO: This need be more idiomatic, removing entity.Param type from DTOs

type CreateSensorInputDTO struct {
	Name      string                  `json:"name"`
	Latitude  float64                 `json:"latitude"`
	Longitude float64                 `json:"longitude"`
	Params    map[string]entity.Param `json:"params"`
}

type CreateSensorOutputDTO struct {
	ID        primitive.ObjectID      `json:"_id"`
	Name      string                  `json:"name"`
	Latitude  float64                 `json:"latitude"`
	Longitude float64                 `json:"longitude"`
	Params    map[string]entity.Param `json:"params"`
}

func NewCreateSensorUseCase(sensorRepository entity.SensorRepository) *CreateSensorUseCase {
	return &CreateSensorUseCase{SensorRepository: sensorRepository}
}

func (c *CreateSensorUseCase) Execute(input CreateSensorInputDTO) (*CreateSensorOutputDTO, error) {
	sensor := entity.NewSensor(input.Name, input.Latitude, input.Longitude, input.Params)
	id, err := c.SensorRepository.CreateSensor(sensor)
	if err != nil {
		return nil, err
	}
	return &CreateSensorOutputDTO{
		ID:        id.InsertedID.(primitive.ObjectID),
		Name:      sensor.Name,
		Latitude:  sensor.Latitude,
		Longitude: sensor.Longitude,
		Params:    sensor.Params,
	}, nil
}
