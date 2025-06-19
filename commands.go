package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func prepareCommands() func() map[string]cliCommand {
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
		}
		return comms
	}
}
