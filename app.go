package main

import (
	"fmt"

	"example.com/anonymous"
)

type transFormFn func(int, int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	// passing function as value in another function
	dNumbers := transformNumbers(numbers, getMultiplyFn(), 2)

	fmt.Println(dNumbers)

	fmt.Println(factorial(5))

	anonymous.WorkingWithAnonymous()

	trnasFormerFn := anonymous.WorkingWithClosures(5)
	fmt.Println(trnasFormerFn(5))
}

func transformNumbers(numbers []int, fn transFormFn, mul int) []int {

	dNumbers := []int{}
	for _, value := range numbers {
		dNumbers = append(dNumbers, fn(value, mul))
	}
	return dNumbers
}

func multiply(val int, mul int) int {
	return val * mul
}

// returning function as value
func getMultiplyFn() transFormFn {
	return multiply
}

// recursive function
func factorial(num int) int {
	// result := 1
	if num == 0 || num == 1 {
		return 1
	}
	/*

		for i := 0; i <= num; i++ {
			result *= i
		}
	*/
	return num * factorial(num-1)
}
