package main

import (
	"encoding/json"
	"os"
)

func saveGame(cfg *config) error {
	file, err := os.Create("savefile.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ") // Pretty print
    return encoder.Encode(cfg)
}

func loadGame() (*config, error) {
	var cfg config

    file, err := os.Open("savefile.json")
    if err != nil {
        return nil, err // You can check for os.IsNotExist(err) if you want to default to empty state
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
    return &cfg, err
}