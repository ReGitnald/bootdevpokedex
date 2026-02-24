package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ReGitnald/pokedexcli/internal/utils"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	fmt.Println("help - Show available commands")
	fmt.Println("exit - Exit the Pokedex")
	return nil
}

func commandMap(cfg *config) error {
	// TODO: Implement the map command to show the Pokemon world map

	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	dat, _ := utils.GetPokedata(url)
	loc := PokeLocation{}
	err := json.Unmarshal(dat, &loc)
	if err != nil {
		fmt.Println(err)
	}
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = &loc.Next
	cfg.Previous = &loc.Previous
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	url := *cfg.Previous
	dat, _ := utils.GetPokedata(url)
	loc := PokeLocation{}
	err := json.Unmarshal(dat, &loc)
	if err != nil {
		fmt.Println(err)
	}
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = &loc.Next
	cfg.Previous = &loc.Previous
	return nil
}
