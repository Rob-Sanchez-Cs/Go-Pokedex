package main

import (
	"fmt"
	"os"
	"errors"
	"github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokecache"
	"github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokedex"
	"math/rand"
)

func commandExit(mainConfig *config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(mainConfig *config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, parameter string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cliCommand := range getCommands() {
		fmt.Printf("%v: %v\n\n", cliCommand.name, cliCommand.description)
	}

	return nil
}

func commandMap(mainConfig *config, usePreviousUrl bool, cache *pokecache.Cache, pokedex *pokedex.Pokedex, parameter string) error {
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

func commandMapNormal(mainConfig *config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, parameter string) error {
	return commandMap(mainConfig, false, cache, pokedex, parameter)
}

func commandMapBack(mainConfig *config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, parameter string) error {
	return commandMap(mainConfig, true, cache, pokedex, parameter)
}

func commandExplore(mainConfig *config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, parameter string) error {
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

func commandCatch(mainConfig *config, cache *pokecache.Cache, myPokedex *pokedex.Pokedex, parameter string) error {
	
	var pokemon pokedex.Pokemon
	err := getPokemon(parameter, &pokemon)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at "+parameter+"...")
	
	var catchChance int

	if pokemon.BaseExperience > 100 {
		catchChance = 50
	} else {
		catchChance = 75
	}
	randomRoll := rand.Intn(100)

	if catchChance >= randomRoll {
		fmt.Println(parameter+" was caught!")
		_, found := myPokedex.Get(parameter)
		if !found {
			myPokedex.Add(parameter, pokemon)
		}
	} else {
		fmt.Println(parameter+" escaped!")
	}


	return nil
}

func commandInspect(mainConfig *config, cache *pokecache.Cache, myPokedex *pokedex.Pokedex, parameter string) error {
	pokemon, found := myPokedex.Get(parameter)
	if !found {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name: "+parameter)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats{
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, myType := range pokemon.Types{
		fmt.Printf("  - %v\n", myType.Type.Name)
	}
	return nil
}

func commandPokedex(mainConfig *config, cache *pokecache.Cache, myPokedex *pokedex.Pokedex, parameter string) error {

	fmt.Println("Your Pokedex:")
	for _, key := range myPokedex.Keys() {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}