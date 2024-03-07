package usecase

import (
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
	"time"
)

type FindAllAlertsUseCase struct {
	AlertRepository entity.AlertRepository
}

type FindAlertsOutputDTO struct {
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
			Option:    alert.Option,
			Latitude:  alert.Latitude,
			Longitude: alert.Longitude,
			Timestamp: alert.Timestamp,
		})
	}
	return output, nil
}