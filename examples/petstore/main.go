package main

import (
	"fmt"

	"examples/petstore"
)

func main() {
	pet := petstore.Pet{
		Name: "Fluffy",
		//PhotoUrls: []string{
		//"foo",
		//"bar",
		//},
	}
	if err := petstore.AssertPetRequired(pet); err != nil {
		fmt.Printf("something wrong with your pet: %v\n", err)
	} else {
		fmt.Println("Nothing wrong with your pet")
	}
}
