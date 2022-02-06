package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("func expression")
	}
	f()

	b := func(x int) {
		fmt.Println("the weirdest number is:", x)
	}
	b(14)
}
