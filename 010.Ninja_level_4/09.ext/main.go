package main

import (
	"fmt"
)

// Create a map of type string which is a person's name and favorite song.
// Delete a record from the map.
// Print out the map.

func main() {

	m := map[string][]string{
		"Yassine": {"Locked out of Heaven", "when you look me in the eyes", "when I was your man"},
		"Younes":  {"Millions years ago", "The night we meet", "Happier"},
	}

	m["Saad"] = []string{"I'm the king of the world", "Halo", "You are the reason"}

	delete(m, "Younes")

	for k, v := range m {
		fmt.Println(k, ":")
		for j := range v {
			fmt.Println("\t", j, ":", v[j])
		}
	}

}
