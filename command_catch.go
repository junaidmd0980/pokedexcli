package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"math/rand"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemonName := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	var body []byte
	if val, ok := cfg.cache.Get(url); ok {
		body = val
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, body)
	}

	pokemon := Pokemon{}
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return err
	}

	fixedThreshold := 50
	roll := rand.Intn(pokemon.BaseExperience)
	if roll < fixedThreshold {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.pokedex[pokemonName] = pokemon

	return nil
}
