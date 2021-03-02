package main

import "fmt"

var y = 43
var z string = "a string variable"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
