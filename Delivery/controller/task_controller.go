package controller

import (
	"net/http"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskUsecase domain.TaskUsecase
}

func NewTaskController(taskUsecase domain.TaskUsecase) *taskController {
	return &taskController{taskUsecase: taskUsecase}

}

// GetTasks retrieves all tasks from the data source and returns them as JSON.
func (tc *taskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskUsecase.GetTasks()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": tasks})
}

// GetTaskById retrieves a task by its ID.
// It takes a gin.Context object as a parameter and uses the ID parameter from the request path to fetch the task from the data package.
// If the task is found, it returns the task data as a JSON response with HTTP status code 200.
// If the task is not found, it returns a JSON response with HTTP status code 404 and a message indicating that the task was not found.

func (tc *taskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskUsecase.GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": task})
}

// CreateTask creates a new task based on the JSON data provided in the request body.
// It binds the JSON data to a Task struct, creates the task using the data, and returns the created task as JSON response.
func (tc *taskController) CreateTask(c *gin.Context) {
	var task domain.Task
	c.BindJSON(&task)

	err := tc.taskUsecase.CreateTask(task)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": task})
}

// UpdateTask updates a task with the given ID.
// It receives the task ID from the request parameters and the updated task data from the request body.
// It returns the updated task if successful, or a JSON response with an error message if the task is not found.
func (tc *taskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	c.BindJSON(&task)

	task, err := tc.taskUsecase.UpdateTask(task, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": task})
}

// DeleteTask deletes a task with the given ID.
// It takes a gin.Context object as a parameter and retrieves the task ID from the URL parameter.
// It then calls the data.DeleteTask function to delete the task with the given ID.
// If the task is not found, it returns a JSON response with a "Task not found" message and a status code of 404 (Not Found).
// If the task is successfully deleted, it returns a JSON response with the deleted task data and a status code of 200 (OK).
func (tc *taskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := tc.taskUsecase.DeleteTask(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "successfully deleted"})
}
