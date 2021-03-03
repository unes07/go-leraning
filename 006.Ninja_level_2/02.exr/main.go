package main

import (
	"fmt"
)

// Using the following operators, write expressions
//  & assign their values to variables

func main() {
	a := (14 == 14)
	b := (14 <= 40)
	c := (14 >= 40)
	d := (14 != 11)
	e := (14 < 11)
	f := (14 > 11)

	fmt.Println(a, b, c, d, e, f)
}
