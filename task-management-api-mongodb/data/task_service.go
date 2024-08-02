package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"task-management-api-mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskNotFoundError = errors.New("task not found")

type TasksService struct {
	Tasks      map[string]models.Task
	Collection *mongo.Collection
}

func Connect() *TasksService {
	dbUri := os.Getenv("MONGODB_URL")
	clientOptions := options.Client().ApplyURI(dbUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("task_management_api").Collection("tasks")
	t := TasksService{Tasks: make(map[string]models.Task), Collection: collection}
	o := options.Find()
	c, err := collection.Find(context.Background(), bson.D{{}}, o)
	if err != nil {
		log.Fatal(err)
	}
	for c.Next(context.TODO()) {
		var e models.Task
		err = c.Decode(&e)
		t.AddTask(e)
	}
	defer c.Close(context.TODO())
	return &t
}

func (t *TasksService) GetTaskByID(taskID string) (models.Task, error) {
	task, ok := t.Tasks[taskID]
	if !ok {
		return task, TaskNotFoundError
	}
	return task, nil
}

func (t *TasksService) AddTask(task models.Task) error {
	t.Tasks[task.ID] = task
	_, err := t.Collection.InsertOne(context.Background(), task)
	return err
}

func (t *TasksService) UpdateTask(taskID string, data models.Task) error {
	task, ok := t.Tasks[taskID]
	if !ok {
		return TaskNotFoundError
	}
	if data.Title != "" {
		task.Title = data.Title
	}
	if data.Description != "" {
		task.Description = data.Description
	}
	if !data.DueDate.IsZero() {
		task.DueDate = data.DueDate
	}
	filter := bson.D{{"id", taskID}}
	updateResult, err := t.Collection.ReplaceOne(context.Background(), filter, task)
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount != 1 {
		return fmt.Errorf("modification error modified %v", updateResult.ModifiedCount)
	}
	t.Tasks[taskID] = task
	return nil
}

func (t *TasksService) DeleteTask(taskID string) error {
	_, ok := t.Tasks[taskID]
	if !ok {
		return TaskNotFoundError
	}
	delete(t.Tasks, taskID)
	filter := bson.D{{"id", taskID}}
	_, err := t.Collection.DeleteOne(context.Background(), filter)
	return err
}
