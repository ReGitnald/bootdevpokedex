package main
import (
	"fmt"
	"slices"
	"testing"
)


func TestCleanInput(t *testing.T) {
	input := "Hello, World"
	expected := []string{"Hello", "World"}
	result := cleanInput(input)
	if !slices.Equal(result, expected) {
		t.Errorf("Test failed: expected %v, got %v\n", expected, result)
	} else {
		fmt.Println("Test passed!")
	}
}