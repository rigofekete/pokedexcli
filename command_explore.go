package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, area *string) error {
	if area == nil || *area == "" {
		return errors.New("\nexplore command needs an area parameter (explore <area_name>)\n") 
	}

	fmt.Printf("\nExploring %s...\n", *area)
	areaList, err := cfg.pokeapiClient.ListAreas(*area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, s := range areaList.PokemonEncounters {
		fmt.Printf(" - %s\n", s.Pokemon.Name)
	}

	return nil


}
