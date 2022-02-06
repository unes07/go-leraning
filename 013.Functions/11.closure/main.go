package main

import "fmt"

// Scope of x is global of package main
var x int

func main() {
	fmt.Println("x:", x)
	x++
	fmt.Println("x:", x)
	foo()
	fmt.Println("x:", x)

	// Scope of y is narrow to func main
	var y int
	fmt.Println("y:", y)
	y++
	fmt.Println("y:", y)
	foo()
	fmt.Println("y:", y)

	{
		// Scope of z is narrow to the code block {}
		z := 0
		fmt.Println("z:", z)
		z++
		fmt.Println("z:", z)
	}
	// fmt.Println("z:", z) can't access z

	// Scope of b is narrow to func main
	// and a it's still clossure of incrementor
	b := incrementor()
	c := incrementor()
	fmt.Println("b1:", b())
	fmt.Println("b2:", b())
	fmt.Println("b3:", b())
	// even if it's the same identifier x but he has a different scope
	fmt.Println("c1:", c())
	fmt.Println("c2:", c())
}

func foo() {
	fmt.Println("foo")
	x++
}

func incrementor() func() int {
	// Scope of a is local to func incrementor
	var a int
	return func() int {
		a++
		return a
	}
}
