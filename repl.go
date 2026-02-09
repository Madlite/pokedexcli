package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Madlite/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		commandArgs := input[1:]

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, commandArgs)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Disaply the next 20 map areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Disaply the previous 20 map areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Disaply the previous 20 map areas",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throw pokeball to try catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display pokemon entry in pokedex",
			callback:    commandInspect,
		},
	}
}
