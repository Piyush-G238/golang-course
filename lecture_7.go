/* profit calculator program to practice functions*/

package main

import (
	"fmt"
)

var revenue float64
var expenses float64
var taxRate float64

func main() {

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("\nExpense: ")
	fmt.Scan(&expenses)

	fmt.Print("\nTax rate: ")
	fmt.Scan(&taxRate)

	ebt := calcEbt()
	profit := calcProfit(ebt)

	ratio := calcRatio(ebt, profit)

	fmt.Printf("Estimated earning before tax: %.2f", ebt)
	fmt.Printf("\nEstimated profit: %.2f", profit)
	fmt.Print("\nEstimated ratio: %.2f", ratio)
}

func calcEbt() float64 {
	return revenue - expenses
}

func calcProfit(ebt float64) float64 {
	return ebt * (1 - taxRate/100)
}

func calcRatio(ebt float64, profit float64) float64 {
	return ebt / profit
}
