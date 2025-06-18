package main

import (
	"fmt"
)

func commandSwitch(cfg *config, args []string) error {
	if len(cfg.Party) < 2 {
		return fmt.Errorf("Not enough Pokemon in party.")
	}
	if len(args) < 2 {
		return fmt.Errorf("Please provide two pokemon in your party to switch.")
	}
	if args[0] == args[1] {
		return fmt.Errorf("Please provide two different pokemon.")
	}
	firstPokemon := args[0]
	secondPokemon := args[1]
	_, ok := cfg.Party[firstPokemon]
	if !ok {
		return fmt.Errorf("%v not in your party.", firstPokemon)
	}
	_, ok = cfg.Party[secondPokemon]
	if !ok {
		return fmt.Errorf("%v not in your party.", secondPokemon)
	}
	var first int
	var second int 
	for i, name := range cfg.PokeKeys {
		if name == firstPokemon {
			first = i
		}
		if name == secondPokemon {
			second = i
		}
	}
	cfg.PokeKeys[first] = secondPokemon
	cfg.PokeKeys[second] = firstPokemon
	return nil
}