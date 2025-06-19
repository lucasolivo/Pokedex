package main

import (
	"fmt"
	"bufio"
	"os"
)

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
}

func level(cfg *config, mon Pokemon, name string) Pokemon {
	if mon.Level == 100 {
		fmt.Printf("Your %v is already max level!\n", mon.Name)
		return mon
	} else {
		mon.Level += 1
		fmt.Printf("Your %v is now level %v\n", mon.Name, mon.Level)
		for moveName, move := range mon.Learnset {
			if move.LevelUp == mon.Level {
				alrLearned := false
				for _, nameOfMove := range mon.Moves {
					if nameOfMove == moveName{
						fmt.Printf("%v already knows %v.\n", name, moveName)
						alrLearned = true
					}
				}
				if !alrLearned {
					if len(mon.Moves) < 4 {
						mon.Moves = append(mon.Moves, moveName)
						fmt.Printf("%v learned %v!\n", name, moveName)
					} else {
						secondScan := bufio.NewScanner(os.Stdin)
						fmt.Println("You can only know up to 4 moves at a time, is there one you would Like to replace?")
						fmt.Printf("New move: %v\n", moveName)
						for i, thisMove := range mon.Moves {
							fmt.Printf("%v: %v\n", i+1, thisMove)
						}
						for {
							if secondScan.Scan() {
								hasDeleted := false
								userInput := secondScan.Text()
								cleaned := cleanInput(userInput)
								if len(cleaned) == 0 {
									continue
								}
								ToDelete := cleaned[0]
								if ToDelete == moveName {
									fmt.Printf("You did not learn %v\n", moveName)
									break
								}
								for dex, thisMove := range mon.Moves {
									if ToDelete == thisMove {
										hasDeleted = true
										mon.Moves = append(mon.Moves[:dex], mon.Moves[dex+1:]...)
										mon.Moves = append(mon.Moves, moveName)
										fmt.Printf("You forgot %v and learned %v\n", thisMove, moveName)
										break
									}
								}
								if !hasDeleted{
									fmt.Println("Please give a valid move to delete.")
								} else {
									break
								}
							}
						}
					}
			}
		}
		}
		return mon
	}
}