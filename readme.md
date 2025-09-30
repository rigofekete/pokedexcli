# Pokédex CLI (PokeAPI REPL)

A command-line Pokédex written in Go with an interactive REPL powered by the PokeAPI.  
Explore regions, encounter and catch Pokémon, view details, and enjoy fast responses via local caching.

---

## Features

- Interactive REPL with parsed commands  
- Explore areas and encounter wild Pokémon  
- Catch and list captured Pokémon  
- View Pokémon details (types, stats, abilities)  
- Caching layer for faster, offline-resilient queries  
- Modular packages for API, cache, and domain logic  
- Easy to extend (battles, parties, evolution timers, persistence)  

---

## Commands (examples)

- `help` — list commands  
- `explore <area>` — discover Pokémon in an area  
- `map` — Get the next page of locations  
- `mapb` — Get the previous page of locations  
- `catch <name>` — attempt to catch a Pokémon  
- `inspect <name>` — view details  
- `pokedex` — show your caught Pokémon  
- `exit` — quit the REPL  


## Acknowledgments

Data provided by the [PokeAPI](https://pokeapi.co).

