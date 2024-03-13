package web

import (
	"encoding/json"
	"net/http"

	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/usecase"
)

type AlertHandlers struct {
	CreateAlertUseCase   *usecase.CreateAlertUseCase
	FindAllAlertsUseCase *usecase.FindAllAlertsUseCase
}

func NewAlertHandlers(createAlertUseCase *usecase.CreateAlertUseCase, findAllAlertsUseCase *usecase.FindAllAlertsUseCase) *AlertHandlers {
	return &AlertHandlers{
		CreateAlertUseCase:   createAlertUseCase,
		FindAllAlertsUseCase: findAllAlertsUseCase,
	}
}

func (a *AlertHandlers) CreateAlertHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateAlertInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := a.CreateAlertUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (a *AlertHandlers) FindAllAlertsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := a.FindAllAlertsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
