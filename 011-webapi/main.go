package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"leroy.africa/leroy/webapi/models"
)

func main() {
	server := gin.Default()

	server.GET("/ping", ping)

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not process request data. " + err.Error(),
		})
		return
	}

	event.ID = 1
	event.UserID = 1
	
	event.Save()
	
	context.JSON(http.StatusCreated, event)
}