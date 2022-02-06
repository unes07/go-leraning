package main

import "fmt"

func main() {
	x := factorial(5)
	fmt.Println(x)

	y := loop_factorial(4)
	fmt.Println(y)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loop_factorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
