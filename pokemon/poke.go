//http://pokeapi.co/api/v2/pokedex/kanto/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
	Name    string    `json: "name`
	Pokemon []Pokemon `json: "pokemon_entries"`
}

type Pokemon struct {
	Numero  int            `json: "entry_number`
	Especie PokemonSpecies `json: "pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json: name`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(responseData))

	var responseObj Response
	json.Unmarshal(responseData, &responseObj)

	//fmt.Println(responseObj.Name)
	//fmt.Println(responseObj.Pokemon)

	for _, pokemon := range responseObj.Pokemon {
		fmt.Println(pokemon.Especie.Name)
	}

}
