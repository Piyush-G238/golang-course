/* formatting texts in go programming*/

package main

import "fmt"

func main() {
	futureValue := 12.345678
	realFutureValue := 13.9486464

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.1f\n", realFutureValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)

	/*formatting multiline string*/
	var lang1 string = "java"
	var lang2 string = "golang"

	fmt.Printf(`Currently I am a %v developer,
but I want to learn %v also`, lang1, lang2)

}
