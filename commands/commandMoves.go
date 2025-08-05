package commands

import (
	"fmt"
)

func commandMoves(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please put a Pokemon here to check its moves.")
	}
	mon := args[0]
	pokemon, ok := cfg.Pokedex[mon]
	if !ok {
		return fmt.Errorf("You have not caught %v!", mon)
	}
	moveData := pokemon.Movedata
	for i, move := range pokemon.Moves {
		fmt.Printf("%v: %v     Power: %v, Accuracy: %v, Type: %v, Damage Type: %v\n", i+1, move, 
		moveData[move].Power, moveData[move].Accuracy, moveData[move].Poketype, moveData[move].Damagetype)
	}
	return nil
}