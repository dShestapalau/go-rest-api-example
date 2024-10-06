package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shestapalau.by/rest/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", GetEvents)
	server.POST("events", CreateEvent)

	server.Run(":8080")
}

func GetEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func CreateEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, event)
}
