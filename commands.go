package main

func getCommands() func() map[string]cliCommand {
	return func() map[string]cliCommand {
		comms := map[string]cliCommand{
			"help": {
				name:        "help",
				description: "Display a help message",
				callback:    commandHelp,
			},
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"map": {
				name:        "map",
				description: "Display a list of next 20 locations",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Display a list of previous 20 locations",
				callback:    commandMapb,
			},
			"explore": {
				name:        "explore",
				description: "List Pokemon available at chosen location",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Throw a Pokeball for a chance to catch a chosen Pokemon",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect a Pokemon in your Pokedex",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "List all caught pokemon",
				callback:    commandPokedex,
			},
		}
		return comms
	}
}
