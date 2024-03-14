package usecase

import (
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateAlertUseCase struct {
	AlertRepository entity.AlertRepository
}

type CreateAlertInputDTO struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Option    string  `json:"option"`
}

type CreateAlertOutputDTO struct {
	ID        primitive.ObjectID `json:"_id"`
	Latitude  float64            `json:"latitude"`
	Longitude float64            `json:"longitude"`
	Option    string             `json:"option"`
	Timestamp time.Time          `json:"timestamp"`
}

func NewCreateAlertUseCase(alertRepository entity.AlertRepository) *CreateAlertUseCase {
	return &CreateAlertUseCase{AlertRepository: alertRepository}
}

func (c *CreateAlertUseCase) Execute(input CreateAlertInputDTO) (*CreateAlertOutputDTO, error) {
	alert := entity.NewAlert(input.Latitude, input.Longitude, input.Option)
	id, err := c.AlertRepository.CreateAlert(alert)
	if err != nil {
		return nil, err
	}
	return &CreateAlertOutputDTO{
		ID:        id.InsertedID.(primitive.ObjectID),
		Latitude:  alert.Latitude,
		Longitude: alert.Longitude,
		Option:    alert.Option,
		Timestamp: alert.Timestamp,
	}, nil
}
