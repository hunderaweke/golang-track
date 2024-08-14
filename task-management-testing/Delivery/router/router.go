package router

import (
	"context"
	"testing-api/Delivery/controllers"
	infrastructure "testing-api/Infrastructure"
	repository "testing-api/Repositories"
	usecases "testing-api/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sv-tools/mongoifc"
)

type taskRouter struct {
	controller controllers.TaskController
}
type userRouter struct {
	controller controllers.UserController
}

func SetupRouter(r *gin.Engine, db mongoifc.Database, timeOut time.Duration, ctx context.Context) {
	AddTaskRouter(r, db, timeOut, ctx)
	AddUserRouter(r, db, timeOut, ctx)
}

func NewTaskRouter(c context.Context, db mongoifc.Database, timeOut time.Duration) taskRouter {
	taskRepository := repository.NewTaskService(c, db)
	taskUsecase := usecases.NewTaskUseCase(taskRepository, timeOut, c)
	taskController := controllers.NewTaskController(taskUsecase)
	return taskRouter{
		controller: *taskController,
	}
}

func NewUserRouter(c context.Context, db mongoifc.Database, timeOut time.Duration) userRouter {
	userRepository := repository.NewUserService(c, db)
	userUsecase := usecases.NewUserUsecase(userRepository, timeOut, c)
	userController := controllers.NewUserController(userUsecase)
	return userRouter{
		controller: *userController,
	}
}

func AddTaskRouter(r *gin.Engine, db mongoifc.Database, timeOut time.Duration, c context.Context) {
	t := NewTaskRouter(c, db, timeOut)
	tasksGroup := r.Group("/tasks")
	tasksGroup.Use(infrastructure.JWTMiddleware())
	{
		tasksGroup.GET("/", t.controller.GetTasks)
		tasksGroup.GET("/:id", t.controller.GetTaskByID)
		tasksGroup.PUT("/:id", infrastructure.AdminMiddleware(), t.controller.UpdateTask)
		tasksGroup.DELETE("/:id", infrastructure.AdminMiddleware(), t.controller.DeleteTask)
		tasksGroup.POST("/", infrastructure.AdminMiddleware(), t.controller.CreateTask)
	}
}

func AddUserRouter(r *gin.Engine, db mongoifc.Database, timeOut time.Duration, c context.Context) {
	u := NewUserRouter(c, db, timeOut)
	admin := r.Group("/users/")
	admin.Use(infrastructure.JWTMiddleware())
	{
		admin.PUT("/promote", infrastructure.AdminMiddleware(), u.controller.PromoteUser)
		admin.GET("/", infrastructure.AdminMiddleware(), u.controller.GetUsers)
		admin.GET("/:id", u.controller.GetUserByID)
		admin.PUT("/:id", infrastructure.AdminMiddleware(), u.controller.UpdateUser)
		admin.DELETE("/:id", infrastructure.AdminMiddleware(), u.controller.DeleteUser)
	}
	r.POST("/register", u.controller.Create)
	r.POST("/login", u.controller.Login)
}
