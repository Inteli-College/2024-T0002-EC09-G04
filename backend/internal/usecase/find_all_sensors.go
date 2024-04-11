// usecase/find_all_sensors.go
package usecase

import "github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"

type FindAllSensorsUseCase struct {
	SensorRepository entity.SensorRepository
}

//TODO: This need be more idiomatic, removing Param type from DTOs

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
			ID:        sensor.Sensor_ID,
			Name:      sensor.Name,
			Latitude:  sensor.Latitude,
			Longitude: sensor.Longitude,
			Params:    sensor.Params,
		})
	}
	return output, nil
}
