package repositories

import (
	"context"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/usecase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task = domain.Task

type taskRepository struct {
	db             mongo.Database
	collectionName string
}

func NewTaskRepository(db mongo.Database, collectionName string) usecase.TaskRepository {

	return &taskRepository{db: db, collectionName: collectionName}
}

func (tr *taskRepository) GetTasks() ([]domain.Task, error) {

	collection := tr.db.Collection(tr.collectionName)

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

func (tr *taskRepository) GetTaskById(id string) (domain.Task, error) {
	collection := tr.db.Collection(tr.collectionName)
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

func (tr *taskRepository) CreateTask(task domain.Task) error {
	collection := tr.db.Collection(tr.collectionName)
	_, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) UpdateTask(task domain.Task, id string) (domain.Task, error) {
	collection := tr.db.Collection(tr.collectionName)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Task{}, err
	}

	filter := bson.D{{"_id", objectId}}

	update := bson.D{{"$set", task}}

	_, errr := collection.UpdateOne(context.TODO(), filter, update)
	if errr != nil {
		return Task{}, errr
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (tr *taskRepository) DeleteTask(id string) error {
	collection := tr.db.Collection(tr.collectionName)

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objectId}}
	_, errr := collection.DeleteOne(context.TODO(), filter)
	if errr != nil {
		return errr
	}

	return nil

}
