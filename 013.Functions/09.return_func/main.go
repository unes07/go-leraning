package main

import "fmt"

func main() {

	// return the insider func of bar()
	x := bar()
	fmt.Printf("%T\n", x)

	// execute the insider func of bar() after storing it in x variable
	i := x()
	fmt.Println(i)
}

func bar() func() int {
	return func() int {
		return 14
	}
}
