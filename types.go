package main

import "github.com/Rob-Sanchez-Cs/Go-Pokedex/internal/pokecache"


type cliCommand struct {
	name        string
	description string
	callback    func(mainConfig *config, cache *pokecache.Cache) error
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

type config struct {
	Next string
	Previous string
}