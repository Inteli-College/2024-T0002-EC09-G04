package usecase

import (
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
)

type CreateSensorUseCase struct {
	SensorRepository entity.SensorRepository
}

type CreateSensorInputDTO struct {
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}

type CreateSensorOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}

func NewCreateSensorUseCase(sensorRepository entity.SensorRepository) *CreateSensorUseCase {
	return &CreateSensorUseCase{SensorRepository: sensorRepository}
}

func (c *CreateSensorUseCase) Execute(input CreateSensorInputDTO) (*CreateSensorOutputDTO, error) {
	sensor := entity.NewSensor(input.Name, input.Latitude, input.Longitude)
	err := c.SensorRepository.CreateSensor(sensor)
	if err != nil {
		return nil, err
	}
	return &CreateSensorOutputDTO{
		ID:        sensor.ID,
		Name:      sensor.Name,
		Latitude:  sensor.Latitude,
		Longitude: sensor.Longitude,
	}, nil
}
