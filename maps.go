package main

import "fmt"

type FloatMap map[string]float64

func (m FloatMap) output(key string) {
	fmt.Println(m[key])
}

func main() {

	userNames := make(map[string]string, 4)
	// make([]string, 2)
	userNames["Amazon"] = "https://www.amazon.in"
	userNames["Microsoft"] = "https://www.microsoft.com"

	fmt.Println(userNames["Amazon"])
	fmt.Println(userNames["Microsoft"])

	workingWithLoops()
}

// working with maps
func workWithMaps() {

	websites := map[string]string{
		"google": "https://google.com",
		"aws":    "https://aws.com",
	}

	fmt.Println(websites["google"])

	// mutating map
	websites["google"] = "https://www.google.co.in"
	fmt.Println(websites["google"])
}

// working with for loops
func workingWithLoops() {

	userNames := []string{"Manuel", "Adam", "Micheal", "Piyush", "Divyanshu"}

	for _, value := range userNames {
		fmt.Println(value)
	}

	courseRatings := FloatMap{
		"Angular": 3.5,
		"React":   4.5,
		"Vue.js":  3.0,
	}

	for key, _ := range courseRatings {
		courseRatings.output(key)
	}
}
