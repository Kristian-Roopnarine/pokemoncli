package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokecache "github.com/Kristian-Roopnarine/pokemoncli/internal/pokecache"
	pokedex "github.com/Kristian-Roopnarine/pokemoncli/internal/pokedex"
)

type Config struct {
	Next     string
	Previous *string
	Cache    pokecache.Cache
	Pokedex  pokedex.Pokedex
}

func startRepl() {
	var commandName string
	var argument string
	reader := bufio.NewScanner(os.Stdin)
	cache, err := pokecache.NewCache(5 * time.Minute)
	if err != nil {
		panic("Error creating Cache")
	}
	pokedex := pokedex.NewPokedex()
	locationConfig := Config{Next: "https://pokeapi.co/api/v2/location-area", Previous: nil, Cache: cache, Pokedex: pokedex}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName = words[0]
		if len(words) > 1 {
			argument = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&locationConfig, argument)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		fmt.Println("")
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config, arg string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get list of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get list of previous locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get list of pokemon in location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all captured pokemon",
			callback:    commandPokedex,
		},
	}
}
