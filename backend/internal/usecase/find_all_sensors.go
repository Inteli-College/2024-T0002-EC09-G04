package usecase

import (
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
)

type FindAllSensorsUseCase struct {
	SensorRepository entity.SensorRepository
}

type FindAllSensorsOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	MinMax    []float64 `json:"min_max"`
	ZCrit     float64   `json:"zcritc"`
}

func NewFindAllSensorsUseCase(sensorRepository entity.SensorRepository) *FindAllSensorsUseCase {
	return &FindAllSensorsUseCase{SensorRepository: sensorRepository}
}

func (f *FindAllSensorsUseCase) Execute() ([]FindAllSensorsOutputDTO, error) {
	sensors, err := f.SensorRepository.FindAllSensors()
	if err != nil {
		return nil, err
	}
	var output []FindAllSensorsOutputDTO
	for _, sensor := range sensors {
		output = append(output, FindAllSensorsOutputDTO{
			ID:       sensor.ID,
			Name:     sensor.Name,
			Latitude: sensor.Latitude,
			Longitude: sensor.Longitude,
		})
	}
	return output, nil
}
