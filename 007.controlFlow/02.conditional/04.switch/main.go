package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this from false")
	case (2 == 4):
		fmt.Println("this from 2 = 4")
	case (3 == 3):
		fmt.Println("this from 3 = 3")
		fallthrough
	case (7 == 4):
		fmt.Println("this from 4 = 4")
		fallthrough
	case (4 == 4):
		fmt.Println("this from 4 = 4")
	default:
		fmt.Println("defaul in case no one true")
	}
	// faltthrough try to not use it

	fmt.Println("==============================")

	// SWITCH IN A VARIABLE
	n := "Bond"
	switch n {
	case "Moneypenny", "XY", "Z":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("james bond")
	case "Q":
		fmt.Println("this is a q")
	default:
		fmt.Println("this is default")
	}

}
