# Pokedex CLI

A command-line Pokedex built in Go, powered by the [PokeAPI](https://pokeapi.co/). Explore locations, catch Pokémon, and inspect your collection—all directly from your terminal.

## 🚀 Features

- **World Map Exploration:** Navigate the Pokémon world map area-by-area.
- **Location Scanning:** Explore specific locations to discover which Pokémon inhabit them.
- **Catch Mechanics:** Attempt to catch wild Pokémon using a probability engine scaled by their base experience.
- **Stat Inspection:** Examine detailed statistics, types, and properties of any Pokémon you catch.
- **Pokedex Ledger:** View a structured inventory of your full collection.
- **Internal Caching:** Built-in caching layer with timed expirations to prevent excessive API throttling.

## 📋 Requirements

- **Go 1.20** or newer

## 🛠️ Installation & Setup

1. **Clone the repository:**
   ```bash
   git clone <your-repo-url>
   cd pokedexcli
   ```

2. **Build the binary:**
   ```bash
   go build -o pokedex
   ```

## 🎮 Usage

Launch the compiled binary to start the interactive REPL loop:
```bash
./pokedex
```

You will be greeted by the Pokedex command prompt:
```text
Pokedex > 
```

## ⌨️ Available Commands

| Command | Description |
| :--- | :--- |
| `help` | Displays a helpful usage guide and lists commands |
| `map` | Displays the next 20 location areas on the map |
| `mapb` | Displays the previous 20 location areas on the map |
| `explore <location-area>` | Lists all Pokémon species native to the given area |
| `catch <pokemon-name>` | Attempts to catch a named Pokémon (chance based on difficulty) |
| `inspect <pokemon-name>` | Displays stats, height, weight, and types of a caught Pokémon |
| `pokedex` | Lists the names of all Pokémon currently in your collection |
| `exit` | Closes the Pokedex interface and exits the program |

## 🕹️ Example Session

```text
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
```

## 📁 Project Structure

```text
pokedexcli/
├── main.go               # Application entry point
├── repl.go               # Core REPL loop, text parsing, and command routing
├── command_*.go          # Individual command handlers (map, catch, inspect, etc.)
├── internal/
│   └── pokecache/        # Custom in-memory cache mechanism with automatic janitor reaping
├── go.mod                # Go module configurations
└── README.md             # Project documentation
```

## 🧰 Built With

- **Go** (Golang)
- [PokeAPI](https://pokeapi.co/) — The RESTful Pokémon API
