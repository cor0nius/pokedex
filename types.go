package main

import pc "github.com/cor0nius/pokedexcli/internal"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pc.Cache) error
}

type locationAreaAPI struct {
	Count    int
	Next     string
	Previous string
	Results  []locationArea
}

type locationArea struct {
	Name string
	Url  string
}

type areaEncounters struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon"`
}

type config struct {
	next     string
	previous string
	area     string
}
