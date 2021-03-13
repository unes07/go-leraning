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

	// add new item to m
	m["todd"] = 33
	fmt.Println(m)

	// Range
	for i, v := range m {
		fmt.Println(i, v)
	}

	// range over a slice
	xi := []int{4, 7, 9, 14}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
