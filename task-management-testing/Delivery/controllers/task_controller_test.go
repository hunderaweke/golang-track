package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	domain "testing-api/Domain"
	infrastructure "testing-api/Infrastructure"
	"testing-api/mocks"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskControllerTestSuite struct {
	suite.Suite
	router         *gin.Engine
	taskUsecase    mocks.TaskUsecase
	taskController TaskController
	tasks          []domain.Task
	admin          domain.User
}

func (suite *TaskControllerTestSuite) SetupSuite() {
	suite.router = gin.Default()
	os.Setenv("JWT_SECRET", "testsecret")
	suite.admin = domain.User{
		ID:      "1",
		IsAdmin: true,
		Email:   "admin@admin.com",
	}
	suite.taskUsecase = *mocks.NewTaskUsecase(suite.T())
	suite.taskController = *NewTaskController(&suite.taskUsecase)
	userID := "user123"
	suite.tasks = []domain.Task{
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
}

func (suite *TaskControllerTestSuite) TestCreateController() {
	assert := assert.New(suite.T())
	for _, t := range suite.tasks {
		suite.taskUsecase.On("Create", t).Return(t, nil)
	}
	suite.router.POST("/tasks", infrastructure.JWTMiddleware(), infrastructure.AdminMiddleware(), suite.taskController.CreateTask)
	token, _ := infrastructure.GenerateToken(suite.admin)
	for _, t := range suite.tasks {
		data, _ := json.Marshal(t)
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(data))
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(http.StatusCreated, w.Code)
		assert.JSONEq(string(data), w.Body.String())
	}
}

func (suite *TaskControllerTestSuite) TestGetTasks() {
	assert := assert.New(suite.T())
	suite.taskUsecase.On("Get").Return(suite.tasks, nil)
	suite.router.GET("/tasks", infrastructure.JWTMiddleware(), suite.taskController.GetTasks)
	token, _ := infrastructure.GenerateToken(suite.admin)
	data, _ := json.Marshal(suite.tasks)
	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)
	assert.JSONEq(string(data), w.Body.String())
}

func (suite *TaskControllerTestSuite) TestGetTaskByID() {
	assert := assert.New(suite.T())
	for _, t := range suite.tasks {
		suite.taskUsecase.On("GetByID", t.ID).Return(t, nil)
	}
	suite.router.GET("/tasks/:id", infrastructure.JWTMiddleware(), suite.taskController.GetTaskByID)
	token, _ := infrastructure.GenerateToken(suite.admin)
	for _, t := range suite.tasks {
		data, _ := json.Marshal(t)
		req, _ := http.NewRequest("GET", "/tasks/"+t.ID, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(http.StatusOK, w.Code)
		assert.JSONEq(string(data), w.Body.String())
	}
}

func (suite *TaskControllerTestSuite) TestUpdateTask() {
	assert := assert.New(suite.T())
	for _, t := range suite.tasks {
		suite.taskUsecase.On("Update", t.ID, t).Return(t, nil)
		suite.taskUsecase.On("GetByID", t.ID).Return(t, nil)
	}
	suite.router.PUT("/tasks/:id", infrastructure.JWTMiddleware(), infrastructure.AdminMiddleware(), suite.taskController.GetTaskByID)
	token, _ := infrastructure.GenerateToken(suite.admin)
	for _, t := range suite.tasks {
		data, _ := json.Marshal(t)
		req, _ := http.NewRequest("PUT", "/tasks/"+t.ID, bytes.NewBuffer(data))
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(http.StatusOK, w.Code)
		assert.JSONEq(string(data), w.Body.String())
	}
}

func (suite *TaskControllerTestSuite) TestDeleteTask() {
	assert := assert.New(suite.T())
	for _, t := range suite.tasks {
		suite.taskUsecase.On("Delete", t.ID).Return(nil)
		suite.taskUsecase.On("GetByID", t.ID).Return(t, nil)
	}
	suite.router.DELETE("/tasks/:id", infrastructure.JWTMiddleware(), infrastructure.AdminMiddleware(), suite.taskController.DeleteTask)
	token, _ := infrastructure.GenerateToken(suite.admin)
	for _, t := range suite.tasks {
		req, _ := http.NewRequest("DELETE", "/tasks/"+t.ID, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(http.StatusNoContent, w.Code)
	}
}

func TestTaskController(t *testing.T) {
	tSuite := new(TaskControllerTestSuite)
	suite.Run(t, tSuite)
}
