package routes

import (
	"database/sql"
	"fmt"
	"time"

	"booking.com/models"
)

func GetEvents(db *sql.DB) []models.Event {

	// context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	events := models.GetAllEvents(db)
	return events
}

func CreateEvents(event models.Event, db *sql.DB) int64 {
	event.UserID = 1
	event.DateTime = time.Now()

	id, _err := event.Save(db)

	if _err != nil {
		fmt.Println(_err)
		panic("unable to save event")
	}

	return id
}
