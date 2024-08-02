package controllers

import (
	"net/http"
	"strconv"
	"task-management-api-mongodb/data"
	"task-management-api-mongodb/models"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	T      data.TasksService = *data.Connect()
	nextID int               = len(T.Tasks) + 1
)

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": T.Tasks})
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	t, err := T.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, t)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	err := c.BindJSON(&updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	T.UpdateTask(id, updatedTask)
	c.IndentedJSON(http.StatusOK, T.Tasks[id])
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	_, ok := T.Tasks[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": data.TaskNotFoundError})
		return
	}
	T.DeleteTask(id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func CreateTask(c *gin.Context) {
	var t models.Task
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if t.DueDate.IsZero() {
		t.DueDate = time.Now()
	}
	t.ID = strconv.Itoa(nextID)
	nextID++
	err := T.AddTask(t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusCreated, t)
}
