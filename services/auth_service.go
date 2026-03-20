package services

import (
	"todo-backend/models"
	"todo-backend/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(email string, password string) error
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (r *AuthService) SignUp(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = r.repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
