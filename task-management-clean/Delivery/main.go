package main

import (
	"clean-architecture/Delivery/router"
	"clean-architecture/database"
	"context"
	"log"
	"os"
	"time"

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
	ctx := context.Background()
	timeOut := time.Duration(4 * time.Second)
	router.SetupRouter(r, db, timeOut, ctx)
	r.Run(":7070")
}
