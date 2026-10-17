package main

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config, _ ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextLocURL != nil {
		url = *cfg.nextLocURL
	}

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


	locAreas := LocationAreas{}
	err := json.Unmarshal(body, &locAreas)
	if err != nil {
		return err
	}

	cfg.nextLocURL = locAreas.Next
	cfg.previousLocURL = locAreas.Previous

	for _, loc := range locAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, _ ...string) error {
	if cfg.previousLocURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.previousLocURL
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

	locAreas := LocationAreas{}
	err := json.Unmarshal(body, &locAreas)
	if err != nil {
		return err
	}

	cfg.nextLocURL = locAreas.Next
	cfg.previousLocURL = locAreas.Previous

	for _, loc := range locAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
