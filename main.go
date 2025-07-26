package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokecache"
	"github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokedex"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var mainConfig config

	cache := pokecache.NewCache(1 * time.Minute)

	pokedex := pokedex.NewPokedex()

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		userInput := cleanInput(scanner.Text())
		enteredCommand := userInput[0]
		var parameter string
		if len(userInput) < 2 {
			parameter = ""
		} else {
			parameter = userInput[1]
		}
		cliCommand, exists := getCommands()[enteredCommand]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			errorReturned := cliCommand.callback(&mainConfig, cache, pokedex, parameter)
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
		"explore": {
			name:        "explore",
			description: "Takes a location area as an argument. Displays a list of all Pokemon located there.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes the name of a pokemon as an argument.\nAttempt to catch the named pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Takes the name of a pokemon as an argument.\nDisplay the stats of a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the names of all the Pokemon you have caught so far",
			callback:    commandPokedex,
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
