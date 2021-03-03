package main

import (
	"fmt"
)

// Write a program the:
// assigns an int to a variable
// prints taht int in decimal, binary, and hex
// shift the bits of that int ober 1 position to the left, and assigns that to a variable
// prints taht variable in decimal, binary, & hex

func main() {
	a := 14
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
