package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ReGitnald/pokedexcli/internal/pokecache"
	"github.com/ReGitnald/pokedexcli/internal/utils"
)

type config struct {
	// Add any configuration fields you need here
	Next          *string
	Previous      *string
	cache         *pokecache.Cache
	caughtPokemon map[string]utils.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func main() {
	cfg := &config{
		cache: pokecache.NewCache(5 * time.Second),
	}
	cfg.caughtPokemon = make(map[string]utils.Pokemon)
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
			"explore": {
				name:        "explore",
				description: "Explore a specific location and list available Pokemon",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Attempt to catch a specific Pokemon",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect a caught Pokemon for details",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "List all caught Pokemon",
				callback:    commandPokedex,
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
			args := words[1:]
			err := commandName.callback(cfg, args...)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", words[0])
			fmt.Println("Type 'help' to see available commands.")
		}
	}
}
