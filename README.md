# Pokedex REPL

This project is a command-line Pokedex application built in Go. It functions as a REPL (Read-Eval-Print Loop), allowing users to interactively explore the Pokemon world through the [PokéAPI](https://pokeapi.co/).

This was a guided project completed as part of the ["Build a Pokedex in Go"](https://www.boot.dev/courses/build-pokedex-cli-golang) course on boot.dev.

## Features

*   Explore different areas in the Pokemon world.
*   Discover which Pokemon can be found in each area.
*   Catch Pokemon to add them to your personal Pokedex.
*   Inspect the details of Pokemon you have caught.
*   View a list of all caught Pokemon.

## Key Concepts & Skills Learned

*   **REPL (Read-Eval-Print Loop):** Built an interactive command-line interface to process user commands.
*   **HTTP Networking:** Made GET requests to the external PokéAPI to fetch data.
*   **Data Serialization:** Handled JSON data from the API, unmarshalling it into Go structs.
*   **Concurrency:** Implemented a thread-safe cache using goroutines and mutexes to store API responses, reducing redundant network requests and improving performance.
*   **State Management:** Managed application state, including pagination for locations and the collection of caught Pokemon.
