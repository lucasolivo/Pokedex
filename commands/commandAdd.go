package commands

import (
	"fmt"
)

func commandAdd(cfg *config, args []string) error {
	if len(cfg.Party) == 0 {
		return fmt.Errorf("Please catch some Pokemon first!")
	}
	if len(cfg.Party) >= 6 {
		return fmt.Errorf("Your party is full.")
	}
	name := args[0]
	mon, ok := cfg.Pokedex[name]
	if !ok {
		return fmt.Errorf("You have not caught %v yet!", name)
	}
	_, ok = cfg.Party[name]
	if ok {
		return fmt.Errorf("%v is already in your party!", name)
	}
	cfg.PokeKeys = append(cfg.PokeKeys, name)
	cfg.Party[name] = mon
	return nil
}