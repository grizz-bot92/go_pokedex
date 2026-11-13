package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}

		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputText := scanner.Text()
		cleanInput(inputText)

		value, ok := commands[inputText]
		if ok {
			switch value.name {
			case "help":
				commandHelp()
			case "exit":
				commandExit()
			default:
				fmt.Println("Unknown command")
			}
		}

	}
}
