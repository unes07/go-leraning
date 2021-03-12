package main

import (
	"fmt"
)

// Print out the remainder for each num between 10 & 100 when it is devided by 4

func main() {
	for i := 10; i < 101; i++ {
		fmt.Printf("When %v is devided by 4, the remainder is %v\n", i, i%4)
	}
}
