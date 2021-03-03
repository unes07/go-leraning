package main

import (
	"fmt"
)

// Automatic Types

// func main() {
// 	x := 42
// 	y := 42.3265
// 	fmt.Printf("%T\n", x)
// 	fmt.Printf("%T\n", y)
// }

// mannuel Types Decleration

var x int
var y float64

func main() {
	x = 42
	y = 42.3265
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}

// uint8 (0 top 255): Byte: unigative int8 (-128 to 127)
// rune : int32
