package web

import (
	"encoding/json"
	"net/http"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/usecase"
)

type SensorHandlers struct {
	CreateSensorUseCase *usecase.CreateSensorUseCase
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
