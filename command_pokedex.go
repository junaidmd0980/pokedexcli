package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {
	pokedex := cfg.pokedex
	if len(pokedex) == 0 {
		fmt.Println(" (empty)")
		return nil
	}

	for _, pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
