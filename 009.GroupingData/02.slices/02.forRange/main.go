package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7, 9, 56}

	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[5])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(cap(x))
}

// a SLICE allows you to group together Values of the same type
