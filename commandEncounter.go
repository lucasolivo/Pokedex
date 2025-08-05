package main

import (
	"net/http"       
    "encoding/json"
	"fmt"
	"io"
    "github.com/lucasolivo/Pokedex/internal/pokecache"
	"math/rand"
	"bufio"
	"os"
	"strconv"
)

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonListResponse struct {
	Count    int                `json:"count"`
	Results  []NamedAPIResource `json:"results"`
}


func getRandomPokemon() (string, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon-species/?limit=1100")
	if err != nil {
		return "", fmt.Errorf("failed to fetch Pokémon list: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch Pokémon list: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var allMons PokemonListResponse
	if err := json.Unmarshal(body, &allMons); err != nil {
		return "", fmt.Errorf("failed to parse Pokémon list: %v", err)
	}

	if len(allMons.Results) == 0 {
		return "", fmt.Errorf("no Pokémon found in API response")
	}

	idx := rand.Intn(len(allMons.Results))
	return "https://pokeapi.co/api/v2/pokemon/" + allMons.Results[idx].Name, nil
}

func commandEncounter(cfg *config, c *pokecache.Cache, args []string) error {
	if len(cfg.Pokedex) == 0 {
		return fmt.Errorf("You need a Pokemon to begin an encounter.")
	}
	leadMonNum := -1
	for dex, name := range cfg.PokeKeys {
		if cfg.Party[name].CurHp > 0 {
			leadMonNum = dex
			break
		}
	}
	if leadMonNum == -1 {
		return fmt.Errorf("Please heal your Pokemon!")
	}
	var url, pokemonName string
	if len(args) > 0{
		pokemonName = args[0]
		url = "https://pokeapi.co/api/v2/pokemon/" + pokemonName 
	} else {
		var err error
		url, err = getRandomPokemon()
		if err != nil {
			return err
		}
		pokemonName = url[34:]
	}
	cachedBody, ok := c.Get(url)
    var body []byte
    if ok {
        body = cachedBody
    } else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Pokemon %s not found", pokemonName)
		}

		body, err = io.ReadAll(res.Body)
        if err != nil {
            return err
        }
        c.Add(url, body)
	}
	var pokemonData map[string]interface{}
	err := json.Unmarshal(body, &pokemonData)
	if err != nil {
		return err
	}
	id, ok := pokemonData["id"].(float64)
    if !ok {
        return fmt.Errorf("Could not parse Pokemon ID")
    }
    
    name, ok := pokemonData["name"].(string)
    if !ok {
        return fmt.Errorf("Could not parse Pokemon name")
    }
    
    baseExperience, ok := pokemonData["base_experience"].(float64)
    if !ok {
        return fmt.Errorf("Could not parse base experience")
    }
    
    height, ok := pokemonData["height"].(float64)
    if !ok {
        return fmt.Errorf("Could not parse height")
    }
    
    weight, ok := pokemonData["weight"].(float64)
    if !ok {
        return fmt.Errorf("Could not parse weight")
    }

	// generate a random ability based on their available ones. (Ignores hidden)
	abilities, ok := pokemonData["abilities"].([]interface{})
	if !ok {
        return fmt.Errorf("Could not parse ability list")
    }

	var ability string
	var toChoose int
	if len(abilities) == 1 {
		toChoose = 1
	} else {
		toChoose = 1 + rand.Intn(len(abilities)-1)
	}

	for _, rawAbility := range abilities {
		abilityEntry, ok := rawAbility.(map[string]interface{})
		if !ok {
			return fmt.Errorf("Invalid ability entry format")
		}
		slot := abilityEntry["slot"].(float64)
		if int(slot) == toChoose {
			toAdd, ok := abilityEntry["ability"].(map[string]interface{})
			if !ok {
				return fmt.Errorf("No ability found")
			}
			ability, ok = toAdd["name"].(string)
			if !ok {
				return fmt.Errorf("No name of ability found")
			}
			break
		}
	}

	speciesURL := "https://pokeapi.co/api/v2/pokemon-species/" + pokemonName
	res, err := http.Get(speciesURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Species info for %s not found", pokemonName)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var speciesData map[string]interface{}
	err = json.Unmarshal(body, &speciesData)
	if err != nil {
		return err
	}

	// Step 2: Get growth_rate name directly
	growthInfo, ok := speciesData["growth_rate"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("Could not find growth rate info")
	}

	growthRateName, ok := growthInfo["name"].(string)
	if !ok {
		return fmt.Errorf("Could not parse growth rate name")
	}

	// Find the moves the Pokemon should know at their current level. 
	lvl := 1 + rand.Intn(10)
	moves, ok := pokemonData["moves"].([]interface{})
	if !ok {
        return fmt.Errorf("Could not parse move list")
    }
	//fmt.Println(moves)
	moveSet := make(map[string]int)
	learnSet := make(map[string]Pokemove)
	for _, rawMove := range moves {
		moveEntry, ok := rawMove.(map[string]interface{})
		//fmt.Println(moveEntry)
		if !ok {
			return fmt.Errorf("can't parse move")
		}
		curMove := moveEntry["move"].(map[string]interface{})
		details := moveEntry["version_group_details"].([]interface{})
		group_to_use := details[len(details)-1].(map[string]interface{})
		//fmt.Println(curMove, group_to_use)
		method := group_to_use["move_learn_method"].(map[string]interface{})
		//fmt.Println(method)
		if method["name"].(string) != "level-up" {
			learnSet[curMove["name"].(string)] = Pokemove{
				LevelUp: 101,
				url: curMove["url"].(string),
			}
			continue
		} else {
			learnSet[curMove["name"].(string)] = Pokemove{
				LevelUp: int(group_to_use["level_learned_at"].(float64)),
				url: curMove["url"].(string),
			}
		}
		if int(group_to_use["level_learned_at"].(float64)) <= lvl {
			moveSet[curMove["name"].(string)] = int(group_to_use["level_learned_at"].(float64))
		}
	}
	

	for {
		if len(moveSet) <= 4 {
			break
		}
		min_val := 101
		var min_move string
		for moveName, val := range moveSet {
			if val <= min_val {
				min_val = val
				min_move = moveName
			}
		}
		delete(moveSet, min_move)
	}

	var thisMoveset []string
	for movename, _ := range moveSet {
		thisMoveset = append(thisMoveset, movename)
	}

	// Extract stats
	stats := make(map[string]int)
	statsArray, ok := pokemonData["stats"].([]interface{})
	if !ok {
		return fmt.Errorf("Could not parse stats")
	}
	for _, statItem := range statsArray {
		statMap, ok := statItem.(map[string]interface{})
		if !ok {
			return fmt.Errorf("Could not parse stat item")
		}
		
		statValue, ok := statMap["base_stat"].(float64)
		if !ok {
			return fmt.Errorf("Could not parse stat value")
		}
		
		statNameMap, ok := statMap["stat"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("Could not parse stat name")
		}
		
		statName, ok := statNameMap["name"].(string)
		if !ok {
			return fmt.Errorf("Could not parse stat name")
		}
		
		stats[statName] = int(statValue)
	}

	// Extract types
	var types []string
	typesArray, ok := pokemonData["types"].([]interface{})
	if !ok {
		return fmt.Errorf("Could not parse types")
	}
	for _, typeItem := range typesArray {
		typeMap, ok := typeItem.(map[string]interface{})
		if !ok {
			return fmt.Errorf("Could not parse type item")
		}
		
		typeNameMap, ok := typeMap["type"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("Could not parse type name")
		}
		
		typeName, ok := typeNameMap["name"].(string)
		if !ok {
			return fmt.Errorf("Could not parse type name")
		}
		
		types = append(types, typeName)
	}

	// Initialize moveData map
	moveData := make(map[string]Move)

	statVals := statCalculator(stats, lvl)
	curHp := statVals["hp"]
    
    // Create the Pokemon struct with the extracted data
    newPokemon := Pokemon{
        ID:             int(id),
        Name:           name,
		NickName:       name,
        BaseExperience: int(baseExperience),
        Height:         int(height),
        Weight:         int(weight),
		Stats: 			stats,
		Types: 			types,
		Level:          lvl,
		CurStats:       statVals,
		CurHp:          curHp,
		Moves:          thisMoveset,
		Movedata:       moveData,
		Ability:        ability,
		Learnset:       learnSet,
		ExpGroup:       growthRateName,
		TotalExp:       expChart[growthRateName][lvl],
    }
	for _, move := range thisMoveset {
		newPokemon, err = addMoveData(newPokemon, move)
		if err != nil {
			return fmt.Errorf("Could not get Pokemon Move Data")
		}
	}
	fmt.Printf("You found a level %v %v!\n", newPokemon.Level, pokemonName)
	leadMonName := cfg.PokeKeys[leadMonNum]
	fmt.Printf("Go! %v!\n", leadMonName)
	leadMon := cfg.Party[leadMonName]
	fmt.Print("What do you want to do? \n1: Catch\n2: run\n3: fight\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			userInput := scanner.Text()
			cleaned := cleanInput(userInput)
			if len(cleaned) == 0 {
				continue
			}
			command := cleaned[0]
			if command == "catch" || command == "1"{
				_, ok := cfg.Pokedex[pokemonName]
				if ok {
					return fmt.Errorf("You already have this Pokemon!")
				}
				fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
				catchRate := 500 - newPokemon.BaseExperience
				if catchRate < 10{
					catchRate = 10
				}
				caught := rand.Intn(100) < catchRate
				
				if caught {
					fmt.Printf("%v was caught!\n", pokemonName)
					secondScan := bufio.NewScanner(os.Stdin)
					fmt.Println("Would you like to nickname your Pokemon? Hit enter to say no.")
					for {
						if secondScan.Scan() {
							userInput := secondScan.Text()
							cleaned := cleanInput(userInput)
							if len(cleaned) == 0 {
								break
							}
							nickName := cleaned[0]
							newPokemon.NickName = nickName
							break
						}
					}
					cfg.Pokedex[pokemonName] = newPokemon
					// add moves after pokemon is added to the cfg
					if (len(cfg.Party) < 6) {
						cfg.Party[pokemonName] = newPokemon
						cfg.PokeKeys = append(cfg.PokeKeys, pokemonName)
					}
					break
				} else {
					fmt.Printf("%v escaped!\n", pokemonName)
				}
			}
			if command == "run" || command == "2"{
				fmt.Printf("You escaped.\n")
				break
			}
			// Begin battle logic against wild Pokemon
			if command == "fight" || command == "3" {
				scannerTwo := bufio.NewScanner(os.Stdin)
				backed := false
				for i, move := range leadMon.Moves {
					fmt.Printf("%v: %v     Power: %v, Accuracy: %v, Type: %v, Damage Type: %v\n", i+1, move, 
					leadMon.Movedata[move].Power, leadMon.Movedata[move].Accuracy, leadMon.Movedata[move].Poketype, leadMon.Movedata[move].Damagetype)
				}
				backNum := len(leadMon.Moves) + 1
				fmt.Printf("%v: Back\n", backNum)
				for {
					if scannerTwo.Scan() {
						backed = false
						userInput := scannerTwo.Text()
						cleaned := cleanInput(userInput)
						if len(cleaned) == 0 {
							continue
						}
						command := cleaned[0]
						if command == strconv.Itoa(backNum) || command == "back"{
							backed = true
							fmt.Print("What do you want to do? \n1: Catch\n2: run\n3: fight\n\n")
							break
						}
						moveUsed := false
						battleOver := false
						for i, move := range leadMon.Moves {
							if command == move || command == strconv.Itoa(i+1){
								moveUsed = true
								newPokemonDex := rand.Intn(len(newPokemon.Moves))
								newPokemonMove := newPokemon.Moves[newPokemonDex]
								cfg.Party[cfg.PokeKeys[leadMonNum]], newPokemon = battle(cfg.Party[cfg.PokeKeys[leadMonNum]], move, newPokemon, newPokemonMove)
								if newPokemon.CurHp == 0 {
									fmt.Println("You won!")
									cfg.Party[cfg.PokeKeys[leadMonNum]] = expGain(leadMon, newPokemon, false, cfg)
									battleOver = true
									break
								}
								availableToSwitch := []string{}
								if cfg.Party[cfg.PokeKeys[leadMonNum]].CurHp == 0 {
									canFight := false
									count := 1
									for _, partyMon := range cfg.PokeKeys {
										if cfg.Party[partyMon].CurHp != 0 {
											availableToSwitch = append(availableToSwitch, partyMon)
											if !canFight {
												canFight = true
												fmt.Println("Pick a Pokemon to switch into!")
											}
											fmt.Printf("%v: %v\n", strconv.Itoa(count), partyMon)
											count += 1
										}
									}
									if !canFight {
										battleOver = true
										fmt.Println("You blacked out!")
										break
									}
									if canFight {
										backed = true
										scannerThree := bufio.NewScanner(os.Stdin)
										for {
											if scannerThree.Scan(){
												userInput3 := scannerThree.Text()
												cleaned3 := cleanInput(userInput3)
												if len(cleaned3) == 0 {
													continue
												}
												command3 := cleaned3[0]
												validCommand := false
												for d, n := range availableToSwitch {
													if command3 == n || command3 == strconv.Itoa(d+1) {
														switchMons(cfg, leadMon, cfg.Party[n])
														leadMonName = n 
														leadMon = cfg.Party[n]
														validCommand = true
														for i, move := range leadMon.Moves {
															fmt.Printf("%v: %v     Power: %v, Accuracy: %v, Type: %v, Damage Type: %v\n", i+1, move, 
															leadMon.Movedata[move].Power, leadMon.Movedata[move].Accuracy, leadMon.Movedata[move].Poketype, leadMon.Movedata[move].Damagetype)
														}
														break
													}
												}
												if validCommand{
													break
												}
												fmt.Println("Please give a valid input")
											}
										}
										break
									}
								}
							}
						}
						if battleOver{
							break
						}
						if !moveUsed{
							println("Please give a valid input.")
						}
					}
				}
				if !backed {
					break
				}
			}
		}
	}
	cfg.Pokedex[leadMonName] = cfg.Party[leadMonName]
	return nil

}