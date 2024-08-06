package controllers

import (
	"context"
	"net/http"
	"strconv"
	"task-management-api-mongodb/data"
	"task-management-api-mongodb/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskController struct {
	taskService data.TasksService
	nextID      int
}

func NewTaskController(c context.Context, db *mongo.Database) *TaskController {
	taskService := data.NewTaskService(c, db)
	return &TaskController{taskService: *taskService, nextID: taskService.Count + 1}
}

func (t *TaskController) GetTasks(c *gin.Context) {
	tasks, err := t.taskService.GetTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (t *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := t.taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	err := c.BindJSON(&updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	task, err := t.taskService.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := t.taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": data.TaskNotFoundError})
		return
	}
	t.taskService.DeleteTask(id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if task.DueDate.IsZero() {
		task.DueDate = time.Now()
	}
	task.ID = strconv.Itoa(t.nextID)
	t.nextID++
	err := t.taskService.AddTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusCreated, task)
}
