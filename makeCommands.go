package main

import (
	"github.com/lucasolivo/Pokedex/internal/pokecache"
	"github.com/lucasolivo/Pokedex/internal/commands"
)

func makeCommands(cfg *config, cache *pokecache.Cache) map[string]cliCommand {
	coms := map[string]cliCommand{}

	coms["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	coms["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func(cfg *config, args []string) error {
			return commands.commandHelp(commands)
		},
	}

	coms["map"] = cliCommand{
		name:        "map",
		description: "Displays the next map of the Pokeworld",
		callback: func(cfg *config, args []string) error {
			return commands.commandMap(cfg, cache)
		},
	}

	coms["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous map of the Pokeworld",
		callback: func(cfg *config, args []string) error {
			return commands.commandMapb(cfg, cache)
		},
	}

	coms["explore"] = cliCommand{
		name:        "explore",
		description: "Displays pokemon that can be found at a location",
		callback: func(cfg *config, args []string) error {
			return commands.commandExplore(cfg, cache, args)
		},
	}

	coms["catch"] = cliCommand{
		name:        "catch",
		description: "Throws a pokeball at a Pokemon",
		callback: func(cfg *config, args []string) error {
			return commands.commandCatch(cfg, cache, args)
		},
	}

	coms["ability"] = cliCommand{
		name:        "ability",
		description: "Generates a random ability (add number to get multiple)",
		callback: func(cfg *config, args []string) error {
			return commands.commandAbility(cfg, args)
		},
	}

	coms["inspect"] = cliCommand{
		name:        "inspect",
		description: "Displays info about a Pokemon in your Pokedex",
		callback:    commands.commandInspect,
	}

	coms["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Lists out all Pokémon in your Pokedex",
		callback:    commands.commandPokedex,
	}

	coms["candy"] = cliCommand{
		name:        "candy",
		description: "Gives one rare candy to level up a Pokemon of your choosing",
		callback:    comamnds.commandCandy,
	}

	coms["party"] = cliCommand{
		name:        "party",
		description: "Displays your current Pokemon party",
		callback: func(cfg *config, args []string) error {
			return commands.commandParty(cfg)
		},
	}

	coms["heal"] = cliCommand {
		name: "heal",
		description: "Heals your party Pokemon to their maximum HP",
		callback: func(cfg *config, args []string) error {
			return commands.commandHeal(cfg)
		},
	}

	coms["reset"] = cliCommand{
		name:        "reset",
		description: "Resets the Pokedex and party for a fresh start",
		callback: func(cfg *config, args []string) error {
			return commands.commandReset(cfg)
		},
	}

	coms["encounter"] = cliCommand{
		name: "encounter",
		description: "Begins an encounter with the specified Pokemon or generates a random Pokemon to encounter if none is found.",
		callback: func(cfg *config, args []string) error {
			return commands.commandEncounter(cfg, cache, args)
		},
	}

	coms["stats"] = cliCommand {
		name: "stats",
		description: "Lists out the individual stats of the Pokemon specified, if in Pokedex. Defaults to lead.",
		callback: func(cfg *config, args []string) error {
			return commands.commandStats(cfg, args)
		},
	}

	coms["box"] = cliCommand {
		name: "box",
		description: "Sends a Pokemon from your Party into your Box",
		callback: func(cfg *config, args []string) error {
			return commands.commandBox(cfg, args)
		},
	}

	coms["switch"] = cliCommand {
		name: "switch",
		description: "Switches the positions of two Pokemon in your party.",
		callback: func(cfg *config, args []string) error {
			return commands.commandSwitch(cfg, args)
		},
	}

	coms["add"] = cliCommand {
		name: "add",
		description: "Adds a Pokemon from your box into your party.",
		callback: func(cfg *config, args []string) error {
			return commands.commandAdd(cfg, args)
		},
	}

	coms["moves"] = cliCommand {
		name: "moves",
		description: "Displays the current Moveset of a Pokemon in your Box.",
		callback: func(cfg *config, args []string) error {
			return commands.commandMoves(cfg, args)
		},
	}

	coms["teach"] = cliCommand {
		name: "teach",
		description: "Teaches a move from the Pokemon's learnset to the specified Pokemon. (Form: teach 'pokemon' 'move')",
		callback: func(cfg *config, args []string) error {
			return commands.commandTeach(cfg, args)
		},
	}

	return coms
}
