package main

import "fmt"

func main() {
	// simple sum
	ii := []int{1, 2, 3, 4, 5, 6}
	s := sum(ii...)
	fmt.Println("sum of all numbers:", s)

	// sum of even numbers using a callback
	s2 := even(sum, ii...)
	fmt.Println("sum of even numbers:", s2)

	// sum of odd numbers using a callback
	s3 := odd(sum, ii...)
	fmt.Println("sum of odd numbers:", s3)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
