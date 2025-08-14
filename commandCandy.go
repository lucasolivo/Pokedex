package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"net/http"
	"encoding/json"
	"math/rand"
	"io"
)

func commandCandy(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please give a Pokemon to level up!")
	}

	mon, ok := cfg.Pokedex[args[0]]
	if !ok {
		return fmt.Errorf("You haven't caught %v yet!", args[0])
	}

	// Let level() return both the updated mon and its name after possible evolution
	newName, updatedMon := level(cfg, mon, args[0])

	// Update Pokédex
	cfg.Pokedex[newName] = updatedMon
	if newName != args[0] {
		delete(cfg.Pokedex, args[0]) // remove old name entry
	}

	// Keep total exp synced
	updatedMon.TotalExp = expChart[updatedMon.ExpGroup][updatedMon.Level]
	cfg.Pokedex[newName] = updatedMon

	// Update party if they’re in it
	for key := range cfg.Party {
		if key == args[0] || key == newName {
			cfg.Party[newName] = updatedMon
			if newName != args[0] {
				delete(cfg.Party, args[0])
			}
		}
	}

	return nil
}


func level(cfg *config, mon Pokemon, name string) (string, Pokemon) {
	if mon.Level == 100 {
		fmt.Printf("Your %v is already max level!\n", mon.Name)
		return mon.Name, mon
	} else {
		mon.Level += 1
		fmt.Printf("Your %v is now level %v\n", mon.Name, mon.Level)
		for moveName, move := range mon.Learnset {
			if move.LevelUp == mon.Level {
				var err error
				mon, err = addMoveData(mon, moveName)
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
		var toEvolve []string
		minLvl := 100
		for evoName, num := range mon.EvolutionLevels{
			if num <= mon.Level && minLvl >= num {
				if minLvl > num {
					toEvolve = toEvolve[:0]
				}
				toEvolve = append(toEvolve, evoName)
				minLvl = num
			}
		}
		if len(toEvolve) > 0{
			if len(toEvolve) > 1 {
				fmt.Println("You have a split evolution! please select the name or number of your choice!")
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Printf("0: Do not evolve\n")
				count := 1
				for evolutionName := range toEvolve {
					fmt.Printf("%v: %v", count, evolutionName)
					count += 1
				}
				for {
					if scanner.Scan() {
						userInput := scanner.Text()
						cleaned := cleanInput(userInput)
						if len(cleaned) == 0 {
							continue
						}
						answer := cleaned[0]
						changeTo := ""
						if answer == "0" || answer == "do not evolve"{
							changeTo = "no"
						}

						for c := 1; c < count; c++ {
							if answer == strconv.Itoa(c) || answer == toEvolve[c-1] {
								changeTo = toEvolve[c-1]
								break
							}
						}

						if changeTo == ""{
							continue
						}
						if changeTo == "no"{
							break
						}
						_, kk := cfg.Pokedex[changeTo]
						if kk {
							fmt.Printf("You already have a %v, would you like to replace him by evolving? (Yes or No)\n", changeTo)
							secondScanner := bufio.NewScanner(os.Stdin)
							for {
									if secondScanner.Scan() {
										input2 := secondScanner.Text()
										clean2 := cleanInput(input2)
										if len(clean2) == 0 {
											continue
										}
										answer2 := clean2[0]
										if answer2 == "yes" {
											delete(cfg.Pokedex, changeTo)
											for i, pokeKey := range cfg.PokeKeys {
												if pokeKey == changeTo {
													cfg.PokeKeys = append(cfg.PokeKeys[:i], cfg.PokeKeys[i+1:]...)
													delete(cfg.Party, changeTo)
												}
											}
											mon, _ = evolve(cfg, mon, changeTo)
											break
										} else {
											break
										}
									}
								}
							break
						} else {
							mon, _ = evolve(cfg, mon, changeTo)
						}
					}
				}
			} else {
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Printf("Your Pokemon is evolving into %v! Would you like him to grow? (Yes or no)\n", toEvolve[0])
				for {
					if scanner.Scan(){
						userInput := scanner.Text()
						cleaned := cleanInput(userInput)
						if len(cleaned) == 0 {
							continue
						}
						answer := cleaned[0]
						if answer == "yes"{
							_, ok := cfg.Pokedex[toEvolve[0]]
							if ok {
								fmt.Printf("You already have a %v, would you like to replace him by evolving? (Yes or No)\n", toEvolve[0])
								secondScanner := bufio.NewScanner(os.Stdin)
								for {
									if secondScanner.Scan() {
										input2 := secondScanner.Text()
										clean2 := cleanInput(input2)
										if len(clean2) == 0 {
											continue
										}
										answer2 := clean2[0]
										if answer2 == "yes" {
											delete(cfg.Pokedex, toEvolve[0])
											for i, pokeKey := range cfg.PokeKeys {
												if pokeKey == toEvolve[0] {
													cfg.PokeKeys = append(cfg.PokeKeys[:i], cfg.PokeKeys[i+1:]...)
													delete(cfg.Party, toEvolve[0])
												}
											}
											mon, _ = evolve(cfg, mon, toEvolve[0])
											answer = "no"
											break
										} else {
											answer = "no"
											break
										}
									}
								}
							} else {
								mon, _ = evolve(cfg, mon, toEvolve[0])
								break
							}
						}
						if answer == "no"{
							break
						}
					}
				}
			}
		}
		// Recalculate stats for the new level
		newStats := statCalculator(mon.Stats, mon.Level) // Or pass base stats instead if that's how your function works

		// Preserve HP % on level up
		hpPercent := float64(mon.CurHp) / float64(mon.CurStats["hp"])
		mon.CurHp = int(hpPercent * float64(newStats["hp"]))
		mon.CurStats = newStats

		return mon.Name, mon
	}
}

func evolve(cfg *config, mon Pokemon, pokemonName string) (Pokemon, error) {
	if mon.Name == pokemonName {
		fmt.Println("Pokemon evolving into itself.")
		return mon, nil
	}

	prevName := mon.Name

	// --- Fetch new Pokemon data ---
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	res, err := http.Get(url)
	if err != nil {
		return mon, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return mon, fmt.Errorf("Pokemon %s not found", pokemonName)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return mon, err
	}

	var pokemonData map[string]interface{}
	if err := json.Unmarshal(body, &pokemonData); err != nil {
		return mon, err
	}

	// Parse basic fields
	id := int(pokemonData["id"].(float64))
	baseExperience := int(pokemonData["base_experience"].(float64))
	height := int(pokemonData["height"].(float64))
	weight := int(pokemonData["weight"].(float64))

	// Handle ability — keep old if possible
	abilities := pokemonData["abilities"].([]interface{})
	ability := mon.Ability
	hasOldAbility := false
	for _, rawAbility := range abilities {
		name := rawAbility.(map[string]interface{})["ability"].(map[string]interface{})["name"].(string)
		if name == mon.Ability {
			hasOldAbility = true
			break
		}
	}
	if !hasOldAbility {
		toChoose := 1
		if len(abilities) > 1 {
			toChoose = 1 + rand.Intn(len(abilities)-1)
		}
		for _, rawAbility := range abilities {
			abilityEntry := rawAbility.(map[string]interface{})
			slot := int(abilityEntry["slot"].(float64))
			if slot == toChoose {
				ability = abilityEntry["ability"].(map[string]interface{})["name"].(string)
				break
			}
		}
	}

	// --- Get growth rate ---
	speciesURL := "https://pokeapi.co/api/v2/pokemon-species/" + pokemonName
	res, err = http.Get(speciesURL)
	if err != nil {
		return mon, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return mon, fmt.Errorf("Species info for %s not found", pokemonName)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return mon, err
	}

	var speciesData map[string]interface{}
	if err := json.Unmarshal(body, &speciesData); err != nil {
		return mon, err
	}
	growthRateName := speciesData["growth_rate"].(map[string]interface{})["name"].(string)

	// --- Learnset ---
	moves := pokemonData["moves"].([]interface{})
	learnSet := make(map[string]Pokemove)
	for _, rawMove := range moves {
		moveEntry := rawMove.(map[string]interface{})
		curMove := moveEntry["move"].(map[string]interface{})
		details := moveEntry["version_group_details"].([]interface{})
		group_to_use := details[len(details)-1].(map[string]interface{})
		method := group_to_use["move_learn_method"].(map[string]interface{})

		if method["name"].(string) != "level-up" {
			learnSet[curMove["name"].(string)] = Pokemove{LevelUp: 101, url: curMove["url"].(string)}
		} else {
			learnSet[curMove["name"].(string)] = Pokemove{
				LevelUp: int(group_to_use["level_learned_at"].(float64)),
				url:     curMove["url"].(string),
			}
		}
	}

	// --- Stats ---
	stats := make(map[string]int)
	statsArray := pokemonData["stats"].([]interface{})
	for _, statItem := range statsArray {
		statMap := statItem.(map[string]interface{})
		statValue := int(statMap["base_stat"].(float64))
		statName := statMap["stat"].(map[string]interface{})["name"].(string)
		stats[statName] = statValue
	}

	// --- Types ---
	var types []string
	typesArray := pokemonData["types"].([]interface{})
	for _, typeItem := range typesArray {
		typeName := typeItem.(map[string]interface{})["type"].(map[string]interface{})["name"].(string)
		types = append(types, typeName)
	}

	// --- Preserve HP percent ---
	hpPercent := float64(mon.CurHp) / float64(mon.CurStats["hp"])
	statVals := statCalculator(stats, mon.Level)
	mon.CurHp = int(hpPercent * float64(statVals["hp"]))

	// --- Update mon fields ---
	if mon.NickName == mon.Name {
		mon.NickName = pokemonName
	}
	mon.CurStats = statVals
	mon.Stats = stats
	mon.Name = pokemonName
	mon.Ability = ability
	mon.ID = id
	mon.BaseExperience = baseExperience
	mon.Height = height
	mon.Weight = weight
	mon.Types = types
	mon.Learnset = learnSet
	mon.ExpGroup = growthRateName
	delete(mon.EvolutionLevels, pokemonName)

	// --- Cleanup old name from all collections ---
	
	for i, k := range cfg.PokeKeys {
		if k == prevName {
			cfg.PokeKeys[i] = pokemonName
		}
	}

	// --- Add new mon to collections ---
	cfg.Pokedex[pokemonName] = mon
	if _, inParty := cfg.Party[prevName]; inParty {
		cfg.Party[pokemonName] = mon
	}
	delete(cfg.Pokedex, prevName)
	delete(cfg.Party, prevName)

	fmt.Printf("Congrats! Your %v evolved into %v\n", prevName, pokemonName)
	return mon, nil
}
