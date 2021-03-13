package main

import (
	"fmt"
)

// Using a composit Literal:
// 	- Create an Array Which holds 5 Values of Types int
// 	- Assign Values to each index position

// Range over the array and print the values out

// Using format printing:
// 	- Print out Type of the array

func main() {
	arr := [5]int{10, 11, 12, 13, 14}
	fmt.Println(arr)

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", arr)

}
