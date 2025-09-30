package main

import (
	"fmt"
	"errors"
	"math/rand"
)


func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("\ncatch command needs a pokemon name (catch <pokemon_name>)\n") 
	}

	name := args[0]

	fmt.Printf("\nThrowing a Pokeball at %s...\n", name)

	pokemonData, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}


	// Debug logs 
	// fmt.Printf("PokemonData 1: %v", pokemonData)
	// fmt.Printf("pokemon base experience: %v\n", pokemonData.BaseXP)
	// fmt.Printf("pokemon name: %s\n", pokemonData.Name)
	// fmt.Printf("pokemon data: %v\n", pokemonData)

	xp := pokemonData.BaseXP
	var chance int
	switch {
		case xp >= 0 && xp <= 49:
			chance = 2
		case xp >= 50 && xp <= 99:
			chance = 5
		case xp >= 100:
			chance = 10
	}

	chanceRes := rand.Intn(chance) + 1

	// Debug log 
	// fmt.Printf("chance result: %v\n", chanceRes)

	if chanceRes == chance {
		fmt.Printf("%s was caught\n", name)
		fmt.Println("You may now inspect it with the inspect command.")

		cfg.pokeapiClient.Pokedex[name] = pokemonData
		// Debug logs 
		// fmt.Printf("PokemonData 1: %v", pokemonData)
		// fmt.Printf("Pokedex list: %v", cfg.pokeapiClient.Pokedex)
		return nil
	}

	fmt.Printf("%s escaped!\n", name)
	return nil
}


