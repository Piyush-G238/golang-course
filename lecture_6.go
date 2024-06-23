/* understanding functions */

package main

import (
	"fmt"
	"math"
)

// global variables
const inflatedRate = 8.0

func main() {
	amount := 12000
	rate := 5.6
	years := 10

	futureValue := calcFuture(amount, years, rate)
	// 12.345678
	realFutureValue := calcFuture(amount, years, inflatedRate)
	// 13.9486464

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.1f\n", realFutureValue)

	outputText(formattedFV)
	outputText(formattedRFV)
}

/* user defined function in golang*/
func outputText(text string) {
	fmt.Printf(text)
}

/*user defined function with return type*/
func calcFuture(amount int, years int, rate float64) (fv float64) {
	fv = float64(amount) * math.Pow(1+rate/100, float64(years))
	return fv
}
