package utils

import (
	"fmt"
	"io"
	"net/http"
)

// func GetPokeLocationURL(id int) string {
// 	return "https://pokeapi.co/api/v2/location-area/" + fmt.Sprint(id) + "/"
// }

func GetPokedata(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	return data, nil
}
