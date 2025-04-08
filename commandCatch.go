package main

import (
	"net/http"       
    "encoding/json"
	"fmt"
	"io"
    "github.com/lucasolivo/Pokedex/internal/pokecache"
	"math/rand"
)

func commandCatch(cfg *config, c *pokecache.Cache, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please submit a Pokemon name")
	}
	pokemonName := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName 
	cachedBody, ok := c.Get(url)
    var body []byte
    if ok {
        fmt.Println("Using cached data")
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
    
    // Create the Pokemon struct with the extracted data
    newPokemon := Pokemon{
        ID:             int(id),
        Name:           name,
        BaseExperience: int(baseExperience),
        Height:         int(height),
        Weight:         int(weight),
    }
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	catchRate := 500 - newPokemon.BaseExperience
	if catchRate < 10{
		catchRate = 10
	}
	caught := rand.Intn(100) < catchRate
	
	if caught {
		fmt.Printf("%v was caught!\n", pokemonName)
		cfg.Pokedex[pokemonName] = newPokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil

}