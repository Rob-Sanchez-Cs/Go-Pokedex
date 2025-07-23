package main

import (
	"fmt"
	"os"
	"errors"
	"github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokecache"
)

func commandExit(mainConfig *config, cache *pokecache.Cache, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(mainConfig *config, cache *pokecache.Cache, parameter string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cliCommand := range getCommands() {
		fmt.Printf("%v: %v\n\n", cliCommand.name, cliCommand.description)
	}

	return nil
}

func commandMap(mainConfig *config, usePreviousUrl bool, cache *pokecache.Cache, parameter string) error {
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

func commandMapNormal(mainConfig *config, cache *pokecache.Cache, parameter string) error {
	return commandMap(mainConfig, false, cache, parameter)
}

func commandMapBack(mainConfig *config, cache *pokecache.Cache, parameter string) error {
	return commandMap(mainConfig, true, cache, parameter)
}

func commandExplore(mainConfig *config, cache *pokecache.Cache, parameter string) error {
	if parameter == "" {
		return errors.New("must pass a location with the explore command")
	}

	var exploreResponse getExploreResponse
	err := exploreLocation(&exploreResponse, cache, parameter)
	if err != nil {
		return err
	}
	fmt.Println("Exploring "+parameter+"...")
	if len(exploreResponse.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon are located in this area")
	} else {
		printPokemonInExploredLocation(exploreResponse)
	}
	
	return nil
}

func printPokemonInExploredLocation(exploreResponse getExploreResponse) {
	fmt.Println("Found Pokemon:")

	for _, encounter := range exploreResponse.PokemonEncounters {
		fmt.Println(" - "+encounter.Pokemon.Name)
	}
}