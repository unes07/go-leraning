package main

import (
	"fmt"
)

// Create a for loop using this synatx
// for {}
// have if print out the years you have been alive

func main() {
	i := 1998
	for {
		if i <= 2021 {
			fmt.Println(i)
			i++
		}
	}
}
