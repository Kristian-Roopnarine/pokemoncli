package main

import (
	"fmt"
	pokeapi "github.com/Kristian-Roopnarine/pokemoncli/pokeapi"
)

func commandMap(config *Config) error {
	resp, err := pokeapi.Get(config.Next)
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	config.Next = resp.Next
	config.Previous = resp.Previous
	return err
}
