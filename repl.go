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

func commandExit(cache *pc.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cache *pc.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	comms := getCommands()
	for comm := range comms() {
		fmt.Printf("%s: %s\n", comms()[comm].name, comms()[comm].description)
	}
	return nil
}

func commandMap(cache *pc.Cache) error {
	var locationAreas locationAreaAPI
	url := nextPage(true)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return err
	}
	for i := range locationAreas.Results {
		fmt.Println(locationAreas.Results[i].Name)
	}
	return nil
}

func commandMapb(cache *pc.Cache) error {
	if mapPage > 1 {
		var locationAreas locationAreaAPI
		url := nextPage(false)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &locationAreas)
		if err != nil {
			return err
		}
		for i := range locationAreas.Results {
			fmt.Println(locationAreas.Results[i].Name)
		}
		return nil
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
}

func getMapPage(url string) func(bool) string {
	return func(fwd bool) string {
		newUrl := ""
		if fwd {
			mapPage++
			newUrl = fmt.Sprintf("%s%d", url, (mapPage-1)*20)
		} else {
			mapPage--
			newUrl = fmt.Sprintf("%s%d", url, (mapPage-1)*20)
		}
		return newUrl
	}
}

var mapPage int
var nextPage = getMapPage("https://pokeapi.co/api/v2/location-area/?offset=")
