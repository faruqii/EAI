package apps

import (
	"fmt"
	"os"

	"github.com/faruqii/EAI/internal/controllers"
)

func Start() {
	fmt.Println("Welcome to EAI Travel App")
	Menus()
}

func SearchDestination() {
	// Prompt the user to enter the search parameters
	var location, checkIn, checkOut, currency string
	var adult, children, infants int

	fmt.Print("Enter location: ")
	fmt.Scanln(&location)

	fmt.Print("Enter check-in date (YYYY-MM-DD): ")
	fmt.Scanln(&checkIn)

	fmt.Print("Enter check-out date (YYYY-MM-DD): ")
	fmt.Scanln(&checkOut)

	fmt.Print("Enter number of adults: ")
	fmt.Scanln(&adult)

	// Prompt for children only if user specifies that they are travelling with children
	fmt.Print("Travelling with children? (yes/no): ")
	var withChildren string
	fmt.Scanln(&withChildren)
	if withChildren == "yes" {
		fmt.Print("Enter number of children: ")
		fmt.Scanln(&children)
	} else {
		children = 0
	}

	// Prompt for infants only if user specifies that they are travelling with infants
	fmt.Print("Travelling with infants? (yes/no): ")
	var withInfants string
	fmt.Scanln(&withInfants)
	if withInfants == "yes" {
		fmt.Print("Enter number of infants: ")
		fmt.Scanln(&infants)
	} else {
		infants = 0
	}

	fmt.Print("Enter currency (USD): ")
	fmt.Scanln(&currency)

	// Call the SearchDestination function with the user input
	controllers.SearchDestination(location, checkIn, checkOut, adult, children, infants, currency)
}

func Menus() {
	var menu int
	fmt.Println("1. Search Destination")
	fmt.Println("0. Exit")
	fmt.Print("Choose menu: ")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		SearchDestination()
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid menu")
	}
}
