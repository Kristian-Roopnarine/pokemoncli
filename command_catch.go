package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	pokeapi "github.com/Kristian-Roopnarine/pokemoncli/internal/pokeapi"
)

func commandCatch(c *Config, arg string) error {
	url := fmt.Sprintf("%v/%v/%v", pokeapi.RootUrl, pokeapi.PokemonApiUrl, arg)
	resp, err := pokeapi.Get[pokeapi.PokemonResponse](url, c.Cache)
	if err != nil {
		return errors.New(err.Error())
	}

	if _, ok := c.Pokedex.Get(arg); ok {
		fmt.Printf("%v is already in your pokdex", arg)
		return nil
	}
	isCaught := throwPokeball(resp)
	if isCaught {
		c.Pokedex.Add(resp)
		fmt.Printf("%v was caught!\n", arg)
		fmt.Printf("You may not inspect it with the inspect command.\n")
	} else {
		fmt.Printf("%v escaped!\n", arg)
	}
	return nil
}

func throwPokeball(pokemon pokeapi.PokemonResponse) bool {
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	// going to use function 1 - e^(-x)
	k := 0.005
	baseChance := 1.0 - math.Exp(k*float64(-1*pokemon.BaseExperience))
	fmt.Printf("Base chance to catch %v : %f\n", pokemon.Name, baseChance)
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	chanceuuu := r.Float64()
	fmt.Printf("Random number : %f\n", chanceuuu)
	return chanceuuu >= baseChance
}
