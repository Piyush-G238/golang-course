package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"booking.com/models"
	"github.com/gin-gonic/gin"

	"booking.com/db"
)

var DB *sql.DB

func main() {

	// initialize database connection
	DB = db.InitDB()

	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	httpServer := gin.Default()

	// GET request on /events
	httpServer.GET("/events", getEvents)
	httpServer.POST("/events", createEvents)

	// running http server on port 8080
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
