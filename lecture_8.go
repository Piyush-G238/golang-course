/*
	understanding control structure, if, else if, loops etc.
	also understanding read and write of data in a file
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const currentFile = "balance.txt"

func main() {
	accountBalance, _error := readBalanceFromFile()

	if _error != nil {
		// The panic built-in function stops normal execution of the current goroutine.
		panic(_error.Error())
	}

	//  = 1000
	var depositAmount float64
	var withdrawAmount float64
	var choice int

	for choice != 4 {
		fmt.Print(`Welcome to Go Bank!
What do you want to do?
1. Check the balance
2. Deposit money
3. Withdraw money
4. Exit	`)

		fmt.Print("\nYour choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			fmt.Println("How much do you want to deposit? ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Print("Deposit amount should be greater than 0\n")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Balance updated, New Amount: %v\n", accountBalance)
		} else if choice == 3 {
			fmt.Println("How much do you want to withdraw? ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Print("Withdrawal amount cannot be less than 0\n")
				continue
			} else if withdrawAmount > accountBalance {
				fmt.Print("Withdrawal amount cannot be greater than account balance\n")
				continue
			}

			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Balance updated, New Amount: %v\n", accountBalance)
		} else {
			fmt.Print("Thank you for using our application.")
			break
		}
	}

}

func writeBalanceToFile(balance float64) {
	currentBalance := fmt.Sprint(balance)
	os.WriteFile(currentFile, []byte(currentBalance), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, _error := os.ReadFile(currentFile)

	if _error != nil {
		return 0, errors.New("Failed to read account balance")
	}

	balanceText := string(data)

	balance, _parseError := strconv.ParseFloat(balanceText, 64)

	if _parseError != nil {
		return 0, errors.New("Unable to parse account balance from file")
	}
	return balance, nil
}
