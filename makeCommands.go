package main

import (
	"github.com/lucasolivo/Pokedex/internal/pokecache"
)

func makeCommands(cfg *config, cache *pokecache.Cache) map[string]cliCommand {
	commands := map[string]cliCommand{}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func(cfg *config, args []string) error {
			return commandHelp(commands)
		},
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next map of the Pokeworld",
		callback: func(cfg *config, args []string) error {
			return commandMap(cfg, cache)
		},
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous map of the Pokeworld",
		callback: func(cfg *config, args []string) error {
			return commandMapb(cfg, cache)
		},
	}

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Displays pokemon that can be found at a location",
		callback: func(cfg *config, args []string) error {
			return commandExplore(cfg, cache, args)
		},
	}

	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Throws a pokeball at a Pokemon",
		callback: func(cfg *config, args []string) error {
			return commandCatch(cfg, cache, args)
		},
	}

	commands["ability"] = cliCommand{
		name:        "ability",
		description: "Generates a random ability (add number to get multiple)",
		callback: func(cfg *config, args []string) error {
			return commandAbility(cfg, args)
		},
	}

	commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Displays info about a Pokemon in your Pokedex",
		callback:    commandInspect,
	}

	commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Lists out all Pokémon in your Pokedex",
		callback:    commandPokedex,
	}

	commands["candy"] = cliCommand{
		name:        "candy",
		description: "Gives one rare candy to a Pokemon of your choosing",
		callback:    commandCandy,
	}

	commands["party"] = cliCommand{
		name:        "party",
		description: "Displays your current Pokemon party",
		callback: func(cfg *config, args []string) error {
			return commandParty(cfg)
		},
	}

	commands["reset"] = cliCommand{
		name:        "reset",
		description: "Resets the Pokedex and party for a fresh start",
		callback: func(cfg *config, args []string) error {
			return commandReset(cfg)
		},
	}

	commands["encounter"] = cliCommand{
		name: "encounter",
		description: "Begins an encounter with the specified Pokemon or generates a random Pokemon to encounter if none is found.",
		callback: func(cfg *config, args []string) error {
			return commandEncounter(cfg, cache, args)
		},
	}

	return commands
}
