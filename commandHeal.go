package main

import (
	"fmt"
)

func commandHeal(cfg *config) error {
	if len(cfg.Party) == 0 {
		return fmt.Errorf("You have no Pokemon to heal")
	}
	for i := range cfg.PokeKeys {
		key := cfg.PokeKeys[i]
		p := cfg.Party[key]
		p.CurHp = p.CurStats["hp"]
		cfg.Party[key] = p                           
	}
	fmt.Println("Your Pokemon are healed and ready for battle! Be sure to protect them!")
	return nil
}