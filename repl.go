package main

import(
	"strings"
)

func cleanInput(text string)[]string{
	cleanedString := []string{}
	lowered := strings.ToLower(text)
	splitString := strings.Fields(lowered)	
	
	cleanedString = append(cleanedString, splitString...)

	return cleanedString
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{}

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}