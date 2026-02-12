package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		fmt.Print(
			"Your command was: ", words[0], "\n",
		)
	}
}
