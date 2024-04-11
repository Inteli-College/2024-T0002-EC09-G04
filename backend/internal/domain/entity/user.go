package entity

import (
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type UserRepository interface {
	SignUp(user *User) error
	ConfirmAccount(confirmation *Confirmation) error
	SignIn(login *Login) (string, error)
	GetUserByToken(token string) (*cognito.GetUserOutput, error)
	UpdatePassword(login *Login) error
}

type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Confirmation struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func NewLogin(email string, password string) *Login {
	return &Login{
		Email:    email,
		Password: password,
	}
}

func NewConfirmation(email string, code string) *Confirmation {
	return &Confirmation{
		Email: email,
		Code:  code,
	}
}

