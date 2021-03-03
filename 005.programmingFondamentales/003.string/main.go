package main

import (
	"fmt"
)

func main() {
	s := "Hello playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
