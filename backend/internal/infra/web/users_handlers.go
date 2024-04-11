package web

import (
	"encoding/json"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/usecase"
	"net/http"
	"strings"
)

type UserHandlers struct {
	CreateUserUsecase       *usecase.CreateUserUsecase
	GetUserUsecase          *usecase.GetUserUsecase
	UserConfirmationUsecase *usecase.UserConfirmationUsecase
	UserSignInUsecase       *usecase.UserSignInUsecase
}

func NewUserHandlers(
	createUserUsecase *usecase.CreateUserUsecase,
	getUserUsecase *usecase.GetUserUsecase,
	userConfirmationUsecase *usecase.UserConfirmationUsecase,
	userSignInUsecase *usecase.UserSignInUsecase,
) *UserHandlers {
	return &UserHandlers{
		CreateUserUsecase:       createUserUsecase,
		GetUserUsecase:          getUserUsecase,
		UserConfirmationUsecase: userConfirmationUsecase,
		UserSignInUsecase:       userSignInUsecase,
	}
}

func (u *UserHandlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = u.CreateUserUsecase.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *UserHandlers) ValidateHandler(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := u.GetUserUsecase.Execute(&usecase.GetUserInputDTO{
		Token: token,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (u *UserHandlers) UserConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.UserConfirmationInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = u.UserConfirmationUsecase.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandlers) UserSignInHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.UserSignInInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := u.UserSignInUsecase.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}