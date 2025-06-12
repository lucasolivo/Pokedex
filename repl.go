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
	Party map[string]Pokemon
	PokeKeys []string
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
	Level         int
}

// get the lowercase words of each string input
func cleanInput(text string) []string{
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func startRepl() {
	fmt.Println("Welcome to the Pokedex! Input 'help' for a list of commands!")
	cfg, err := loadGame()
	if err != nil {
		cfg = &config{
			// Your existing initialization
			Pokedex: make(map[string]Pokemon),
			Party: make(map[string]Pokemon),
			PokeKeys: []string{},
		}
	}
	cache := pokecache.NewCache(30 * time.Second)

	// create commands map, initialized with exit
	commands := makeCommands(cfg, cache)

	scanner := bufio.NewScanner(os.Stdin) //create a scanner
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			cleaned := cleanInput(userInput)
			if len(cleaned) == 0 {
				continue
			}
			command := cleaned[0] //the command should be the first word
			args := cleaned[1:]
			cmd, ok := commands[command]
			if ok {
				err := cmd.callback(cfg, args) 
				if err != nil {
					fmt.Println(err)
				}
				err = saveGame(cfg)
				if err != nil {
					fmt.Println("Failed to save game state:", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
