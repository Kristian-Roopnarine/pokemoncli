package main

import "fmt"

func commandPokedex(c *Config, arg string) error {
	if len(c.Pokedex.CaughtPokemon) == 0 {
		fmt.Println("Your pokedex is empty..")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for key, _ := range c.Pokedex.CaughtPokemon {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}
