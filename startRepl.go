package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}

		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputText := cleanInput(scanner.Text())
		commandName := inputText[0]

		value, ok := cfg.commands[commandName]
		if ok{
			err := value.callback(cfg)
			if err != nil{
				fmt.Println(err)
			}	
		}else{
			fmt.Println("Unknown command")
		}

	}
}
