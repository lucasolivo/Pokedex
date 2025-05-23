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
	callback    func(*config, []string) error // pointer to config allows us to directly mutate the config struct instead of making a copy
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	Pokedex map[string]Pokemon
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

type PokemonEncounter struct {
    Pokemon struct {
        Name string `json:"name"`
    } `json:"pokemon"`
}

type PokemonAreaResponse struct {
    PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type Pokemon struct {
    ID            int
    Name          string
    BaseExperience int
    Height        int
    Weight        int
    Stats         map[string]int  // For storing stats like "hp": 40
    Types         []string        // For storing types like ["normal", "flying"]
}

// get the lowercase words of each string input
func cleanInput(text string) []string{
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func startRepl() {
	fmt.Println("Welcome to the Pokedex! Input 'help' for a list of commands!")
	cfg := &config{
        // Your existing initialization
        Pokedex: make(map[string]Pokemon),
    }
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
	helpCallback := func(cfg *config, args []string) error {
		return commandHelp(commands)
	}

	// Closure binding cache for the map command
	mapCallback := func(cfg *config, args []string) error {
		return commandMap(cfg, cache) // Pass the cache into commandMap
	}

	// Closure binding cache for the map back command
	mapbCallback := func(cfg *config, args []string) error {
		return commandMapb(cfg, cache) // Assume commandMapb supports cache
	}

	// Closure binding cache for the explore command
	exploreCallback := func(cfg *config, args []string) error {
		return commandExplore(cfg, cache, args) // Pass the cache into commandExplore
	}

	// Closure binding cache for the catch command
	catchCallback := func(cfg *config, args []string) error {
		return commandCatch(cfg, cache, args)
	}

	abilityCallback := func(cfg *config, args []string) error {
		return commandAbility(cfg, args)
	}

	// add the map command
	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next map of the Pokeworld",
		callback:    mapCallback,
	}


	// add a command to generate random abilities.
	commands["ability"] = cliCommand{
		name: "ability",
		description: "Generates a random ability, multiple (up to 10) if a number is added.",
		callback: abilityCallback,
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
	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Displays pokemon that can be found at a location",
		callback:    exploreCallback,
	}

	// Add the catch command
	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Throws a pokeball at a Pokemon",
		callback:    catchCallback,
	}

	// Add the inspect command
	commands["inspect"] = cliCommand{
		name: "catch",
		description: "Inspects a pokemon the user asks for, displaying stats if in the pokedex",
		callback: commandInspect,
	}

	// Add the pokedex command
	commands["pokedex"] = cliCommand {
		name: "pokedex",
		description: "Lists out all the pokemon that have been added to your Pokedex",
		callback: commandPokedex,
	}

	scanner := bufio.NewScanner(os.Stdin) //create a scanner
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			cleaned := cleanInput(userInput)
			command := cleaned[0] //the command should be the first word
			args := cleaned[1:]
			cmd, ok := commands[command]
			if ok {
				err := cmd.callback(cfg, args) 
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
