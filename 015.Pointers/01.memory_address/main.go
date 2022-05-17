package main

import (
	"fmt"
)

func main() {
	a := 14

	fmt.Println(a)
	// get the adress where a is stored
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	// the type of &a is *int: pointer of an int
	fmt.Printf("%T\n", &a)

	var b *int = &a
	// b is a pointer to the adress where a is stored
	// a and b share the same adress
	fmt.Println(b)
	// get the value stored in the adress where a is stored
	fmt.Println(*b)
	// * give the value stored at an adress when you have an adress
	fmt.Println(*&a)

	*b = 15
	// change the value stored in the adress where a is stored
	fmt.Println(a)
}
