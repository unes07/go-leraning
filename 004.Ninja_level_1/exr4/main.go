package main

import (
	"fmt"
)

type unes int

var x unes

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
