package pokedex

import (
	"sync"

	"github.com/Kristian-Roopnarine/pokemoncli/internal/pokeapi"
)

type Pokedex struct {
	CaughtPokemon map[string]pokeapi.PokemonResponse
	mu            *sync.RWMutex
}

func NewPokedex() Pokedex {
	return Pokedex{CaughtPokemon: make(map[string]pokeapi.PokemonResponse), mu: &sync.RWMutex{}}
}

func (p Pokedex) Add(pokemon pokeapi.PokemonResponse) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.CaughtPokemon[pokemon.Name] = pokemon
	return nil
}

func (p Pokedex) Get(name string) (pokeapi.PokemonResponse, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	poke, ok := p.CaughtPokemon[name]
	if !ok {
		return pokeapi.PokemonResponse{}, false
	}
	return poke, ok
}
