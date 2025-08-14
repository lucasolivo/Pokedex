package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

func commandTeach(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please input a Pokemon to teach a move to")
	}
	name := args[0]
	mon, ok := cfg.Pokedex[name]
	if !ok {
		return fmt.Errorf("You have not caught %v", name)
	}
	if len(args) == 1 {
		return fmt.Errorf("Please input a move to teach to %v", args[0])
	}
	move := args[1]
	_, ok = mon.Learnset[move]
	if !ok {
		return fmt.Errorf("%v cannot learn %v", name, move)
	}
	for _, nameOfMove := range mon.Moves {
		if nameOfMove == move {
			return fmt.Errorf("%v already knows %v.", name, move)
		}
	}
	mon, err := addMoveData(mon, move)
	if err != nil {
		return fmt.Errorf("Could not get Move Data")
	}
	if len(mon.Moves) < 4 {
		mon.Moves = append(mon.Moves, move)
		fmt.Printf("%v learned %v!\n", name, move)
	} else {
		secondScan := bufio.NewScanner(os.Stdin)
		fmt.Println("You can only know up to 4 moves at a time, is there one you would Like to replace?")
		fmt.Printf("New move: %v     Power: %v, Accuracy: %v, Type: %v, Damage Type: %v\n\n", 
						move, mon.Movedata[move].Power, mon.Movedata[move].Accuracy, 
						mon.Movedata[move].Poketype, mon.Movedata[move].Damagetype)
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
				if ToDelete == move {
					fmt.Printf("You did not learn %v\n", move)
					break
				}
				for dex, thisMove := range mon.Moves {
					strDex := strconv.Itoa(dex+1)
					if ToDelete == thisMove || ToDelete == strDex{
						hasDeleted = true
						mon.Moves = append(mon.Moves[:dex], mon.Moves[dex+1:]...)
						mon.Moves = append(mon.Moves, move)
						fmt.Printf("You forgot %v and learned %v\n", thisMove, move)
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

	cfg.Pokedex[name] = mon
	if _, inParty := cfg.Party[name]; inParty {
		cfg.Party[name] = mon
	}
	return nil
}