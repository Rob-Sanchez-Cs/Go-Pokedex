package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		userInput := cleanInput(scanner.Text())
		enteredCommand := userInput[0]
		cliCommand, exists := getCommands()[enteredCommand]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			errorReturned := cliCommand.callback()
			if errorReturned != nil {
				fmt.Println(errorReturned)
			}
		}

		fmt.Print("Pokedex > ")
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cliCommand := range getCommands() {
		fmt.Printf("%v: %v\n", cliCommand.name, cliCommand.description)
	}

	return nil
}

func cleanInput(input string) []string {
	trimmed := strings.TrimSpace(input)

	slice := strings.Fields(trimmed)
	for i := range slice {
		slice[i] = strings.ToLower(slice[i])
	}
	return slice
}
