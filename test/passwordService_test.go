package test_test

import (
	"testing"

	infrastructure "github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/Infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestHashePassword(t *testing.T) {
	service := infrastructure.NewPasswordInfrastructureService()

	// Test case: valid password
	password := "123456"
	hashedPassword, err := service.HashePassword(password)
	assert.Nil(t, err)
	assert.NotEqual(t, password, hashedPassword)

	// Test case: check that hashed password can be compared correctly
	isSame := service.ComparePassword(hashedPassword, password)
	assert.True(t, isSame)
}

func TestComparePassword(t *testing.T) {
	service := infrastructure.NewPasswordInfrastructureService()

	// Test case: valid password
	password := "123456"
	hashedPassword, _ := service.HashePassword(password)

	// Test case: correct password comparison
	isSame := service.ComparePassword(hashedPassword, password)
	assert.True(t, isSame)

	// Test case: incorrect password comparison
	isSame = service.ComparePassword(hashedPassword, "wrongPassword")
	assert.False(t, isSame)
}
