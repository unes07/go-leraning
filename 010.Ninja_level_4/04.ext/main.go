package main

import (
	"fmt"
)

// Append 52, 53, 54, 55 to this slice {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

func main() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(arr)

	arr = append(arr, 52)
	fmt.Println(arr)

	arr = append(arr, 53, 54, 55)
	fmt.Println(arr)

	x := []int{56, 57, 58, 59, 60}
	arr = append(arr, x...)
	fmt.Println(arr)
}
