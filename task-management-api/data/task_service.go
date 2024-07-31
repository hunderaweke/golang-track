package data

import (
	"errors"
	"task-management-api/models"
)

var TaskNotFoundError = errors.New("task not found")

type TasksService struct {
	Tasks map[string]models.Task
}

func (t *TasksService) GetTaskByID(taskID string) (models.Task, error) {
	task, ok := t.Tasks[taskID]
	if !ok {
		return task, TaskNotFoundError
	}
	return task, nil
}

func (t *TasksService) AddTask(task models.Task) {
	t.Tasks[task.ID] = task
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
	t.Tasks[taskID] = task
	return nil
}

func (t *TasksService) DeleteTask(taskID string) error {
	_, ok := t.Tasks[taskID]
	if !ok {
		return TaskNotFoundError
	}
	delete(t.Tasks, taskID)
	return nil
}
