package usecases

import (
	"context"
	"testing"
	domain "testing-api/Domain"
	"testing-api/mocks"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskUsecaseTestSuite struct {
	suite.Suite
	taskRepository *mocks.TaskRepository
	taskUsecase    domain.TaskUsecase
}

func (suite *TaskUsecaseTestSuite) SetupSuite() {
	suite.taskRepository = mocks.NewTaskRepository(suite.T())
	suite.taskUsecase = NewTaskUseCase(suite.taskRepository, time.Duration(4*time.Second), context.TODO())
}

func (suite *TaskUsecaseTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	taskInput := domain.Task{
		ID:          "task123", // Manually set the ID
		UserID:      "user123",
		Title:       "Test Task",
		Description: "This is a test task.",
		DueDate:     time.Now(),
		Status:      "pending",
	}
	expectedTask := taskInput
	suite.taskRepository.On("Create", mock.Anything, taskInput).Return(taskInput, nil)
	createdTask, err := suite.taskUsecase.Create(taskInput)
	assert.NoError(err)
	assert.Equal(expectedTask, createdTask)
}

func (suite *TaskUsecaseTestSuite) TestGet() {
	assert := assert.New(suite.T())
	tasks := []domain.Task{
		{
			ID:          "task1",
			UserID:      "user123",
			Title:       "Complete project report",
			Description: "Finish the final report and submit it to the manager.",
			DueDate:     time.Date(2024, 8, 15, 12, 0, 0, 0, time.UTC),
			Status:      "pending",
		},
		{
			ID:          "task2",
			UserID:      "user123",
			Title:       "Prepare presentation",
			Description: "Create slides for the upcoming conference.",
			DueDate:     time.Date(2024, 8, 20, 9, 0, 0, 0, time.UTC),
			Status:      "pending",
		},
		{
			ID:          "task3",
			UserID:      "user456",
			Title:       "Update software documentation",
			Description: "Revise the user manual and update the API docs.",
			DueDate:     time.Date(2024, 8, 18, 17, 0, 0, 0, time.UTC),
			Status:      "done",
		},
	}
	suite.taskRepository.On("Get", mock.Anything).Return(tasks, nil)
	got, err := suite.taskUsecase.Get()
	assert.NoError(err)
	assert.Equal(tasks, got)
}

func (suite *TaskUsecaseTestSuite) TestGetByID() {
	assert := assert.New(suite.T())
	expectedTask := domain.Task{
		ID:          "task3",
		UserID:      "user456",
		Title:       "Update software documentation",
		Description: "Revise the user manual and update the API docs.",
		DueDate:     time.Date(2024, 8, 18, 17, 0, 0, 0, time.UTC),
		Status:      "done",
	}
	suite.taskRepository.On("GetByID", mock.Anything, expectedTask.ID).Return(expectedTask, nil)
	got, err := suite.taskUsecase.GetByID(expectedTask.ID)
	assert.NoError(err)
	assert.Equal(expectedTask, got)
}

func (suite *TaskUsecaseTestSuite) TestGetByUserID() {
	assert := assert.New(suite.T())
	userID := "user123"
	tasks := []domain.Task{
		{
			ID:          "task1",
			UserID:      userID,
			Title:       "Complete project report",
			Description: "Finish the final report and submit it to the manager.",
			DueDate:     time.Date(2024, 8, 15, 12, 0, 0, 0, time.UTC),
			Status:      "pending",
		},
		{
			ID:          "task2",
			UserID:      userID,
			Title:       "Prepare presentation",
			Description: "Create slides for the upcoming conference.",
			DueDate:     time.Date(2024, 8, 20, 9, 0, 0, 0, time.UTC),
			Status:      "pending",
		},
		{
			ID:          "task3",
			UserID:      userID,
			Title:       "Update software documentation",
			Description: "Revise the user manual and update the API docs.",
			DueDate:     time.Date(2024, 8, 18, 17, 0, 0, 0, time.UTC),
			Status:      "done",
		},
	}
	suite.taskRepository.On("GetByUserID", mock.Anything, userID).Return(tasks, nil)
	got, err := suite.taskUsecase.GetByUserID(userID)
	assert.NoError(err)
	assert.Equal(got, tasks)
}

func (suite *TaskUsecaseTestSuite) TestUpdate() {
	assert := assert.New(suite.T())
	task := domain.Task{
		ID:          "task3",
		UserID:      "userid",
		Title:       "Update software documentation",
		Description: "Revise the user manual and update the API docs.",
		DueDate:     time.Date(2024, 8, 18, 17, 0, 0, 0, time.UTC),
		Status:      "done",
	}
	expectedTask := task
	suite.taskRepository.On("Update", mock.Anything, task.ID, task).Return(&task, nil)
	updatedTask, err := suite.taskUsecase.Update(task.ID, task)
	assert.NoError(err)
	assert.Equal(expectedTask, *updatedTask)
}

func (suite *TaskUsecaseTestSuite) TestDelete() {
	assert := assert.New(suite.T())
	suite.taskRepository.On("Delete", mock.Anything, "taskID").Return(nil)
	err := suite.taskUsecase.Delete("taskID")
	assert.NoError(err)
}

func TestTaskUsecase(t *testing.T) {
	tSuite := new(TaskUsecaseTestSuite)
	suite.Run(t, tSuite)
	tSuite.taskRepository.AssertExpectations(t)
}
