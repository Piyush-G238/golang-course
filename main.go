package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"booking.com/models"
	"github.com/gin-gonic/gin"

	"booking.com/db"
	"booking.com/routes"
)

var DB *sql.DB

func main() {

	// initialize database connection
	DB = db.InitDB()

	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	httpServer := gin.Default()

	routes.RegisterRoutes(httpServer, DB)

	/*
		// GET request on /events
		httpServer.GET("/events", getEvents)
		httpServer.POST("/events", createEvents)
		httpServer.GET("/events/:eventId", getEventById)
		httpServer.PUT("/events/:eventId", updateEvent)
		httpServer.DELETE("/events/:eventId", deleteEvent)
		// running http server on port 8080
	*/

	httpServer.Run(":8080")
}

func getEvents(context *gin.Context) {

	// context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	events := models.GetAllEvents(DB)
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	_error := context.ShouldBindJSON(&event)
	if _error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}
	// event.ID = 1
	event.UserID = 1
	event.DateTime = time.Now()

	id, _err := event.Save(DB)

	if _err != nil {

		fmt.Println(_err)
		panic("unable to save event")
	}
	context.JSON(http.StatusCreated, gin.H{"Generated ID": id})
}

func getEventById(context *gin.Context) {
	eventId, parseErr := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if parseErr != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "unable to parse eventId"})
		return
	}

	event := models.GetEventByID(eventId, DB)
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {

	eventId, parseError := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if parseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse event Id"})
		return
	}

	fetchedEvent := models.GetEventByID(eventId, DB)
	if fetchedEvent.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found with provided event ID"})
		return
	}

	var event models.Event
	context.BindJSON(&event)

	event.ID = int(eventId)
	event.DateTime = time.Now()
	event.UserID = 1

	_, _updateError := event.Update(DB)
	if _updateError != nil {
		fmt.Println(_updateError)
		context.JSON(http.StatusBadRequest, gin.H{"message": "unable to update event"})
	}

	context.JSON(http.StatusOK, gin.H{"Event ID": event.ID})
}

func deleteEvent(context *gin.Context) {

	eventId, _parseError := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if _parseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse eventId"})
		return
	}

	models.DeleteEventByID(eventId, DB)
	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully!"})
}
