package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello")

}

func foo() {
	fmt.Println("world")
}
