package main

// Public API - https://pokeapi.co/docs/v2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map the Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

// Querying the API endpoint - http://pokeapi.co/api/v2/pokedex/kanto/

func main() {
	// GET Request
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(string(responseData))

	// Marshalling means converting data structures
	// into a format suitable for storage or transmission (like a byte stream).

	// Unmarshalling is the reverse process
	// of converting a format back into the original data structure.

	// Unmarshalling the returned JSON string into a new variable

	var responseObject Response

	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	// Listing all the pokemons

	for _, pokemon := range responseObject.Pokemon {
		fmt.Printf("%v \n", pokemon.Species.Name)
	}
}
