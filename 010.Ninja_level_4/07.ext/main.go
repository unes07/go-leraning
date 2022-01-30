package main

import (
	"fmt"
)

// a slice of a slice string

func main() {
	xs := []string{"james", "bond", "shaken, not stirred"}
	fmt.Println(xs)

	xy := []string{"miss", "moneypenny", "stole", "all", "your", "gold"}
	fmt.Println(xy)

	arr := [][]string{xs, xy}
	fmt.Println(arr)

	for i, v := range arr {
		fmt.Println("array:", i)
		for j, v2 := range v {
			fmt.Println("\t", j, v2)
		}
	}

}
