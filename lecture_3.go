/*
Scan func in go
Taking input from user through console
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var years float64
	const rateOfInterest = 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("\nDuration of Investment: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+rateOfInterest/100, years)
	var futureRealValue float64 = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
