package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changedMe(p *person) {
	p.first = "Mr."
	// OR WE CAN USE: (*p).first = "Mr."
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
	changedMe(&p1)
	fmt.Println(p1)

}
