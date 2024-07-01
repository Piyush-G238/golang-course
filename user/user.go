package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	dateOfBirth string
	createdAt   time.Time
}

type Admin struct {
	email    string
	password string
	user     User
}

// method declaration in golang
func (u User) PrintUserDetails() {

	fmt.Printf("First Name: %v\n", u.firstName)
	fmt.Printf("Last Name: %v\n", u.lastName)
	fmt.Printf("Date of Birth: %v\n", u.dateOfBirth)
	fmt.Printf("Created At: %v\n", u.createdAt)
}

// mutating method in structs
func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

// constructor function for structs
func NewUser(firstName string, lastName string, dob string) (*User, error) {
	if firstName == "" || lastName == "" || dob == "" {
		return nil, errors.New("first name, last name and date of birth are mandatory to fill")
	}

	return &User{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: dob,
		createdAt:   time.Now(),
	}, nil
}

func NewAdmin(email string, password string, user User) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are mandatory to fill")
	}
	return &Admin{
		email:    email,
		password: password,
		user:     user,
	}, nil
}
