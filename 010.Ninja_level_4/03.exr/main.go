package main

import (
	"fmt"
)

// Use Slicing to create the following new slices:

// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]

func main() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(arr)

	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[2:7])
	fmt.Println(arr[1:6])

}
