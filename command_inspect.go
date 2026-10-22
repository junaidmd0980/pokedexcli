package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, exists := cfg.pokedex[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
