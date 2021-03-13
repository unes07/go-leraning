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

	delete(m, "James")
	fmt.Println(m)

	// delete something that not exist (no error)
	delete(m, "Colt")
	fmt.Println(m)

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Colt"])

	// comma ok idom
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)

}
