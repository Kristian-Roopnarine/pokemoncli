package main

import (
	"fmt"
	pokeapi "github.com/Kristian-Roopnarine/pokemoncli/pokeapi"
)

func commandMap(config *Config) error {
	currentUrl := config.Next
	resp, err := pokeapi.Get(config.Next)
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	config.Next = resp.Next
	if config.Previous == nil {
		config.Previous = &currentUrl
	} else {
		config.Previous = resp.Previous
	}
	return err
}
