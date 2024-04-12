package usecase

import (
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	"strconv"
)

type GetUserUsecase struct {
	UserRepository entity.UserRepository
}

type GetUserInputDTO struct {
	Token string
}

type GetUserOutputDTO struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	CustomID      string `json:"custom_id"`
	EmailVerified bool   `json:"email_verified"`
}

func NewGetUserUsecase(userRepository entity.UserRepository) *GetUserUsecase {
	return &GetUserUsecase{UserRepository: userRepository}
}

func (u *GetUserUsecase) Execute(input *GetUserInputDTO) (*GetUserOutputDTO, error) {
	output, err := u.UserRepository.GetUserByToken(input.Token)
	dto := &GetUserOutputDTO{}
	for _, attribute := range output.UserAttributes {
		switch *attribute.Name {
		case "sub":
			dto.ID = *attribute.Value
		case "name":
			dto.Name = *attribute.Value
		case "email":
			dto.Email = *attribute.Value
		case "custom:custom_id":
			dto.CustomID = *attribute.Value
		case "email_verified":
			verification, err := strconv.ParseBool(*attribute.Value)
			if err == nil {
				dto.EmailVerified = verification
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return dto, nil
}