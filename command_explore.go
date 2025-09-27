package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("\nexplore command needs an area parameter (explore <area_name>)\n") 
	}

	name := args[0]
	fmt.Printf("\nExploring %s...\n", name)
	areaList, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range areaList.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
