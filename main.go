package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"shestapalau.by/rest/db"
	"shestapalau.by/rest/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/events", CreateEvent)

	server.Run(":8080")
}

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	if len(events) == 0 {
		events = []models.Event{}
	}

	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func CreateEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	event.UserID = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	context.JSON(http.StatusCreated, event)
}
