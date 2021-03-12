package main

import (
	"fmt"
)

// create a program that uses "switch" statement with no switch expresiion specified

func main() {
	switch {
	case (4 == 4):
		fmt.Println("thats true")
	case (5 == 4):
		fmt.Println("thats not")
	}
}
