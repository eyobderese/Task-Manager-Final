package test_test

import (
	"testing"
	"time"

	infrastructure "github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/Infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestTokenGeneretor(t *testing.T) {
	service := infrastructure.NewJwtService()

	// Test case: valid claims
	claims := map[string]interface{}{
		"username": "testuser",
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token, err := service.TokenGeneretor(claims)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestTotokenParser(t *testing.T) {
	service := infrastructure.NewJwtService()

	// Test case: valid token
	claims := map[string]interface{}{
		"username": "testuser",
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token, _ := service.TokenGeneretor(claims)
	parsedClaims, err := service.TotokenParser(token)
	assert.Nil(t, err)
	assert.Equal(t, claims["username"], parsedClaims["username"])
}
