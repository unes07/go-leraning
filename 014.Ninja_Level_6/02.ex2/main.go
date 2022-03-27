package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := foo(xi...)
	fmt.Println(n)

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := bar(ii)
	fmt.Println(m)

}

func foo(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
