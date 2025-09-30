package main

import (
	"fmt"
	"errors"
)


func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("\ninspect command needs a pokemon name (catch <pokemon_name>)\n") 
	}

	name := args[0]

	pokemon, exists := cfg.pokeapiClient.Pokedex[name]
	if !exists {
		return fmt.Errorf("you have not caught the %s pokemon", name)
	}


	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, statInfo := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", statInfo.Stat.Name, statInfo.BaseStat)
	}

	fmt.Println("Types: ")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf(" - %s\n", typeInfo.Type.Name)
	}

	return nil
}


