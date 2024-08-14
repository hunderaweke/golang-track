package main

import (
	"context"
	"log"
	"testing-api/Delivery/router"
	"testing-api/database"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	clnt, err := database.NewMongoClient(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	db := clnt.Database("task_management_api")
	ctx := context.Background()
	timeOut := time.Duration(4 * time.Second)
	router.SetupRouter(r, db, timeOut, ctx)
	r.Run(":7070")
}
