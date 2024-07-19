package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

// var events []Event = []Event{}

func (event Event) Save(db *sql.DB) (int64, error) {

	insertEventQuery := "INSERT INTO events (name, description, location, dateTime, userId) VALUES (?, ?, ?, ?, ?)"
	stmt, _prepareErr := db.Prepare(insertEventQuery)

	if _prepareErr != nil {
		fmt.Println(_prepareErr)
		panic("unable to prepare provided sql statement")
	}

	result, _stmtErr := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if _stmtErr != nil {
		fmt.Println(_stmtErr)
		panic("unable to execute sql query")
	}

	defer stmt.Close()

	return result.LastInsertId()
}

func GetAllEvents(db *sql.DB) []Event {
	var events []Event

	selectAllEvents := "SELECT id, name, description, location, userId, dateTime FROM events"
	rows, _queryErr := db.Query(selectAllEvents)

	if _queryErr != nil {
		fmt.Print("unable to execute select query")
	}

	for rows.Next() {
		var event Event
		// = Event{}
		_err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.UserID,
			&event.DateTime)

		if _err != nil {
			fmt.Println(_err)
			fmt.Println("unable to map values of event")
		}
		events = append(events, event)
	}
	defer rows.Close()
	return events
}
