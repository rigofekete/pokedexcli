package pokeapi

type RespShallowLocations struct {
	Count int 		`json:"count"`
	Next *string 		`json:"next"`
	Previous *string 	`json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string  `json:"url"`
	}`json:"results"`
}


type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}


type Pokemon struct {
	Name 	string 		`json:"name"`
	BaseXP 	int 		`json:"base_experience"`
	Height 	int 		`json:"height"`
	Weight 	int 		`json:"weight"`
	Stats 	[]PokemonStat 	`json:"stats"`
	Types 	[]PokemonType 	`json:"types"`
}


type PokemonStat struct {
	BaseStat int         `json:"base_stat"`
	Effort   int   	     `json:"effort"`
	Stat     NamedAPI    `json:"stat"`
}

type PokemonType struct {
	Type NamedAPI  `json:"type"`
}

type NamedAPI struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
