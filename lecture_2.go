/*
constants in go
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64 = 1000
	years := 10.0
	rateOfInterest := 5.5

	futureValue := investmentAmount * math.Pow(1+rateOfInterest/100, years)
	var futureRealValue float64 = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
