package main

import(
	"fmt"
)

func commandHelp(c *config) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n\n")

	for _, cmd := range c.commands{
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}