package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     time.Time          `json:"due_date"`
	Status      string             `json:"status"`
}

const (
	TaskCollection = "tasks"
)

type TaskUsecase interface {
	GetTasks() ([]Task, error)
	GetTaskById(id string) (Task, error)
	CreateTask(task Task) error
	UpdateTask(task Task, id string) (Task, error)
	DeleteTask(id string) error
}



type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `json:"email" `
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

type UserUseCase interface {
	CreateUser(user User) (User, error)
	LoginUser(user User) (string, error)
	PromoteUser(userId string) (User, error)
}

type UserController interface {
	SignUp(c *gin.Context)
	LoginUser(c *gin.Context)
	PromoteUser(c *gin.Context)
}
