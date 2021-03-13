package main

import (
	"fmt"
)

// Using a composit Literal:
// 	- Create an slice Which  of Types int
// 	- Assign 10 Values

// Range over the array and print the values out

// Using format printing:
// 	- Print out Type of the slice

func main() {
	arr := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(arr)

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", arr)

}
