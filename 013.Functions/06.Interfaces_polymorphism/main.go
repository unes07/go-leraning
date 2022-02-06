package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak!")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak!")
}

// interfaces allow to a value to be more than one type
type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr() as a person", h.(person).first, h.(person).last)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr() as a secretAgent", h.(secretAgent).first, h.(secretAgent).last)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	// a value can be of more than one type
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int = int(x)
	// y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
