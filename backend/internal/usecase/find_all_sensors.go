// usecase/find_all_sensors.go
package usecase

import (
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
)

type FindAllSensorsUseCase struct {
	SensorRepository entity.SensorRepository
}

type FindAllSensorsOutputDTO struct {
	ID        string                  `json:"sensor_id"`
	Name      string                  `json:"name"`
	Latitude  float64                 `json:"latitude"`
	Longitude float64                 `json:"longitude"`
	Params    map[string]entity.Param `json:"params"`
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
			ID:        sensor.ID,
			Name:      sensor.Name,
			Latitude:  sensor.Latitude,
			Longitude: sensor.Longitude,
			Params:    sensor.Params,
		})
	}
	return output, nil
}
