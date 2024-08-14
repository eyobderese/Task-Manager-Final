package usecase

import (

	// "time"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
)

type Task = domain.Task

// Mock data for tasks
//
//	var tasks = []Task{
//		{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
//		{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
//		{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
//	}
type TaskRepository interface {
	GetTasks() ([]Task, error)
	GetTaskById(id string) (Task, error)
	CreateTask(task Task) error
	UpdateTask(task Task, id string) (Task, error)
	DeleteTask(id string) error
}

type taskUsecase struct {
	taskRepository TaskRepository
}

func NewTaskUsecase(taskRepository TaskRepository) domain.TaskUsecase {
	return &taskUsecase{taskRepository: taskRepository}
}

func (tu *taskUsecase) GetTasks() ([]Task, error) {
	return tu.taskRepository.GetTasks()
}

func (tu *taskUsecase) GetTaskById(id string) (Task, error) {
	return tu.taskRepository.GetTaskById(id)
}

func (tu *taskUsecase) CreateTask(task Task) error {
	err := tu.taskRepository.CreateTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (tu *taskUsecase) UpdateTask(task Task, id string) (Task, error) {
	return tu.taskRepository.UpdateTask(task, id)
}

func (tu *taskUsecase) DeleteTask(id string) error {
	err := tu.taskRepository.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
}

/*


// GetTasks returns a slice of tasks.
func GetTasks() ([]Task, error) {
	collection := getCollection()

	var tasks []Task
	cur, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return tasks, err
	}

	for cur.Next(context.TODO()) {
		var task Task
		err := cur.Decode(&task)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	cur.Close(context.TODO())

	return tasks, nil

}

// GetTaskById retrieves a task from the tasks slice based on the provided ID.
// If a task with the given ID is found, it is returned. Otherwise, an empty task is returned.
func GetTaskById(id string) (Task, error) {
	collection := getCollection()
	var task Task
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return Task{}, err
	}

	filter := bson.D{{"_id", objectId}}

	err = collection.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

// CreateTask creates a new task and adds it to the list of tasks.
// It takes a Task object as a parameter and returns the created task.
func CreateTask(task Task) (Task, error) {
	collection := getCollection()
	_, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

// UpdateTask updates a task with the given ID in the tasks slice.
// It replaces the existing task with the provided task and returns the updated task.
// If no task with the given ID is found, it returns an empty Task struct.
func UpdateTask(task Task, id string) (Task, error) {
	collection := getCollection()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Task{}, err
	}

	filter := bson.D{{"_id", objectId}}

	update := bson.D{{"$set", task}}

	updatedResult, errr := collection.UpdateOne(context.TODO(), filter, update)
	fmt.Println(updatedResult)
	if errr != nil {
		return Task{}, errr
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

// DeleteTask deletes a task with the specified ID from the tasks slice.
// It returns the deleted task if found, otherwise it returns an empty task.
func DeleteTask(id string) (Task, error) {
	collection := getCollection()
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return Task{}, err
	}

	filter := bson.D{{"_id", objectId}}
	_, errr := collection.DeleteOne(context.TODO(), filter)
	if errr != nil {
		return Task{}, errr
	}

	return Task{}, nil

}

/*

func EditTask(task Task, id string) Task {

	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		return Task{}
	}

	if task.ID != "" {
		tasks[index].ID = task.ID
	}
	if task.Title != "" {

		tasks[index].Title = task.Title
	}
	if task.Description != "" {
		tasks[index].Description = task.Description
	}
	if (task.DueDate != time.Time{}) {
		tasks[index].DueDate = task.DueDate
	}
	if task.Status != "" {
		tasks[index].Status = task.Status
	}
	return tasks[index]
}
*/
