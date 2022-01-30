package main

import (
	"fmt"
)

// Use this slice {42, 43, 44, 45, 46, 47, 48, 49, 50, 51} to print this slice [42 43 44 48 49 50 51]

func main() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(arr)

	arr = append(arr[:3], arr[6:]...)
	fmt.Println(arr)

}
