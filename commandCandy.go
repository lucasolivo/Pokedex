package main

import "fmt"

func commandCandy(cfg *config, args []string) error {
	if len(args) == 0{
		return fmt.Errorf("Please give a Pokemon to level up!")
	}
	mon, ok := cfg.Pokedex[args[0]]
	if !ok {
		return fmt.Errorf("You haven't caught %v yet!", args[0])
	}
	if mon.Level == 100 {
		fmt.Printf("Your %v is already max level!\n", mon.Name)
		return nil
	} else {
		mon.Level += 1
		fmt.Printf("Your %v is now level %v\n", mon.Name, mon.Level)
		cfg.Pokedex[args[0]] = mon
		return nil
	}
}