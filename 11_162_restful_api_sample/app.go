package main

import (
	"restful-api-sample/db"
	"restful-api-sample/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	router.EventRoutes(server)
	router.AuthRoutes(server)

	server.Run(":8080")
}
