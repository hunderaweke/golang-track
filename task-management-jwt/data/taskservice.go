package data

import (
	"context"
	"errors"
	"fmt"
	"task-management-api-mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskNotFoundError = errors.New("task not found")

type TasksService struct {
	collection *mongo.Collection
	Count      int
}

func NewTaskService(c context.Context, db *mongo.Database) *TasksService {
	collection := db.Collection("tasks")
	t := TasksService{collection: collection}
	count, _ := t.collection.CountDocuments(context.Background(), bson.D{{}}, options.Count())
	t.Count = int(count)
	return &t
}

func (t *TasksService) GetTasks() ([]models.Task, error) {
	opts := options.Find()
	tasks := []models.Task{}
	c, err := t.collection.Find(context.TODO(), bson.D{{}}, opts)
	if err != nil {
		return tasks, err
	}
	for c.Next(context.TODO()) {
		var e models.Task
		err = c.Decode(&e)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, e)
	}
	defer c.Close(context.TODO())
	return tasks, nil
}

func (t *TasksService) GetTaskByID(userID, taskID string) (models.Task, error) {
	tasks, err := t.GetTaskByUserID(userID)
	if err != nil {
		return models.Task{}, err
	}
	for _, t := range tasks {
		if t.ID == taskID {
			return t, nil
		}
	}
	return models.Task{}, fmt.Errorf("the task doesn't belong to the current user")
}

func (t *TasksService) GetTaskByUserID(userID string) ([]models.Task, error) {
	filter := bson.D{{"user_id", userID}}
	opts := options.Find()
	cursor, err := t.collection.Find(context.Background(), filter, opts)
	var tasks []models.Task
	if err != nil {
		return tasks, err
	}
	for cursor.Next(context.TODO()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *TasksService) AddTask(task models.Task) error {
	_, err := t.collection.InsertOne(context.Background(), task)
	return err
}

func (t *TasksService) UpdateTask(userID, taskID string, data models.Task) (models.Task, error) {
	task, err := t.GetTaskByID(userID, taskID)
	if err != nil {
		return task, err
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
	updateResult, err := t.collection.ReplaceOne(context.Background(), filter, task)
	if err != nil {
		return task, err
	}
	if updateResult.ModifiedCount != 1 {
		return task, fmt.Errorf("modification error modified %v", updateResult.ModifiedCount)
	}
	return task, nil
}

func (t *TasksService) DeleteTask(taskID string) error {
	filter := bson.D{{"id", taskID}}
	opts := options.Delete()
	res, err := t.collection.DeleteOne(context.Background(), filter, opts)
	if res.DeletedCount == 0 {
		return TaskNotFoundError
	}
	return err
}
