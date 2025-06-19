package main

import (
	"fmt"
)

func commandMoves(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please put a Pokemon here to check its moves.")
	}
	mon := args[0]
	pokemon, ok := cfg.Pokedex[mon]
	if !ok {
		return fmt.Errorf("You have not caught %v!", mon)
	}
	for i, move := range pokemon.Moves {
		fmt.Printf("%v: %v\n", i+1, move)
	}
	return nil
}