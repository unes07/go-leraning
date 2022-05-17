package main

import (
	"fmt"
)

// *note: we use pointer to pass an adress instead of a chunk of data
// *y : dereferencr the value of y (y is a pointer)

func main() {
	x := 0
	fmt.Println("x befor", x)
	fmt.Println("x befor", &x)
	foo(&x)
	fmt.Println("x after", x)
	fmt.Println("x after", &x)
}

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 14
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
