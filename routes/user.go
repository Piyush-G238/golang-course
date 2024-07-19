package routes

import (
	"database/sql"

	"booking.com/models"
)

func CreateUser(user models.User, db *sql.DB) int64 {

	userId, _error := user.Save(db)
	if _error != nil {
		panic("unable to save user details")
	}
	return userId
}
