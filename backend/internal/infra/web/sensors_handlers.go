package web

import (
	"encoding/json"
	"net/http"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/usecase"
)


type SensorHandlers struct {
	CreateSensorUseCase   *usecase.CreateSensorUseCase
}

func NewSensorHandlers(createSensorUseCase *usecase.CreateSensorUseCase) *SensorHandlers {
	return &SensorHandlers{CreateSensorUseCase: createSensorUseCase}
}

func (s *SensorHandlers) CreateSensorHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateSensorInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := s.CreateSensorUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}