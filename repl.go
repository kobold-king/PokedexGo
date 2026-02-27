package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kobold-king/PokedexGo/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	prevURL       *string
}

func startRepl(cfg *config) {
	// 1. Create a scanner to wait for user input from Standard Input
	scanner := bufio.NewScanner(os.Stdin)
	// 2. Start an infinite for loop for the REPL
	for {
		// 3. Print the prompt without a newline character
		fmt.Print("Pokedex > ")
		// 4. Use .Scan() to block and wait for the user to press enter
		success := scanner.Scan()
		if !success {
			break
		}
		// 5. Get the text and clean the user's input
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		// 6. Capture the first word and print it
		commandinput := cleaned[0]

		// Check to see if the inputted command exists
		command, exists := getCommands()[commandinput]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// Command structure
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		//custom clicomman for practice
		"banana": {
			name:        "banana",
			description: "Posts a banana",
			callback:    commandBanana,
		},
	}
}
