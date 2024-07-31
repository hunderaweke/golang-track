package main

import (
	"task-management-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.AddRouter(r)
	r.Run(":7070")
}
