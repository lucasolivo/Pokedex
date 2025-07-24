package main

import (
	"fmt"
	"bufio"
	"strconv"
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
				mon, err := addMoveData(mon, moveName)
				if err != nil {
					fmt.Printf("Failed to add move data for %s: %v\n", moveName, err)
				}
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
						fmt.Printf("New move: %v     Power: %v, Accuracy: %v, Type: %v, Damage Type: %v\n\n", 
						moveName, mon.Movedata[moveName].Power, mon.Movedata[moveName].Accuracy, 
						mon.Movedata[moveName].Poketype, mon.Movedata[moveName].Damagetype)
						for i, thisMove := range mon.Moves {
							fmt.Printf("%v: %v     Power: %v, Accuracy: %v, Type: %v, Damage Type: %v\n", i+1, thisMove, 
							mon.Movedata[thisMove].Power, mon.Movedata[thisMove].Accuracy, 
							mon.Movedata[thisMove].Poketype, mon.Movedata[thisMove].Damagetype)
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
									strDex := strconv.Itoa(dex+1)
										if ToDelete == thisMove || ToDelete == strDex{
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