package main

import (
	"fmt"
	"math"
)

// *note:
// *Recivier   Value
// *(t T)      T AND *T
// *(t *T)     *T

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s *square) area() float64 {
	return s.side * s.side
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{
		radius: 10,
	}

	s := square{
		side: 10,
	}

	// NON-POINTER RECIVER & NON-POINTER VALUE
	info(c)

	// NON-POINTER RECIVER & POINTER VALUE
	info(&c)

	// POINTER RECIVER & POINTER VALUE
	info(&s)

	// POINTER RECIVER & NON-POINTER VALUE: not working by definition
	// info(s)

}
