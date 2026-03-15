package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/ReGitnald/pokedexcli/internal/utils"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("help - Show available commands")
	fmt.Println("exit - Exit the Pokedex")
	return nil
}

func commandMap(cfg *config, args ...string) error {
	// TODO: Implement the map command to show the Pokemon world map

	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	dat, ok := cfg.cache.Get(url)
	if !ok {
		dat, _ = utils.GetPokedata(url)
		cfg.cache.Add(url, dat)
	}
	loc := utils.PokeLocation{}
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

func commandMapb(cfg *config, args ...string) error {
	if cfg.Previous == nil || *cfg.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	url := *cfg.Previous
	dat, ok := cfg.cache.Get(url)
	if !ok {
		dat, _ = utils.GetPokedata(url)
		cfg.cache.Add(url, dat)
	}
	loc := utils.PokeLocation{}
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

func commandExplore(cfg *config, args ...string) error {
	placename := args[0]
	if placename == "" {
		fmt.Println("Please provide a location to explore")
		return nil
	}
	fmt.Println("Exploring ", placename, "...")
	fmt.Println("Found Pokemon:")
	//This is where the Pokemon found in the location would be printed, but since the API doesn't provide this information, we'll just print a placeholder message for now.
	url := utils.GetPokeLocationURL(placename)
	dat, ok := cfg.cache.Get(url)
	if !ok {
		dat, _ = utils.GetPokedata(url)
		cfg.cache.Add(url, dat)
	}
	loc := utils.PokeLocationArea{}
	err := json.Unmarshal(dat, &loc)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, pokemon := range loc.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	pokemonName := args[0]
	if pokemonName == "" {
		fmt.Println("Please provide a Pokemon to catch")
		return nil
	}
	fmt.Println("Throwing a Pokeball at ", pokemonName, "...")
	// Simulate catching the Pokemon with a chance based on catch rate
	url := utils.GetPokemonURL(pokemonName)
	dat, ok := cfg.cache.Get(url)
	if !ok {
		var err error
		dat, err = utils.GetPokedata(url)
		if err != nil {
			fmt.Println("Pokemon may not exist:", err)
			return err
		}
		cfg.cache.Add(url, dat)
	}
	pok := utils.Pokemon{}
	err := json.Unmarshal(dat, &pok)
	if err != nil {
		fmt.Println(err)
		return err
	}
	catchrate := pok.CaptureRate
	fmt.Println("Catch rate for ", pokemonName, ": ", catchrate)
	if rand.Intn(256) <= catchrate {
		fmt.Println("Congratulations! You caught ", pokemonName, "!")
		cfg.caughtPokemon[pokemonName] = pok
	} else {
		fmt.Println("Oh no! The ", pokemonName, " escaped!")
	}
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	pokemonName := args[0]
	if pokemonName == "" {
		fmt.Println("Please provide a Pokemon to inspect")
		return nil
	}
	fmt.Println("Inspecting ", pokemonName, "...")
	if pok, exists := cfg.caughtPokemon[pokemonName]; exists {
		fmt.Printf("ID: %d\nName: %s\nOrder: %d\nGender Rate: %d\nCapture Rate: %d\n",
			pok.ID, pok.Name, pok.Order, pok.GenderRate, pok.CaptureRate)
		return nil
	}
	return nil
}
