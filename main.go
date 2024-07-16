package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	httpServer := gin.Default()

	// GET request on /events
	httpServer.GET("/events", getEvents)

	// running http server on port 8080
	httpServer.Run(":8080")
}

func getEvents(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
