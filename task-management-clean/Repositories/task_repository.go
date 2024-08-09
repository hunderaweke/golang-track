package repository

import (
	domain "clean-architecture/Domain"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskNotFoundError = errors.New("task not found")

type TasksRepository struct {
	collection *mongo.Collection
}

func NewTaskService(c context.Context, db *mongo.Database) domain.TaskRepository {
	collection := db.Collection(domain.TaskCollection)
	return &TasksRepository{collection: collection}
}

func (t *TasksRepository) Get(c context.Context) ([]domain.Task, error) {
	opts := options.Find()
	tasks := []domain.Task{}
	cursor, err := t.collection.Find(c, bson.D{{}}, opts)
	if err != nil {
		return tasks, err
	}
	for cursor.Next(context.TODO()) {
		var e domain.Task
		err = cursor.Decode(&e)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, e)
	}
	defer cursor.Close(c)
	return tasks, nil
}

func (t *TasksRepository) GetByID(c context.Context, taskID string) (domain.Task, error) {
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("task with id %v not found", taskID)
	}
	res := t.collection.FindOne(c, bson.D{{"_id", id}}, options.FindOne())
	var task domain.Task
	if err := res.Decode(&task); err != nil {
		return domain.Task{}, fmt.Errorf("task with id %v not found", taskID)
	}
	return task, nil
}

func (t *TasksRepository) GetByUserID(c context.Context, userID string) ([]domain.Task, error) {
	filter := bson.D{{"user_id", userID}}
	opts := options.Find()
	cursor, err := t.collection.Find(c, filter, opts)
	var tasks []domain.Task
	if err != nil {
		return tasks, err
	}
	for cursor.Next(context.TODO()) {
		var task domain.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *TasksRepository) Create(c context.Context, task domain.Task) (domain.Task, error) {
	task.ID = string(primitive.NewObjectIDFromTimestamp(time.Now()).Hex())
	_, err := t.collection.InsertOne(c, task)
	return task, err
}

func (t *TasksRepository) Update(c context.Context, taskID string, data domain.Task) (*domain.Task, error) {
	task, err := t.GetByID(c, taskID)
	if err != nil {
		return &task, err
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
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return &domain.Task{}, err
	}
	filter := bson.D{{"_id", id}}
	updateResult, err := t.collection.ReplaceOne(c, filter, task)
	if err != nil {
		return &task, err
	}
	if updateResult.ModifiedCount != 1 {
		return &task, fmt.Errorf("modification error modified %v", updateResult.ModifiedCount)
	}
	return &task, nil
}

func (t *TasksRepository) Delete(c context.Context, taskID string) error {
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", id}}
	opts := options.Delete()
	res, err := t.collection.DeleteOne(c, filter, opts)
	if res.DeletedCount == 0 {
		return TaskNotFoundError
	}
	return err
}
