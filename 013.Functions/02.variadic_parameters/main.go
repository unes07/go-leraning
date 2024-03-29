package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5, 6)
	fmt.Println("The total is,", x)
}

func sum(x ...int) int {
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
