package test_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/Delivery/controller"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUsecase) UpdateTask(task domain.Task, id string) (domain.Task, error) {
	args := m.Called(task, id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) GetTaskById(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

type TaskControllerSuit struct {
	suite.Suite
}

func (suite *TaskControllerSuit) SetupSuite() {

}
func (suite *TaskControllerSuit) TestCreateTask() {
	mockTaskUsecase := new(MockTaskUsecase)
	taskController := controller.NewTaskController(mockTaskUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/tasks", taskController.CreateTask)

	task := domain.Task{ /* fill with valid task data */ }
	mockTaskUsecase.On("CreateTask", task).Return(nil)

	body, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	suite.Equal(http.StatusCreated, resp.Code)
}

func (suite *TaskControllerSuit) TestUpdateTask() {
	mockTaskUsecase := new(MockTaskUsecase)
	taskController := controller.NewTaskController(mockTaskUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.PUT("/tasks/:id", taskController.UpdateTask)

	task := domain.Task{ /* fill with valid task data */ }
	mockTaskUsecase.On("UpdateTask", task, "1").Return(task, nil)

	body, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	suite.Equal(http.StatusOK, resp.Code)

}

func (suite *TaskControllerSuit) TestGetTasks() {
	mockTaskUsecase := new(MockTaskUsecase)
	taskController := controller.NewTaskController(mockTaskUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/tasks", taskController.GetTasks)

	tasks := []domain.Task{ /* fill with valid tasks data */ }
	mockTaskUsecase.On("GetTasks").Return(tasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)

}

func (suite *TaskControllerSuit) TestGetTaskById() {
	mockTaskUsecase := new(MockTaskUsecase)
	taskController := controller.NewTaskController(mockTaskUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/tasks/:id", taskController.GetTaskById)

	task := domain.Task{ /* fill with valid task data */ }
	mockTaskUsecase.On("GetTaskById", "1").Return(task, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *TaskControllerSuit) TestDeleteTask() {
	mockTaskUsecase := new(MockTaskUsecase)
	taskController := controller.NewTaskController(mockTaskUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.DELETE("/tasks/:id", taskController.DeleteTask)

	mockTaskUsecase.On("DeleteTask", "1").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	suite.Equal(http.StatusNoContent, resp.Code)

}

func Test_taskControllerSuite(t *testing.T) {
	/// we still need this to run all tests in our suite
	suite.Run(t, &TaskControllerSuit{})
}
