package router

import (
	"restful-api-sample/middlewares"

	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {
	const routerBase string = "/events"
	server.GET(routerBase, getEvents)
	server.GET(routerBase+"/:id", getEvent)

	authenticated := server.Group("/events")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/", createEvent)
	authenticated.PUT("/:id", updateEvent)
	authenticated.DELETE("/:id", deleteEvent)
	authenticated.POST("/:id/register", registerForEvent)
	authenticated.DELETE("/:id/register", cancelRegisterForEvent)
}

func AuthRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
}
