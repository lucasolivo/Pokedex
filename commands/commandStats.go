package commands

import (
	"fmt"
)

func commandStats(cfg *config, args []string) error {
	if len(cfg.Pokedex) == 0 {
		return fmt.Errorf("You don't have any Pokemon!")
	}
	var mon Pokemon
	if len(args) == 0 {
		mon = cfg.Party[cfg.PokeKeys[0]]
	} else {
		var ok bool
		mon, ok = cfg.Pokedex[args[0]]
		if !ok {
			return fmt.Errorf("You do not have this Pokemon!")
		}
	}
	for name, val := range mon.CurStats {
		fmt.Printf("%v: %v\n", name, val)
	}
	return nil
}