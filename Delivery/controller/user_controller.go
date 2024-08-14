package controller

import (
	"net/http"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/gin-gonic/gin"
)

type User = domain.User

type userController struct {
	userUsecase domain.UserUseCase
}

func NewUserController(userUsecase domain.UserUseCase) domain.UserController {
	return &userController{userUsecase: userUsecase}

}

func (uc *userController) SignUp(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Invalid Content"})
		return
	}

	newuser, errr := uc.userUsecase.CreateUser(user)

	if errr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newuser})
}

func (uc *userController) LoginUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Invalid Content"})
		return
	}

	token, errr := uc.userUsecase.LoginUser(user)

	if errr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	// create "authorithetion key with token cocke"

	c.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{"Authorization": token})

}

func (uc *userController) PromoteUser(c *gin.Context) {
	userId := c.Param("id")

	newuser, errr := uc.userUsecase.PromoteUser(userId)

	if errr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newuser})
}
