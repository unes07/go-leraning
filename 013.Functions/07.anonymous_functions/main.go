package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("the value of x is:", x)
	}(42)
}

func foo() {
	fmt.Println("foo")
}
