package main

import (
	"net/http"       
    "encoding/json"
	"fmt"
	"io"
    "github.com/lucasolivo/Pokedex/internal/pokecache"
)

func commandExplore(cfg *config, c *pokecache.Cache, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please submit a location name")
	}
	areaName := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + areaName
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
			return fmt.Errorf("area %s not found", areaName)
		}

		body, err = io.ReadAll(res.Body)
        if err != nil {
            return err
        }
        c.Add(url, body)
	}
	var locRes PokemonAreaResponse
    err := json.Unmarshal(body, &locRes)
    if err != nil {
        return err
    }
	fmt.Printf("Exploring %s...\n", areaName)
    fmt.Println("Found Pokemon:")
    
    for _, encounter := range locRes.PokemonEncounters {
        fmt.Printf(" - %s\n", encounter.Pokemon.Name)
    }
	return nil
}