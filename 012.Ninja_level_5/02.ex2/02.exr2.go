package main

import (
	"fmt"
)

// Create your own type Person which will have an underlying type of struct
// so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two Values of type Person
// Print out the data from the type using Printf

type Person struct {
	firstName string
	lastName  string
	flavors   []string
}

func main() {
	p1 := Person{
		firstName: "Younes",
		lastName:  "boudoul",
		flavors:   []string{"vanilla", "citron"},
	}

	p2 := Person{
		firstName: "Yassine",
		lastName:  "taki",
		flavors:   []string{"chocolate", "strawberry"},
	}

	m := map[string]Person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Printf("%v\n", m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
}
