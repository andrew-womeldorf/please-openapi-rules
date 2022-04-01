package main

import (
	"fmt"

	"examples/generator"
)

func main() {
	pet := generator.Pet{
		Name: "Fluffy",
		//PhotoUrls: []string{
		//"foo",
		//"bar",
		//},
	}
	if err := generator.AssertPetRequired(pet); err != nil {
		fmt.Printf("something wrong with your pet: %v\n", err)
	} else {
		fmt.Println("Nothing wrong with your pet")
	}
}
