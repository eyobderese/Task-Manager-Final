package test_test

import (
	"context"
	"log"
	"testing"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/repositories"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/usecase"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepositorySuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the funcionalities we need to test
	repository usecase.TaskRepository
	db         *mongo.Database
	// some helper function to clean-up any used tables
}

func (suite *taskRepositorySuite) SetupSuite() {
	// this function runs once before all tests in the suite

	// some initialization setup
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("testdb")
	repository := repositories.NewTaskRepository(*db, "tasks")

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.repository = repository
	suite.db = db

}

func (suite *taskRepositorySuite) TearDownSuite() {
	// this function runs once after all tests in the suite
	// we need this to clean up any data we used in the tests
	// we need to drop the table we used in the tests

	defer suite.db.Drop(context.Background())

}

func (suite *taskRepositorySuite) TestCreateTask_Positive() {
	// instantiate an entity to be used by the function we want to test
	task := domain.Task{
		Title:       "test",
		Description: "test",
	}

	// real function we need to test
	err := suite.repository.CreateTask(task)
	// assertion for the result of our test
	suite.NoError(err, "no error when create tweet with valid input")
}

func (suite *taskRepositorySuite) TestGetTasks_Positive() {
	// real function we need to test
	tasks, err := suite.repository.GetTasks()
	// assertion for the result of our test
	suite.NoError(err, "no error when get tasks")
	suite.NotNil(tasks, "tasks are not nil")
}

func (suite *taskRepositorySuite) TestGetTaskById_Positive() {
	// instantiate an entity to be used by the function we want to test
	task := domain.Task{
		Title:       "test",
		Description: "test",
		ID:          primitive.NewObjectID(),
	}

	// real function we need to test
	err := suite.repository.CreateTask(task)

	suite.NoError(err, "no error when create tweet with valid input")

	// real function we need to test
	task, err = suite.repository.GetTaskById(task.ID.Hex())
	// assertion for the result of our test
	suite.NoError(err, "no error when get task by id")
	suite.NotNil(task, "task is not nil")
}

func (suite *taskRepositorySuite) TestUpdateTask_Positive() {
	task := domain.Task{
		Title:       "test",
		Description: "this is Test",
		ID:          primitive.NewObjectID(),
	}
	err := suite.repository.CreateTask(task)
	suite.NoError(err, "no error when create tweet with valid input")

	task.Description = "this is updated test"
	updatedTask, err := suite.repository.UpdateTask(task, task.ID.Hex())
	suite.NoError(err, "no error when update task with valid input")
	suite.NotNil(updatedTask, "updated task is not nil")
	suite.Equal(task.Description, updatedTask.Description, "updated task description is equal to the new description")

}

func Test_taskRepositorySuite(t *testing.T) {
	/// we still need this to run all tests in our suite
	suite.Run(t, &taskRepositorySuite{})
}
