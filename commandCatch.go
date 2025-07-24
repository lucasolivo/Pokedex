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
)

func commandCatch(cfg *config, c *pokecache.Cache, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please submit a Pokemon name")
	}
	pokemonName := args[0]
	_, ok := cfg.Pokedex[pokemonName]
	if ok {
		return fmt.Errorf("You already have this Pokemon!")
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName 
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
    }

	for _, move := range thisMoveset {
		newPokemon, err = addMoveData(newPokemon, move)
		if err != nil {
			return fmt.Errorf("Could not get Pokemon Move Data")
		}
	}
	
	fmt.Printf("You found a level %v %v!\n", newPokemon.Level, pokemonName)
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
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil

}