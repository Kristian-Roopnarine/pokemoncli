package main

import (
	"errors"
	"fmt"
	pokeapi "github.com/Kristian-Roopnarine/pokemoncli/internal/pokeapi"
	"strings"
)

func commandExplore(c *Config, arg string) error {
	url := fmt.Sprintf("%v/%v/%v", pokeapi.RootUrl, pokeapi.LocationAreaApi, arg)
	fmt.Printf("Exploring %v...\n", arg)
	resp, err := pokeapi.Get[pokeapi.LocationAreaResponse](url, c.Cache)
	if err != nil {
		return errors.New(err.Error())
	}

	var outputStr strings.Builder
	if len(resp.PokemonEncounters) > 0 {
		outputStr.WriteString(fmt.Sprintf("Found Pokemon:\n"))
		for _, result := range resp.PokemonEncounters {
			outputStr.WriteString(fmt.Sprintf(" - %v\n", result.Pokemon.Name))
		}
	} else {
		outputStr.WriteString(fmt.Sprintf("No pokemon found in %v", arg))
	}
	fmt.Println(outputStr.String())
	return nil
}
