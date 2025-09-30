package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokeapiClient.Pokedex) == 0 {
		return errors.New("Pokedex is empty, go and catch some Pokemon first!")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokeapiClient.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
