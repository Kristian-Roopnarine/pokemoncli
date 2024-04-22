package main

import (
	"errors"
	"fmt"
	pokeapi "github.com/Kristian-Roopnarine/pokemoncli/pokeapi"
)

func commandMapb(config *Config) error {
	if config.Previous == nil {
		return errors.New("No Previous locations to show.")
	}
	config.Next = *config.Previous
	resp, err := pokeapi.Get(*config.Previous)
	if err != nil {
		return err
	}
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	config.Previous = resp.Previous
	return nil

}
