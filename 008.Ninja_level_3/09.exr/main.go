package main

import (
	"fmt"
)

// create a program that uses "switch" statement
// with the switch expresiion specified
// as a variable of type string with the identifier "favSport"

func main() {
	favSport := "football"
	switch favSport {
	case "rugby":
		fmt.Println("rugby")
	case "football":
		fmt.Println("football")
	case "surfing":
		fmt.Println("surfing")
	}
}
