package anonymous

import (
	"fmt"
)

func WorkingWithAnonymous() {

	newFunction := func(num int) int {
		return num * 2
	}

	fmt.Println(newFunction(2))
}

// example of closure
func WorkingWithClosures(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}

// variadic function
func SumUp(nums ...int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}

	return sum
}
