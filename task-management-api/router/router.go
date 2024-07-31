package router

import (
	"task-management-api/controllers"

	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.Engine) {
	r.GET("/tasks/", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTaskByID)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	r.POST("/tasks/", controllers.CreateTask)
}
