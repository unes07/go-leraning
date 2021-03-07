package main

import "fmt"

func main() {
	x := 14
	if x == 11 {
		fmt.Println("if")
	} else if x == 14 {
		fmt.Println("from else if")
	} else {
		fmt.Println("from else")
	}
}
