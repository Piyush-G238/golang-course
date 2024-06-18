/*Profit calculator program using go*/
package main

import "fmt"

func main() {

	var revenue int
	var expenses int
	const taxRate float64 = 5.5

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	earningBeforeTax := revenue - expenses
	fmt.Printf("Earning before Tax: %v", earningBeforeTax)

	profit := float64(earningBeforeTax) * (1 - taxRate/100)
	fmt.Printf("\nTotal profit collected: %.2f", profit)
}
