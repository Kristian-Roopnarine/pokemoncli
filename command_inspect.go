package main

import "fmt"

func commandInspect(c *Config, arg string) error {
	poke, ok := c.Pokedex.Get(arg)
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", poke.Name)
	fmt.Printf("Height: %v\n", poke.Height)
	fmt.Println("Stat: ")
	for _, data := range poke.Stats {
		fmt.Printf(" - %v: %v", data.Stat.Name, data.BaseStat)
		fmt.Println()
	}
	fmt.Println("Types: ")
	for _, data := range poke.Types {
		fmt.Printf(" - %v", data.Type.Name)
		fmt.Println()
	}
	return nil
}
