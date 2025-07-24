package main

import (
//	"bufio"
	"fmt"
//	"os"
	"strings"
	"time"
	"github.com/lucasolivo/Pokedex/internal/pokecache"
	"github.com/chzyer/readline"
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
	NickName      string
    BaseExperience int
    Height        int
    Weight        int
    Stats         map[string]int  // For storing stats like "hp": 40
    Types         []string        // For storing types like ["normal", "flying"]
	Level         int
	CurStats      map[string]int // Current stats of the pokemon, calculated based on level
	CurHp         int
	Moves         []string  // Current learned moves of the pokemon. 4 Maximum slots.
	Movedata      map[string]Move // The actual data about a particular move. Currently only dealing with power
	Ability       string
	Learnset      map[string]Pokemove
}

type Pokemove struct {
	LevelUp int
	url string
}

type Move struct {
	Power string
	Accuracy string
	Poketype string
	Damagetype string
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
			Pokedex:  make(map[string]Pokemon),
			Party:    make(map[string]Pokemon),
			PokeKeys: []string{},
		}
	}
	cache := pokecache.NewCache(30 * time.Second)
	commands := makeCommands(cfg, cache)

	// Set up interactive prompt with readline
	rl, err := readline.New("Pokedex > ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // Handles Ctrl+D and other exit signals
			break
		}

		cleaned := cleanInput(line)
		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
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
