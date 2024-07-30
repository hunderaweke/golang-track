package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/tasks/", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.POST("/tasks/", createTask)
	router.Run(":8080")
}

func getTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func getTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, t := range tasks {
		if t.ID == id {
			ctx.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "object not found",
	})
}

func updateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var t Task
	if err := ctx.BindJSON(&t); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	for i, task := range tasks {
		if task.ID == id {
			if t.Title != "" {
				tasks[i].Title = t.Title
			}
			if t.Description != "" {
				tasks[i].Description = t.Description
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "task updated"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func createTask(ctx *gin.Context) {
	var t Task
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	tasks = append(tasks, t)
	ctx.JSON(http.StatusOK, gin.H{"message": "task added"})
}
