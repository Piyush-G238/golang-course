package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	db, _err := sql.Open("mysql", "root:hH&qeV%y12@tcp(localhost)/golang_dev?parseTime=true")

	if _err != nil {
		fmt.Println(_err)
		panic("unable to connect with given database")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		firstName VARCHAR(50) NOT NULL,
		lastName VARCHAR(50),
		username VARCHAR(100) NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	`
	_, _queryErr := db.Exec(createUserTable)

	if _queryErr != nil {
		fmt.Println(_queryErr)
		panic("unable to create users table")
	}

	createEventTable := `
CREATE TABLE IF NOT EXISTS events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    dateTime TIMESTAMP NOT NULL,
    userId INT NOT NULL,
	FOREIGN KEY (userId) REFERENCES users(id)
);
`
	_, _queryErr = db.Exec(createEventTable)

	if _queryErr != nil {
		fmt.Println(_queryErr)
		panic("unable to create event table")
	}

	return db
}
