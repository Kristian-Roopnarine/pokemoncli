package main

import (
	"errors"
	"fmt"
	pokeapi "github.com/Kristian-Roopnarine/pokemoncli/internal/pokeapi"
)

func commandMap(config *Config, arg string) error {
	resp, err := pokeapi.Get[pokeapi.LocationResponse](config.Next, config.Cache)
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	config.Next = resp.Next
	config.Previous = resp.Previous
	return err
}

func commandMapb(config *Config, arg string) error {
	if config.Previous == nil {
		return errors.New("No Previous locations to show.")
	}
	resp, err := pokeapi.Get[pokeapi.LocationResponse](*config.Previous, config.Cache)
	if err != nil {
		return err
	}
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	config.Next = resp.Next
	config.Previous = resp.Previous
	return nil

}
