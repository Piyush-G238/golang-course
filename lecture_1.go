/*
every go program starts with package declaration. package main statement indicates that file is part of main package.
main package is special because it is entry point for executable programs.
*/
package main

/*
import keyword is used to include code from other packages to use in our program.
fmt(format) package provides functions for formatting input & output.
*/
import (
	"fmt"
	"math"
)

/*
main function is special function in Go. It serves as entry point for executable program.
*/
func main() {

	/*
		This line outputs text "Hello, world" to the console.
	*/
	fmt.Println("Hello, world")

	var investedAmount float64 = 2000
	var returnRate float64 = 5
	var years = float64(10)

	var returnedAmount = investedAmount * math.Pow(1+returnRate/100, years)

	// fmt.Println(years)
	fmt.Println(returnedAmount)
}
