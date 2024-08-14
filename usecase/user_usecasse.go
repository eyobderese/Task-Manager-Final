package usecase

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
)

type User = domain.User

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUser(user User) (User, error)
	PromoteUser(userId string) (User, error)
}

type PasswordInfrastructure interface {
	HashePassword(pass string) (string, error)
	ComparePassword(existingPassword string, newPassword string) bool
}

type TokenInfrastructure interface {
	TokenGeneretor(claims map[string]interface{}) (string, error)
	TotokenParser(authPartToken string) (jwt.MapClaims, error)
}

type userUseCase struct {
	userRepository         UserRepository
	passwordInfrastructure PasswordInfrastructure
	tokenInfrastructure    TokenInfrastructure
}

func NewUserUsecase(userRepository UserRepository, passwordInfrastructure PasswordInfrastructure, tokenInfrastructure TokenInfrastructure) domain.UserUseCase {
	return &userUseCase{userRepository: userRepository, passwordInfrastructure: passwordInfrastructure, tokenInfrastructure: tokenInfrastructure}
}

// CreateUser inserts a new user into the database.
func (uuc *userUseCase) CreateUser(user User) (User, error) {
	hashedPassword, err := uuc.passwordInfrastructure.HashePassword(user.Password)

	if err != nil {
		return User{}, err
	}
	user.Password = hashedPassword
	return uuc.userRepository.CreateUser(user)

}

func (uuc *userUseCase) LoginUser(user User) (string, error) {
	newUser, err := uuc.userRepository.GetUser(user)
	if err != nil {
		return "", err
	}

	if !uuc.passwordInfrastructure.ComparePassword(newUser.Password, user.Password) {
		return "", err
	}

	claims := map[string]interface{}{
		"id":    newUser.ID.Hex(), // Assuming ID is of type primitive.ObjectID
		"email": user.Email,
		"role":  newUser.Role,
	}

	token, err := uuc.tokenInfrastructure.TokenGeneretor(claims)

	// return the token
	return token, err

}

func (uuc *userUseCase) PromoteUser(userId string) (User, error) {
	return uuc.userRepository.PromoteUser(userId)
}
