package test_test

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userUsecaseSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the funcionalities we need to test
	usecase domain.UserUseCase

	// some helper function to clean-up any used tables
}

type userRepositoryMock struct {
	mock.Mock
}

type PasswordService struct {
	mock.Mock
}

type TokenService struct {
	mock.Mock
}

func (m *userRepositoryMock) CreateUser(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *userRepositoryMock) GetUser(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *userRepositoryMock) PromoteUser(userId string) (domain.User, error) {
	args := m.Called(userId)
	return args.Get(0).(domain.User), args.Error(1)

}

func (m *PasswordService) HashePassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *PasswordService) ComparePassword(hashedPassword string, password string) bool {
	args := m.Called(hashedPassword, password)
	return args.Bool(0)
}

func (m *TokenService) TokenGeneretor(claims map[string]interface{}) (string, error) {
	args := m.Called(claims)
	return args.String(0), args.Error(1)
}

func (m *TokenService) TotokenParser(token string) (jwt.MapClaims, error) {
	args := m.Called(token)
	return args.Get(0).(jwt.MapClaims), args.Error(1)

}

func (suit *userUsecaseSuite) SetupSuite() {
	// this function runs once before all tests in the suite

	// some initialization setup
	repository := new(userRepositoryMock)
	PasswordService := new(PasswordService)
	TokenService := new(TokenService)
	usecase := usecase.NewUserUsecase(repository, PasswordService, TokenService)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suit.usecase = usecase
}

func (suit *userUsecaseSuite) TearDownSuite() {

}

func (suit *userUsecaseSuite) TestCreateUser() {
	user := domain.User{
		Email:    "job.com",
		Password: "123456",
		Role:     "admin",
	}

	mockRepo := new(userRepositoryMock)
	mockRepo.On("CreateUser", mock.Anything).Return(user, nil)
	mockPassService := new(PasswordService)
	mockPassService.On("HashePassword", mock.Anything).Return("123456", nil)

	suit.usecase = usecase.NewUserUsecase(mockRepo, mockPassService, new(TokenService))
	result, err := suit.usecase.CreateUser(user)
	fmt.Println("this is the orginal user", user)
	fmt.Println("this is the result", result)
	suit.NoError(err)
	suit.Equal(user, result)

}

func (suit *userUsecaseSuite) TestLoginUser() {
	user := domain.User{
		Email:    "job.com",
		Password: "123456",
		Role:     "admin",
	}

	mockRepo := new(userRepositoryMock)
	mockRepo.On("GetUser", mock.Anything).Return(user, nil)
	mockPassService := new(PasswordService)
	mockPassService.On("ComparePassword", mock.Anything, mock.Anything).Return(true)
	mockTokenService := new(TokenService)
	mockTokenService.On("TokenGeneretor", mock.Anything).Return("token", nil)

	suit.usecase = usecase.NewUserUsecase(mockRepo, mockPassService, mockTokenService)
	result, err := suit.usecase.LoginUser(user)
	suit.NoError(err)
	suit.Equal("token", result)
}

func (suit *userUsecaseSuite) TestPromoteUser() {

	promotedUser := domain.User{
		Email:    "job.com",
		Password: "123456",
		Role:     "user",
	}

	mockRepo := new(userRepositoryMock)
	mockRepo.On("PromoteUser", mock.Anything).Return(promotedUser, nil)

	suit.usecase = usecase.NewUserUsecase(mockRepo, new(PasswordService), new(TokenService))
	result, err := suit.usecase.PromoteUser("1")
	suit.NoError(err)
	suit.Equal(promotedUser, result)
}

func Test_userUsecaseSuite(t *testing.T) {
	/// we still need this to run all tests in our suite
	suite.Run(t, &userUsecaseSuite{})
}
