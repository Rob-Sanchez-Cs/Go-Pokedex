package main

type cliCommand struct {
	name        string
	description string
	callback    func(mainConfig *config) error
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