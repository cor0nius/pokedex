package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	comms := getCommands()
	for comm := range comms() {
		fmt.Printf("%s: %s\n", comms()[comm].name, comms()[comm].description)
	}
	return nil
}

func commandMap() error {
	url := nextPage(true)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var locationAreas []string
	decoder := json.NewDecoder(resp.Body)
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if token == "name" {
			token, _ := decoder.Token()
			locationAreas = append(locationAreas, token.(string))
		}
	}
	for i := range locationAreas {
		fmt.Println(locationAreas[i])
	}
	return nil
}

func commandMapb() error {
	if mapPage > 1 {
		url := nextPage(false)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var locationAreas []string
		decoder := json.NewDecoder(resp.Body)
		for {
			token, err := decoder.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			if token == "name" {
				token, _ := decoder.Token()
				locationAreas = append(locationAreas, token.(string))
			}
		}
		for i := range locationAreas {
			fmt.Println(locationAreas[i])
		}
		return nil
	} else {
		fmt.Println("you're on the first page")
	}
	return nil
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
