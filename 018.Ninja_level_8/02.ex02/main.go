package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"Unes","Age":24}]`
	fmt.Println(s)

	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Age)
	}
}
