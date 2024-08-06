package router

import (
	"context"
	"task-management-api-mongodb/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTaskRouter(r *gin.Engine, db *mongo.Database) {
	t := controllers.NewTaskController(context.TODO(), db)
	r.GET("/tasks/", t.GetTasks)
	r.GET("/tasks/:id", t.GetTaskByID)
	r.PUT("/tasks/:id", t.UpdateTask)
	r.DELETE("/tasks/:id", t.DeleteTask)
	r.POST("/tasks/", t.CreateTask)
}

func AddUserRouter(r *gin.Engine, db *mongo.Database) {
	u := controllers.NewUserController(db)
	r.GET("/users/", u.GetUsers)
	r.POST("/register", u.Create)
	r.POST("/login", u.Login)
}
