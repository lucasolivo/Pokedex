package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/lucasolivo/Pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error // pointer to config allows us to directly mutate the config struct instead of making a copy
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

type LocationAreaResponse struct {
    Count    int    `json:"count"`
    Next     *string `json:"next"`
    Previous *string `json:"previous"`
    Results  []LocationArea `json:"results"`
}

type LocationArea struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

// get the lowercase words of each string input
func cleanInput(text string) []string{
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func startRepl() {
	cfg := &config{}
	cache := pokecache.NewCache(30 * time.Second)

	// create commands map, initialized with exit
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	
	// Create a closure for the help command
	helpCallback := func(cfg *config) error {
		return commandHelp(commands)
	}

	// Closure binding cache for the map command
	mapCallback := func(cfg *config) error {
		return commandMap(cfg, cache) // Pass the cache into commandMap
	}

	// Closure binding cache for the map back command
	mapbCallback := func(cfg *config) error {
		return commandMapb(cfg, cache) // Assume commandMapb supports cache
	}

	// add the map command
	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next map of the Pokeworld",
		callback:    mapCallback,
	}

	// add the map back command
	commands["mapb"] = cliCommand {
		name: "mapb",
		description: "Displays the previous page of the map of the Pokeworld",
		callback: mapbCallback,
	}
	
	// Add the help command
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    helpCallback,
	}

	// Add the explore command

	scanner := bufio.NewScanner(os.Stdin) //create a scanner
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			cleaned := cleanInput(userInput)
			command := cleaned[0] //the command should be the first word
			cmd, ok := commands[command]
			if ok {
				err := cmd.callback(cfg) 
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
