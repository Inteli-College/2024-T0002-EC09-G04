package usecase

import (
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
)

type UserConfirmationUsecase struct {
	UserRepository entity.UserRepository
}

type UserConfirmationInputDTO struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

func NewUserConfirmationUsecase(userRepository entity.UserRepository) *UserConfirmationUsecase {
	return &UserConfirmationUsecase{UserRepository: userRepository}
}

func (u *UserConfirmationUsecase) Execute(input *UserConfirmationInputDTO) error {
	confirmation := entity.NewConfirmation(input.Email, input.Code)
	return u.UserRepository.ConfirmAccount(confirmation)
}