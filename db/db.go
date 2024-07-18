package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	db, _err := sql.Open("mysql", "root:hH&qeV%y12@tcp(localhost:3306)/golang_dev")

	if _err != nil {
		fmt.Println(_err)
		panic("unable to connect with given database")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	createEventTable := `
CREATE TABLE IF NOT EXISTS events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    location VARCHAR(255) NOT NULL,
    registration_time DATETIME NOT NULL,
    user_id INT NOT NULL
);`
	_, _queryErr := db.Exec(createEventTable)

	if _queryErr != nil {
		fmt.Println(_queryErr)
		panic("unable to create event table")
	}

	return db
}
