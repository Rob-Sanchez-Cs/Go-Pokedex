package main

import (
	"github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokecache"
	"github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokedex"
)


type cliCommand struct {
	name        string
	description string
	callback    func(mainConfig *config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, parameter string) error
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type getMapResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type getExploreResponse struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             LocationArea      	   `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod EncounterMethod       `json:"encounter_method"`
	EncounterMethodVersionDetails  []EncounterMethodVersionDetails 	  `json:"version_details"`
}

type EncounterMethodVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Name struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource            `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetails   `json:"version_details"`
}

type PokemonEncounterVersionDetails struct {
	MaxChance       int                     `json:"max_chance"`
	Version         NamedAPIResource        `json:"version"`
	EncounterDetails []EncounterDetails      `json:"encounter_details"`
}

type EncounterDetails struct {
	Chance          int                `json:"chance"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	MaxLevel        int                `json:"max_level"`
	Method          NamedAPIResource   `json:"method"`
	MinLevel        int                `json:"min_level"`
}

type config struct {
	Next string
	Previous string
}