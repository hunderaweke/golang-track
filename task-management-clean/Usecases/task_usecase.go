package usecases

import (
	domain "clean-architecture/Domain"
	"context"
	"time"
)

type taskUsecase struct {
	taskRepository  domain.TaskRepository
	timeoutDuration time.Duration
	ctx             context.Context
}

func NewTaskUseCase(taskRepository domain.TaskRepository, timeout time.Duration, c context.Context) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository:  taskRepository,
		timeoutDuration: timeout,
		ctx:             c,
	}
}

func (tu *taskUsecase) Create(t domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(tu.ctx, tu.timeoutDuration)
	defer cancel()
	return tu.taskRepository.Create(ctx, t)
}

func (tu *taskUsecase) Get() ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(tu.ctx, tu.timeoutDuration)
	defer cancel()
	return tu.taskRepository.Get(ctx)
}

func (tu *taskUsecase) GetByID(taskID string) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(tu.ctx, tu.timeoutDuration)
	defer cancel()
	return tu.taskRepository.GetByID(ctx, taskID)
}

func (tu *taskUsecase) GetByUserID(userID string) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(tu.ctx, tu.timeoutDuration)
	defer cancel()
	return tu.taskRepository.GetByUserID(ctx, userID)
}

func (tu *taskUsecase) Update(taskID string, data domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(tu.ctx, tu.timeoutDuration)
	defer cancel()
	return tu.taskRepository.Update(ctx, taskID, data)
}

func (tu *taskUsecase) Delete(taskID string) error {
	ctx, cancel := context.WithTimeout(tu.ctx, tu.timeoutDuration)
	defer cancel()
	return tu.taskRepository.Delete(ctx, taskID)
}
