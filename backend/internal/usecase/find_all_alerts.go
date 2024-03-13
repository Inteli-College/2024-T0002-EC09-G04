package usecase

import (
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
)

type FindAllAlertsUseCase struct {
	AlertRepository entity.AlertRepository
}

type FindAlertsOutputDTO struct {
	ID        string    `json:"_id"`
	Option    string    `json:"option"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}

func NewFindAllAlertsUseCase(alertRepository entity.AlertRepository) *FindAllAlertsUseCase {
	return &FindAllAlertsUseCase{AlertRepository: alertRepository}
}

func (f *FindAllAlertsUseCase) Execute() ([]FindAlertsOutputDTO, error) {
	alerts, err := f.AlertRepository.FindAllAlerts()
	if err != nil {
		return nil, err
	}
	var output []FindAlertsOutputDTO
	for _, alert := range alerts {
		output = append(output, FindAlertsOutputDTO{
			ID:        alert.ID,
			Option:    alert.Option,
			Latitude:  alert.Latitude,
			Longitude: alert.Longitude,
			Timestamp: alert.Timestamp,
		})
	}
	return output, nil
}
