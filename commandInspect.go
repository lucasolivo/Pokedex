package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please submit a Pokemon name")
	}
	pokemonName := args[0]

	mon, ok := cfg.Pokedex[pokemonName]

	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", mon.Name)
	fmt.Printf("Height: %v\n", mon.Height)
	fmt.Printf("Weight: %v\n", mon.Weight)
	fmt.Printf("Stats:\n")
	for stat, val := range mon.Stats {
		fmt.Printf("  -%v: %v\n", stat, val)
	}
	fmt.Printf("Types:\n")
	for _, typ := range mon.Types {
		fmt.Printf("  - %v\n", typ)
	}
	return nil
}