package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	pc "github.com/cor0nius/pokedexcli/internal"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func commandExit(cfg *config, cache *pc.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, cache *pc.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	comms := getCommands()
	for comm := range comms() {
		fmt.Printf("%s: %s\n", comms()[comm].name, comms()[comm].description)
	}
	return nil
}

func commandMap(cfg *config, cache *pc.Cache) error {
	var locationAreas locationAreaAPI
	var data []byte
	url := cfg.next
	_, ok := cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	} else {
		data, _ = cache.Get(url)
	}
	err := json.Unmarshal(data, &locationAreas)
	if err != nil {
		return err
	}
	cfg.next = locationAreas.Next
	cfg.previous = locationAreas.Previous
	for i := range locationAreas.Results {
		fmt.Println(locationAreas.Results[i].Name)
	}
	return nil
}

func commandMapb(cfg *config, cache *pc.Cache) error {
	if url := cfg.previous; url != "" {
		var locationAreas locationAreaAPI
		var data []byte
		_, ok := cache.Get(url)
		if !ok {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			data, err = io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			cache.Add(url, data)
		} else {
			data, _ = cache.Get(url)
		}
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return err
		}
		cfg.next = locationAreas.Next
		cfg.previous = locationAreas.Previous
		for i := range locationAreas.Results {
			fmt.Println(locationAreas.Results[i].Name)
		}
		return nil
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
}

func commandExplore(cfg *config, cache *pc.Cache) error {
	url := cfg.area
	var encounters areaEncounters
	var data []byte
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &encounters)
	if err != nil {
		return err
	}
	for encounter := range encounters.PokemonEncounters {
		fmt.Println(encounters.PokemonEncounters[encounter].Pokemon.Name)
	}
	return nil
}
