package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var mainConfig config

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		userInput := cleanInput(scanner.Text())
		enteredCommand := userInput[0]
		cliCommand, exists := getCommands()[enteredCommand]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			errorReturned := cliCommand.callback(&mainConfig)
			if errorReturned != nil {
				fmt.Println(errorReturned)
			}
		}

		fmt.Print("Pokedex > ")
	}
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
			description: "Displays the names of 20 location areas in the Pokemon world.\nSubsequent calls will display the next 20 locations.",
			callback:    commandMapNormal,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.\nSubsequent calls will display the previous 20 locations.",
			callback:    commandMapBack,
		},
	}
}

func cleanInput(input string) []string {
	trimmed := strings.TrimSpace(input)

	slice := strings.Fields(trimmed)
	for i := range slice {
		slice[i] = strings.ToLower(slice[i])
	}
	return slice
}
