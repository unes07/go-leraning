package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	// not recomanded
	fmt.Println(m["Boney"])

	// better way
	v, ok := m["Boney"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Boney"]; ok {
		fmt.Println("This is from if", v)
	}

	if v, ok := m["James"]; ok {
		fmt.Println("This is from if", v)
	}
}
