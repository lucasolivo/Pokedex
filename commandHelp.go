package main

import "fmt"

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
	fmt.Println()
    
    // Dynamically generate the help text
    for _, cmd := range commands {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    
    return nil
}