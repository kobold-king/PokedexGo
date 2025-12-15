package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
		command := cleaned[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

//edit2
