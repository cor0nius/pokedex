package main

import pc "github.com/cor0nius/pokedexcli/internal"

type cliCommand struct {
	name, description string
	callback          func(*config, *pc.Cache) error
}

type locationAreaAPI struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type areaEncounters struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name    string        `json:"name"`
	Url     string        `json:"url"`
	BaseExp int           `json:"base_experience"`
	Height  int           `json:"height"`
	Weight  int           `json:"weight"`
	Types   []PokemonType `json:"types"`
}

type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type config struct {
	next, previous string
	aux, name      string
	pokedex        Pokedex
}

type Pokedex map[string]Pokemon
