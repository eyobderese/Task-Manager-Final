package test_test

import (
	"testing"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/stretchr/testify/suite"
)

type domainTestSuit struct {
	suite.Suite
}

func (suite *domainTestSuit) SetupSuite() {

}

func (suit *domainTestSuit) Test_userEntity() {
	// the entitiy we need to test

	user := domain.User{
		Email:    "test",
		Password: "testPassword",
	}

	suit.Equal("test", user.Email)
	suit.Equal("testPassword", user.Password)

}

func (suit *domainTestSuit) Test_TaskEntity() {
	task := domain.Task{
		Title:       "titile1",
		Description: "this is the first task",
	}

	suit.Equal("titile1", task.Title)
	suit.Equal("this is the first task", task.Description)
	// test the DueDate is exist
	suit.NotNil(task.DueDate)
}

func Test_domainSuite(t *testing.T) {
	/// we still need this to run all tests in our suite
	suite.Run(t, &domainTestSuit{})
}
