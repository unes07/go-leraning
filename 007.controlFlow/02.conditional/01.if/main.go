package main

import "fmt"

func main() {
	if true {
		fmt.Println("Hello unes")
	}

	if false {
		fmt.Println("don't hello me")
	}

	if x := 14; x == 2 {
		fmt.Println(x)
	}
}
