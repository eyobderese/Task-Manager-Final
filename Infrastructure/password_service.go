package infrastructure

import (
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/usecase"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func NewPasswordInfrastructureService() usecase.PasswordInfrastructure {
	return &PasswordService{}
}

func (pr *PasswordService) HashePassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return " ", err
	}

	return string(hashedPassword), nil

}

func (pr *PasswordService) ComparePassword(existingPassword string, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(newPassword))

	if err != nil {
		return false
	}

	return true

}
