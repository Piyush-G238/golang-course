package lists

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

// excercise based on list and slices in go
func excercise1() {

	hobbies := [3]string{"Playing GTA V", "Eating food", "Coding"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1], hobbies[2])

	mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)

	mainHobbies = hobbies[1:]
	fmt.Println(mainHobbies)

	goals := []string{"Completing Golang this week", "Starting microservices next week"}
	goals = append(goals, "Low level design from august")
	fmt.Println(goals)

	myProducts := []Product{
		{
			title: "A book",
			price: 12.99,
			id:    1,
		},
		{
			title: "A laptop",
			price: 14.99,
			id:    2,
		},
	}
	myProducts = append(
		myProducts,
		Product{
			title: "A Chair",
			price: 18.99,
			id:    3,
		})
	fmt.Println(myProducts)
}

func workingWithList() {

	var productNames [4]string = [4]string{"A Book", "A pencil"}
	fmt.Println(productNames[0])

	productNames[2] = "A Laptop"
	fmt.Println(productNames[2])

	prices := [10]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	// creating slice from array
	featuredPrice := prices[1:]
	// := prices[1:3]
	fmt.Println(featuredPrice)

	// creating slice from another slice
	highlightedPrices := featuredPrice[:1]
	fmt.Println(highlightedPrices)

	featuredPrice[0] = 33.11
	fmt.Println(len(featuredPrice), cap(featuredPrice))

	// creating dynamic list
	newPrices := []float64{10.99, 8.99}
	fmt.Println(newPrices[1])

	newPrices = append(newPrices, 5.99)
	// newPrices[2] = 11.99
	fmt.Println(newPrices[2])
}

// unpacking list values
func unpackingList() {

	prices := []float64{10.99, 8.99}
	prices = append(prices, 5.99, 12.99, 29.99, 100.10)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
