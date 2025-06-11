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
	cfg.Pokedex[args[0]] = level(cfg, mon, args[0])
	return nil
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

func level(cfg *config, mon Pokemon, name string) Pokemon {
	if mon.Level == 100 {
		fmt.Printf("Your %v is already max level!\n", mon.Name)
		return mon
	} else {
		//speciesURL := "https://pokeapi.co/api/v2/pokemon-species/" + strings.ToLower(mon.Name)
		mon.Level += 1
		fmt.Printf("Your %v is now level %v\n", mon.Name, mon.Level)
		cfg.Pokedex[name] = mon
		return mon
	}
}