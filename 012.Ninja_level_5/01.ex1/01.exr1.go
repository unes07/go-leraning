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

	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)

}
