package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(s))
	fmt.Println(string(bs))

	pass := "password123"
	err = bcrypt.CompareHashAndPassword(bs, []byte(pass))
	if err != nil {
		fmt.Println("Password does not match")
		log.Fatal(err)
		return
	}
	fmt.Println("Password is correct")
}
