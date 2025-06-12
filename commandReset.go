package main

import (
	"fmt"
)

func commandReset(cfg *config) error {
	cfg.Pokedex = make(map[string]Pokemon)
	cfg.Party = make(map[string]Pokemon)
	cfg.PokeKeys = []string{}
	fmt.Println("Your Pokedex has been reset.")
	return nil
}