package main

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

type LocationDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area name")
	}

	areaName := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	fmt.Printf("Exploring %s...\n", areaName)

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

	locDetail := LocationDetail{}
	err := json.Unmarshal(body, &locDetail)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locDetail.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
