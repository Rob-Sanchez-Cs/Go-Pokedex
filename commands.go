package main

import (
	"fmt"
	"os"

	"github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokecache"
)

func commandExit(mainConfig *config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(mainConfig *config, cache *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cliCommand := range getCommands() {
		fmt.Printf("%v: %v\n\n", cliCommand.name, cliCommand.description)
	}

	return nil
}

func commandMap(mainConfig *config, usePreviousUrl bool, cache *pokecache.Cache) error {
	var mapsResponse getMapResponse
	error := getMaps(&mapsResponse, mainConfig, usePreviousUrl, cache)
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

func commandMapNormal(mainConfig *config, cache *pokecache.Cache) error {
	return commandMap(mainConfig, false, cache)
}

func commandMapBack(mainConfig *config, cache *pokecache.Cache) error {
	return commandMap(mainConfig, true, cache)
}