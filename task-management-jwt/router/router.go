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
		tasksGroup.GET("/", t.GetTasks)
		tasksGroup.GET("/:id", t.GetTaskByID)
		tasksGroup.PUT("/:id", middlewares.AdminMiddleware(), t.UpdateTask)
		tasksGroup.DELETE("/:id", middlewares.AdminMiddleware(), t.DeleteTask)
		tasksGroup.POST("/", middlewares.AdminMiddleware(), t.CreateTask)
	}
}

func AddUserRouter(r *gin.Engine, db *mongo.Database) {
	u := controllers.NewUserController(db)
	admin := r.Group("/users/")
	admin.Use(middlewares.JWTMiddleware())
	{
		admin.PUT("/promote", middlewares.AdminMiddleware(), u.PromoteUser)
		admin.GET("/", middlewares.AdminMiddleware(), u.GetUsers)
		admin.GET("/:id", u.GetUserByID)
		admin.PUT("/:id", u.UpdateUser)
		admin.DELETE("/:id", middlewares.AdminMiddleware(), u.DeleteUser)
	}
	r.POST("/register", u.Create)
	r.POST("/login", u.Login)
}
