package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
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


		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg)
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
	callback    func(*config) error
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
			name:        "map",
			description: "Disaply the previous 20 map areas",
			callback:    commandMap,
		},
	}
}