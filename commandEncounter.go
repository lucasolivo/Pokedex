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
    
    // Create the Pokemon struct with the extracted data
    newPokemon := Pokemon{
        ID:             int(id),
        Name:           name,
        BaseExperience: int(baseExperience),
        Height:         int(height),
        Weight:         int(weight),
		Stats: 			stats,
		Types: 			types,
		Level:          1 + rand.Intn(10),
    }
	fmt.Printf("You found a level %v %v!\n", newPokemon.Level, pokemonName)
	fmt.Print("What do you want to do? Catch or run?\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			userInput := scanner.Text()
			cleaned := cleanInput(userInput)
			if len(cleaned) == 0 {
				continue
			}
			command := cleaned[0]
			if command == "catch" {
				fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
				catchRate := 500 - newPokemon.BaseExperience
				if catchRate < 10{
					catchRate = 10
				}
				caught := rand.Intn(100) < catchRate
				
				if caught {
					fmt.Printf("%v was caught!\n", pokemonName)
					cfg.Pokedex[pokemonName] = newPokemon
					if (len(cfg.Party) < 6) {
						cfg.Party[pokemonName] = newPokemon
						cfg.PokeKeys = append(cfg.PokeKeys, pokemonName)
					}
					break
				} else {
					fmt.Printf("%v escaped!\n", pokemonName)
				}
			}
			if command == "run" {
				fmt.Printf("You escaped.\n")
				break
			}
		}
	}

	return nil

}