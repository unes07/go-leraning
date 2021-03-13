package main

import (
	"fmt"
)

func main() {
	jb := []string{"Colt", "steel", "coffee", "Vanilla"}
	fmt.Println(jb)
	mp := []string{"Maxi", "Millan", "Thea", "Mint"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
