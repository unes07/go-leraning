package main

import "fmt"

func main() {
	// pass a slice of int as an argument while sum waiting for ints
	// unfurling a slice of int
	xi := []int{2, 3, 4, 5, 6}
	x := sum(xi...)
	fmt.Println("The total is,", x)

	// passing a zero or more values as arguments (varidic function)
	y := sum()
	fmt.Println("The total is,", y)
}

func sum(x ...int) int {
	// NB: the varidic parameter must be the last one
	// type of x here is a slice of int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are adding", v, "to the total which is now", sum)
	}

	return sum
}
