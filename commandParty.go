package main

import (
	"fmt"
)

func commandParty(cfg *config) error {
	if len(cfg.Party) == 0 {
		return fmt.Errorf("You have no Pokemon in your party!")
	}
	for i := range cfg.PokeKeys{
		if i == 0{
			fmt.Printf("Lead: ")
		}
		mon := cfg.Party[cfg.PokeKeys[i]]
		if mon.NickName != mon.Name {
			fmt.Printf("%v the %v at level %v\n", mon.NickName, mon.Name, mon.Level)
		} else {
			fmt.Printf("%v at level %v\n", mon.Name, mon.Level)
		}
	}
	return nil
}