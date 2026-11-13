package main

import(
	"fmt"
)

func commandHelp() error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n\n")

	for _, v := range commands{
		fmt.Println(v.name)
	}

	return nil
}