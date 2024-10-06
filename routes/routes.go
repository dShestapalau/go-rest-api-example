package routes

import (
	"github.com/gin-gonic/gin"
	"shestapalau.by/rest/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/signUp", signUp)
	server.POST("/login", login)

}
