package main

import (
	"task-management-api-mongodb/router"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	router.AddRouter(r)
	r.Run(":7070")
}
