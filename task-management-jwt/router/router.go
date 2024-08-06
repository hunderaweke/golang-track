package router

import (
	"context"
	"task-management-api-mongodb/controllers"
	"task-management-api-mongodb/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTaskRouter(r *gin.Engine, db *mongo.Database) {
	t := controllers.NewTaskController(context.TODO(), db)
	tasksGroup := r.Group("/tasks")
	tasksGroup.Use(middlewares.JWTMiddleware())
	{
		tasksGroup.GET("/", middlewares.JWTMiddleware(), t.GetTasks)
		tasksGroup.GET("/:id", t.GetTaskByID)
		tasksGroup.PUT("/:id", t.UpdateTask)
		tasksGroup.DELETE("/:id", t.DeleteTask)
		tasksGroup.POST("/", t.CreateTask)
	}
}

func AddUserRouter(r *gin.Engine, db *mongo.Database) {
	u := controllers.NewUserController(db)
	r.GET("/users/", u.GetUsers)
	r.POST("/register", u.Create)
	r.POST("/login", u.Login)
}
