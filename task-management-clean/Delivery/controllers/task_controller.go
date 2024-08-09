package controllers

import (
	domain "clean-architecture/Domain"
	infrastructure "clean-architecture/Infrastructure"
	repository "clean-architecture/Repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUsecase domain.TaskUsecase
}

func NewTaskController(taskUsecase domain.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase: taskUsecase}
}

func getUserClaims(c *gin.Context) infrastructure.UserClaims {
	claims, _ := c.Get("claims")
	userClaims := claims.(*infrastructure.UserClaims)
	return *userClaims
}

func (t *TaskController) GetTasks(c *gin.Context) {
	tasks, err := t.taskUsecase.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (t *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := t.taskUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	userClaims := getUserClaims(c)
	if !userClaims.IsAdmin {
		c.JSON(http.StatusNotModified, gin.H{"error": "updating task requires admin access"})
		return
	}
	var updatedTask domain.Task
	err := c.BindJSON(&updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	task, err := t.taskUsecase.Update(taskID, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	userClaims := getUserClaims(c)
	if !userClaims.IsAdmin {
		c.JSON(http.StatusNotModified, gin.H{"error": "deleting task requires admin access"})
		return
	}
	_, err := t.taskUsecase.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": repository.TaskNotFoundError})
		return
	}
	t.taskUsecase.Delete(taskID)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if task.DueDate.IsZero() {
		task.DueDate = time.Now()
	}
	task, err := t.taskUsecase.Create(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusCreated, task)
}
