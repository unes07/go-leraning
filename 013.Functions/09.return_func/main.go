package main

import "fmt"

func main() {

	// return the insider func of bar()
	fmt.Printf("%T\n", bar())
	// execute the insider func of bar() after storing it in x variable
	fmt.Println(bar()())
}

func bar() func() int {
	return func() int {
		return 14
	}
}
