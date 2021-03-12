package main

import (
	"fmt"
)

// Building on the previous hands-on exercice,
// create a program that uses "else if" and "else"

func main() {
	x := "Steel"

	if x == "colt" {
		fmt.Println(x)
	} else if x == "Steel" {
		fmt.Println(x)
	} else {
		fmt.Println("Nothing")
	}
}
