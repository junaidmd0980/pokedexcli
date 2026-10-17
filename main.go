package main

import (
	"time"
	"github.com/junaidmd0980/pokedexcli/internal/pokecache"
)

func main() {
	cfg := &config{
		commands: getCommands(),
		cache: pokecache.NewCache(5 * time.Second),
		pokedex: make(map[string]Pokemon),
	}
	startRepl(cfg)
}
