package main

import (
	"strings"
)

func cleanInput(text string) []string {
	sep := " "
	i := strings.Split(text, sep)
	var result []string
	for _, word := range i {
		result = append(result, strings.ToLower(strings.Trim(word, ".,!?;:\"()[]{}")))
	}
	return result
}
