package test_test

import (
	"testing"
	"time"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type taskUsecaseSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the funcionalities we need to test
	usecase domain.TaskUsecase

	// some helper function to clean-up any used tables
}

type taskRepositoryMock struct {
	mock.Mock
}

func (m *taskRepositoryMock) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *taskRepositoryMock) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *taskRepositoryMock) GetTaskById(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *taskRepositoryMock) UpdateTask(task domain.Task, id string) (domain.Task, error) {
	args := m.Called(task, id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *taskRepositoryMock) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (suite *taskUsecaseSuite) SetupSuite() {
	// this function runs once before all tests in the suite

	// some initialization setup

	repository := new(taskRepositoryMock)
	usecase := usecase.NewTaskUsecase(repository)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.usecase = usecase

}

func (suite *taskUsecaseSuite) TearDownSuite() {
	// we clean the suite after the suite excute all tasks
}

func (suite *taskUsecaseSuite) TestCreateTask_Positive() {
	// instantiate an entity to be used by the function we want to test
	mockRepo := new(taskRepositoryMock)
	mockRepo.On("CreateTask", mock.Anything).Return(nil)
	suite.usecase = usecase.NewTaskUsecase(mockRepo)
	task := domain.Task{
		Title:       "test",
		Description: "test",
	}

	// real function we need to test
	err := suite.usecase.CreateTask(task)
	// assertion for the result of our test
	suite.NoError(err)
}

func (suit *taskUsecaseSuite) TestGetTasks_Positive() {
	mockRepo := new(taskRepositoryMock)
	task := domain.Task{
		Title:       "Task 1",
		Description: "First task",
		DueDate:     time.Now(),
	}

	mockRepo.On("GetTasks").Return([]domain.Task{task}, nil)
	suit.usecase = usecase.NewTaskUsecase(mockRepo)

	tasks, err := suit.usecase.GetTasks()
	suit.NoError(err)
	suit.NotNil(tasks)
	suit.NotEmpty(tasks)
}

func (suit *taskUsecaseSuite) TestGetTaskById_Positive() {
	mockRepo := new(taskRepositoryMock)
	task := domain.Task{
		Title:       "Task 1",
		Description: "First task",
		DueDate:     time.Now(),
	}

	mockRepo.On("GetTaskById", "1").Return(task, nil)
	suit.usecase = usecase.NewTaskUsecase(mockRepo)

	task, err := suit.usecase.GetTaskById("1")
	suit.NoError(err)
	suit.NotNil(task)
}

func (suit *taskUsecaseSuite) TestUpdateTask_Positive() {
	mockRepo := new(taskRepositoryMock)
	task := domain.Task{
		Title:       "Task 1",
		Description: "First task",
		DueDate:     time.Now(),
	}

	mockRepo.On("UpdateTask", task, "1").Return(task, nil)
	suit.usecase = usecase.NewTaskUsecase(mockRepo)

	task, err := suit.usecase.UpdateTask(task, "1")
	suit.NoError(err)
	suit.NotNil(task)
	suit.Equal("Task 1", task.Title)
}

func (suit *taskUsecaseSuite) TestDeleteTask_Positive() {
	mockRepo := new(taskRepositoryMock)
	mockRepo.On("DeleteTask", "1").Return(nil)
	suit.usecase = usecase.NewTaskUsecase(mockRepo)

	err := suit.usecase.DeleteTask("1")
	suit.Nil(err)
}

func Test_taskUsecaseSuite(t *testing.T) {
	/// we still need this to run all tests in our suite
	suite.Run(t, &taskUsecaseSuite{})
}
