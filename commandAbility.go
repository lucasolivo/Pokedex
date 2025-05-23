package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

type AbilityEffectEntry struct {
	ShortEffect string `json:"short_effect"`
	Language    struct {
		Name string `json:"name"`
	} `json:"language"`
}

type AbilityResponse struct {
	Name          string               `json:"name"`
	EffectEntries []AbilityEffectEntry `json:"effect_entries"`
}

type AbilityListResponse struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandAbility(cfg *config, args []string) error {
	num := 1
	if len(args) > 0 { //default value to 1
		var err error
		num, err = strconv.Atoi(args[0])
		if err != nil || num < 1 || num > 10 { //if not a number of less than 1 or more than 10 (the range we allow)
			return fmt.Errorf("Please enter a number between 1 and 10 after ability")
		}
	}

	// Step 1: Fetch all valid abilities
	res, err := http.Get("https://pokeapi.co/api/v2/ability/?limit=1000")
	if err != nil {
		return fmt.Errorf("failed to fetch ability list: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch ability list: status %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var allAbilities AbilityListResponse
	if err := json.Unmarshal(body, &allAbilities); err != nil {
		return fmt.Errorf("failed to parse ability list: %v", err)
	}

	if len(allAbilities.Results) == 0 {
		return fmt.Errorf("no abilities found in API response")
	}

	// Step 2: Pick random unique abilities
	used := make(map[int]bool)
	for i := 0; i < num; i++ {
		var idx int
		for {
			idx = rand.Intn(len(allAbilities.Results))
			if !used[idx] {
				used[idx] = true
				break
			}
		}

		abilityURL := allAbilities.Results[idx].URL
		res, err := http.Get(abilityURL)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Printf("Warning: failed to fetch ability at %s\n", abilityURL)
			continue
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var ability AbilityResponse
		if err := json.Unmarshal(body, &ability); err != nil {
			return err
		}

		shortEffect := "No English short effect found."
		for _, entry := range ability.EffectEntries {
			if entry.Language.Name == "en" {
				shortEffect = entry.ShortEffect
				break
			}
		}

		fmt.Printf("Ability: %s\nEffect: %s\n\n", ability.Name, shortEffect)
	}

	return nil
}