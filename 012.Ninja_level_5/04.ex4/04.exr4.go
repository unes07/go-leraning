package main

import "fmt"

// Create and use an anonymous struct.

func main() {
	s := struct {
		firstName string
		friends   map[string]int
		favDrinks []string
	}{
		firstName: "John",
		friends:   map[string]int{"Bob": 42, "Alice": 37},
		favDrinks: []string{"coke", "water", "coffee"},
	}

	for k, v := range s.friends {
		fmt.Printf("%s: %d\n", k, v)
	}

}
