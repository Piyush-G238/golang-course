package models

import (
	"database/sql"
	"errors"

	"booking.com/utils"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Password  string
}

func (user User) Save(db *sql.DB) (int64, error) {

	createUserQuery := "INSERT INTO users (firstName, lastName, username, password) values(?, ?, ?, ?)"
	stmt, _prepareError := db.Prepare(createUserQuery)

	if _prepareError != nil {
		panic("unable to generate create user query")
	}

	utils.HashPassword(&user.Password)

	result, _execError := stmt.Exec(user.FirstName, user.LastName, user.UserName, user.Password)
	if _execError != nil {
		panic("unable to execute create user query")
	}

	return result.LastInsertId()
}

func (user User) ValidateCredentials(db *sql.DB) error {

	query := "SELECT password FROM users WHERE username = ?"
	row := db.QueryRow(query, user.UserName)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return err
	}

	var credentialMatches bool = utils.CheckPasswordHash(
		user.Password,
		retrievedPassword)

	if !credentialMatches {
		return errors.New("credentials invalid")
	}

	return nil
}
