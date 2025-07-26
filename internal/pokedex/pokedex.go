package pokedex


import (
	"sync"
)

type Pokedex struct {
	mu sync.RWMutex
	data map[string]Pokemon
}

type StatEntry struct {
    BaseStat int              `json:"base_stat"`
    Effort   int              `json:"effort"`
    Stat     NamedAPIResource `json:"stat"`
}

type TypeEntry struct {
    Slot int              `json:"slot"`
    Type NamedAPIResource `json:"type"`
}

type Pokemon struct {
    Height int         `json:"height"`
    Weight int         `json:"weight"`
    Stats  []StatEntry `json:"stats"`
    Types  []TypeEntry `json:"types"`
	BaseExperience int `json:"base_experience"`
}

type NamedAPIResource struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

func NewPokedex() *Pokedex {
	c := Pokedex{
		data: make(map[string]Pokemon),
	}

	return &c

}

func (c *Pokedex) Add(key string, pokemon Pokemon) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = pokemon
}

func (c *Pokedex) Get(key string) (Pokemon, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	pokemon, found := c.data[key]

	if !found {
		return Pokemon{}, false
	}

	return pokemon, true
}
