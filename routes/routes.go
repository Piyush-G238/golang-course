package routes

import (
	"database/sql"
	"net/http"

	"booking.com/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, db *sql.DB) {

	server.GET("/events", func(ctx *gin.Context) {
		events := GetEvents(db)
		ctx.JSON(http.StatusOK, events)
	})

	server.POST("/events", func(ctx *gin.Context) {
		var event models.Event
		_error := ctx.ShouldBindJSON(&event)

		if _error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
			return
		}

		id := CreateEvents(event, db)
		ctx.JSON(http.StatusCreated, gin.H{"Event ID": id})

	})

	server.POST("/signup", func(ctx *gin.Context) {
		var user models.User

		_parseError := ctx.ShouldBindJSON(&user)

		if _parseError != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "unable to parse request body"})
			return
		}

		userId := CreateUser(user, db)

		ctx.JSON(http.StatusCreated, gin.H{"User ID": userId})
	})

	server.POST("/login", func(ctx *gin.Context) {
		var user models.User

		_parseError := ctx.ShouldBindJSON(&user)

		if _parseError != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse request body"})
			return
		}

		err := user.ValidateCredentials(db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "User Authenticated"})

	})
	// server.GET("/events/:eventId", getEventById)
	// server.PUT("/events/:eventId", updateEvent)
	// server.DELETE("/events/:eventId", deleteEvent)
}
