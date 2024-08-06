package main

import (
	"context"
	"log"
	"os"
	"task-management-api-mongodb/database"
	"task-management-api-mongodb/router"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	dbUri := os.Getenv("MONGODB_URL")
	clnt, err := database.NewConnection(context.TODO(), dbUri)
	if err != nil {
		log.Fatal(err)
	}
	db := clnt.Database("task_management_api")
	router.AddTaskRouter(r, db)
	router.AddUserRouter(r, db)
	r.Run(":7070")
}
