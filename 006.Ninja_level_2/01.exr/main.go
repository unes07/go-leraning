package main

import (
	"fmt"
)

// Write a program that prints a number in:
// decimal, binary and hex

func main() {
	x := 14
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}
