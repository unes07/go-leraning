package main

import "fmt"

// func (r reciver) identifier(parameters) (return(s)) {...}

func main() {
	foo()

	bar("Younes!")

	s1 := woo("Yassine")
	fmt.Println(s1)

	x, y := mouse("Younes", "BOUDOUL")
	fmt.Println(x)
	fmt.Println(y)
}

// Basic func
func foo() {
	fmt.Println("Hello from foo!")
}

// everything in GO is pass by value

// takes an argument of type string
func bar(s string) {
	fmt.Println("Hello, ", s)
}

// returns a string
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

// multiple return values
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}
