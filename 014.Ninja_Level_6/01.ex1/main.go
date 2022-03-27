package main

func main() {
	n := foo()
	x, s := bar()

	println(n, x, s)
}

func foo() int {
	return 14
}

func bar() (int, string) {
	return 11, "just another number"
}
