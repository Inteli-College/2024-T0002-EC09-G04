package usecase

import (
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
)

type UserSignInUsecase struct {
	UserRepository entity.UserRepository
}

type UserSignInInputDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewUserSignInUsecase(userRepository entity.UserRepository) *UserSignInUsecase {
	return &UserSignInUsecase{UserRepository: userRepository}
}

func (u *UserSignInUsecase) Execute(input *UserSignInInputDTO) (string, error) {
	login := entity.NewLogin(input.Email, input.Password)
	token, err := u.UserRepository.SignIn(login)
	if err != nil {
		return "", err
	}
	return token, nil
}