package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)


type LocationPage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return fmt.Errorf("Error getting pokeapi data: %w ", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 {
		return fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}
	

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var page LocationPage

	if err := json.Unmarshal(data, &page); err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}

	for i := range page.Results {
		fmt.Println(page.Results[i].Name)
	}

	// fmt.Println(page)
		
	return nil
}
