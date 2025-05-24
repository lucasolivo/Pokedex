package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	dex := cfg.Pokedex
	if len(dex) == 0 {
		return fmt.Errorf("Your pokedex is empty, try catching some pokemon!")
	}
	fmt.Println("Your pokedex:")
	for name, mon := range dex {
		fmt.Printf("  - %v Level %v\n", name, mon.Level)
	}
	return nil
}