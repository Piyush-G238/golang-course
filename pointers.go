/* understanding concept of pointers in go*/
package main

import "fmt"

func main() {
	age := 32
	fmt.Print("Age: ", age)

	// adultYears := getAdultYears(&age)
	getAdultYears(&age)

	// var agePointer *int = &age
	fmt.Print("\nAdult Years: ", age)
}

/*
func getAdultYears(age int) int {
	return age - 18
}
*/

func getAdultYears(age *int) {
	*age -= 18
	// return *age
}
