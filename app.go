package main

import (
	"fmt"

	"example.com/user"
)

/*
type User struct {
	firstName   string
	lastName    string
	dateOfBirth string
	createdAt   time.Time
}
*/

func main() {

	var firstname string
	var lastname string
	var dateOfBirth string

	fmt.Println("Please enter firstName: ")
	fmt.Scanln(&firstname)

	fmt.Println("Please enter LastName: ")
	fmt.Scanln(&lastname)

	fmt.Println("Please enter Date of Birth: ")
	fmt.Scanln(&dateOfBirth)

	user1, _error := user.NewUser(firstname, lastname, dateOfBirth)

	if _error != nil {
		fmt.Printf("%v\n", _error.Error())
		panic("Invalid creation of user!")
	}
	// User{}

	// user1.createdAt = time.Now()

	/*
		user1 := User{
			firstName:   "Piyush",
			lastName:    "Gupta",
			dateOfBirth: "28/06/1998",
			createdAt:   time.Now(),
		}
	*/

	// userPointer := &user1

	user1.PrintUserDetails()
	user1.ClearUsername()
	user1.PrintUserDetails()

	// fmt.Printf("Modified First Name: %v\n", user1.firstName)

}

/*
func printUserDetails(user *User) {
	fmt.Printf("First Name: %v\n", user.firstName)
	fmt.Printf("Last Name: %v\n", user.lastName)
	fmt.Printf("Date of Birth: %v\n", user.dateOfBirth)
	fmt.Printf("Created At: %v\n", user.createdAt)

	// user.firstName = "Arjun"
}
*/
