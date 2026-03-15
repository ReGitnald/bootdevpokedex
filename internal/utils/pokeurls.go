package utils

import (
	"fmt"
	"io"
	"net/http"
)

func GetPokeLocationURL(name string) string {
	return "https://pokeapi.co/api/v2/location-area/" + name + "/"
}

func GetPokedata(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	return data, nil
}
func GetPokemonURL(name string) string {
	return "https://pokeapi.co/api/v2/pokemon/" + name + "/"
}
