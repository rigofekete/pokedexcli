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
		// fmt.Printf("you have not caught the %s pokemon", name)
		return fmt.Errorf("you have not caught the %s pokemon", name)
	}

	// hp := 	           pokemon.Stats[0].BaseStat
	// attack :=  	   pokemon.Stats[1].BaseStat
	// defense :=     	   pokemon.Stats[2].BaseStat
	// specialAttack :=   pokemon.Stats[3].BaseStat
	// specialDefense :=  pokemon.Stats[4].BaseStat
	// speed := 	   pokemon.Stats[5].BaseStat
	//
	//
	// fmt.Printf("Name: %s\nHeight: %v\nWeight: %v\nStats: \n -hp: %v\n -attack: %v\n -defense: %v\n -special-attack: %v\n -special-defense: %v\n -speed: %v\n",pokemon.Name, pokemon.Height, pokemon.Weight, hp, attack, defense, specialAttack, specialDefense, speed)
	//
	//

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, statInfo := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", statInfo.Stat.Name, statInfo.BaseStat)
	}

	fmt.Println("Types: ")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf(" -%s", typeInfo.Type.Name)
	}

	return nil
}


