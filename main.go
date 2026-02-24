package main

import (
	"bufio"
	"fmt"
	"os"
)

type PokeLocation struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	// Add any configuration fields you need here
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func main() {
	cfg := &config{}
	commands :=
		map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Show available commands",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Show the map of the Pokemon world",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Show the map of the Pokemon world (backwards)",
				callback:    commandMapb,
			},
		}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		commandName, exists := commands[words[0]]
		if exists {
			err := commandName.callback(cfg)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", words[0])
			fmt.Println("Type 'help' to see available commands.")
		}
	}
}
