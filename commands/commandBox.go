package commands

import (
	"fmt"
)

func commandBox(cfg *config, args []string) error {
	if len(cfg.PokeKeys) == 0 {
		return fmt.Errorf("Try catching a Pokemon!")
	}
	if len(cfg.PokeKeys) == 1 {
		return fmt.Errorf("You must have at least one Pokemon in your party.")
	}
	if len(args) == 0 {
		return fmt.Errorf("Please give a Pokemon to send to the box")
	}
	changed := false
	for i, name := range cfg.PokeKeys {
		if name == args[0] {
			changed = true
			cfg.PokeKeys = append(cfg.PokeKeys[:i], cfg.PokeKeys[i+1:]...)
			delete(cfg.Party, name)
		}
	}
	if !changed {
		return fmt.Errorf("This Pokemon is not in your Party!")
	}
	return nil
}