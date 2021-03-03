package main

import (
	"fmt"
)

const (
	a int     = 42
	b float64 = 42.4654
	c string  = "Ronaldo, Colt"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}