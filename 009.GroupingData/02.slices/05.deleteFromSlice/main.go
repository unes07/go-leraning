package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7, 9, 56}
	fmt.Println(x)
	x = append(x, 17, 14, 03)
	fmt.Println(x)

	y := []int{100, 101, 404}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
