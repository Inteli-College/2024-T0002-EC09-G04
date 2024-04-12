package usecase

import (
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
)

type CreateUserUsecase struct {
	UserRepository entity.UserRepository
}

type CreateUserInputDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewCreateUserUsecase(userRepository entity.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{UserRepository: userRepository}
}

func (u *CreateUserUsecase) Execute(input *CreateUserInputDTO) error {
	user := entity.NewUser(input.Name, input.Email, input.Password)
	return u.UserRepository.SignUp(user)
}
