/* understanding switch statement*/

package main

import "fmt"

func main() {

	fmt.Println("Welcome to MyApp!")

	fmt.Print(`
1. View profile
2. Edit settings
3. Exit`)

	fmt.Print("\nYour choice: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Viewing profile...")
		break

	case 2:
		fmt.Println("Editing settings...")
		break

	case 3:
		fmt.Println("Exiting...")
		break

	default:
		fmt.Println("Invalid choice!")
	}
}
