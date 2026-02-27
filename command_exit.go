package main

import (
	"fmt"
	"os"
)

// The function of the exit command
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
