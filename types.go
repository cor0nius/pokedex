package main

import pc "github.com/cor0nius/pokedexcli/internal"

type cliCommand struct {
	name        string
	description string
	callback    func(*pc.Cache) error
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
