package main

import (
	"net/http"       
    "encoding/json"
	"fmt"
	"io"
    "github.com/lucasolivo/Pokedex/internal/pokecache"
)

func commandMapb(cfg *config, c *pokecache.Cache) error {

    url := "https://pokeapi.co/api/v2/location-area"
    if cfg.prevLocationsURL != nil {
        url = *cfg.prevLocationsURL
    }

    // Step 1: Check the cache
    cachedBody, ok := c.Get(url)
    var body []byte
    if ok {
        fmt.Println("Using cached data")
        body = cachedBody
    } else {
        // Step 2: Make the HTTP request if not in cache
        res, err := http.Get(url)
        if err != nil {
            return err
        }
        defer res.Body.Close()

        body, err = io.ReadAll(res.Body)
        if err != nil {
            return err
        }

        // Step 3: Add the response body to the cache
        c.Add(url, body)
    }

    // Step 4: Parse the response
    var locRes LocationAreaResponse
    err := json.Unmarshal(body, &locRes)
    if err != nil {
        return err
    }

    // Print location names
    for _, loc := range locRes.Results {
        fmt.Println(loc.Name)
    }

    // Update your config with the next/previous URLs
    cfg.nextLocationsURL = locRes.Next
    cfg.prevLocationsURL = locRes.Previous

    return nil

}