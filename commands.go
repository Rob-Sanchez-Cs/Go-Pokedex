package main

import (
	"fmt"
	"os"
)

func commandExit(mainConfig *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(mainConfig *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cliCommand := range getCommands() {
		fmt.Printf("%v: %v\n", cliCommand.name, cliCommand.description)
	}

	return nil
}

func commandMap(mainConfig *config, usePreviousUrl bool) error {
	var mapsResponse getMapResponse
	error := getMaps(&mapsResponse, mainConfig, usePreviousUrl)
	if error != nil {
		return error
	}

	for _, area := range mapsResponse.Results {
		fmt.Println(area.Name)
	}

	mainConfig.Next = mapsResponse.Next
	mainConfig.Previous = mapsResponse.Previous

	return nil

}

func commandMapNormal(mainConfig *config) error {
	return commandMap(mainConfig, false)
}

func commandMapBack(mainConfig *config) error {
	return commandMap(mainConfig, true)
}