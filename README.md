# Pokedex CLI

A command-line Pokedex built in Go, powered by the [PokeAPI](https://pokeapi.co/). Explore locations, catch Pokemon, and inspect your collection, all from the terminal.

## Features

- Explore the Pokemon world map, one page at a time
- Explore individual locations to see which Pokemon live there
- Attempt to catch Pokemon (success is probability-based on their base experience)
- Inspect Pokemon you've caught to see their stats and types
- View your Pokedex of all caught Pokemon
- Built-in caching layer to avoid hammering the API with repeat requests

## Installation

Clone the repo and build the binary:

\`\`\`bash
git clone <your-repo-url>
cd pokedexcli
go build -o pokedex
\`\`\`

## Usage

Run the compiled binary to start the REPL:

\`\`\`bash
./pokedex
\`\`\`

You'll be dropped into an interactive prompt:

\`\`\`
Pokedex >
\`\`\`

## Commands

| Command | Description |
|---|---|
| `help` | Displays a help message |
| `map` | Displays the next 20 location areas |
| `mapb` | Displays the previous 20 location areas |
| `explore <location-area>` | Lists the Pokemon found in a given location area |
| `catch <pokemon-name>` | Attempts to catch the named Pokemon |
| `inspect <pokemon-name>` | Shows details about a caught Pokemon |
| `pokedex` | Lists the names of all Pokemon you've caught |
| `exit` | Exits the Pokedex |

## Example Session

\`\`\`
Pokedex > catch pidgey
Throwing a Pokeball at pidgey...
pidgey was caught!
You may now inspect it with the inspect command.
Pokedex > catch caterpie
Throwing a Pokeball at caterpie...
caterpie was caught!
You may now inspect it with the inspect command.
Pokedex > pokedex
Your Pokedex:
 - pidgey
 - caterpie
\`\`\`

## Project Structure

- `main.go` / `repl.go` — REPL loop and command dispatch
- `command_*.go` — individual command implementations (map, explore, catch, inspect, pokedex, help, exit)
- `internal/pokecache` — simple in-memory cache with expiration

## Built With

- Go
- [PokeAPI](https://pokeapi.co/)
