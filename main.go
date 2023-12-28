package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Pokemon struct {
	Name string `json:"name"`
	Type []PokemonType
}

type PokemonType string

const (
	Normal   = "normal"
	Fire     = "fire"
	Water    = "water"
	Electric = "electic"
	Grass    = "grass"
	Ice      = "ice"
	Fighting = "fighting"
	Poison   = "poison"
	Ground   = "ground"
	Flying   = "flying"
	Psychic  = "psychic"
	Bug      = "bug"
	Rock     = "rock"
	Ghost    = "ghost"
	Dragon   = "dragon"
	Dark     = "dark"
	Steel    = "steel"
	Fairy    = "fairy"
	Iron     = "iron"
)

var allPokemon = []Pokemon{}

func getAllPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(allPokemon)
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, p := range allPokemon {
		if p.Name == params["name"] {
			json.NewEncoder(w).Encode(p)
		}
	}
}

func deletePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, p := range allPokemon {
		if p.Name == params["name"] {
			allPokemon = append(allPokemon[:index], allPokemon[index+1:]...)
			fmt.Fprintln(w, "Found")
			break
		} else {
			w.Write([]byte("no such pokemon found"))
			return
		}
	}

	json.NewEncoder(w).Encode(allPokemon)
}

func main() {
	addPokemon("Treecko", []PokemonType{Grass})
	addPokemon("Lairon", []PokemonType{Steel, Iron})

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./static")))
	r.HandleFunc("/getPokemon", getAllPokemon).Methods("GET")
	r.HandleFunc("/getPokemon/{name}", getPokemon).Methods("GET")
	r.HandleFunc("/deletePokemon/{name}", deletePokemon).Methods("DELETE")

	port := getEnv("PORT")
	fmt.Println("running server on port", port)
	http.ListenAndServe(":"+port, r)
}

func getEnv(v string) string {
	err := godotenv.Load(".env")

	// Catch errors while opening the ENV file
	if err != nil {
		return fmt.Sprintln(err)
	} else {
		// Catch errors if the ENV variable is not found
		if _, ok := os.LookupEnv(v); !ok {
			return fmt.Sprintf("env variable %s not found", v)
		}
		return os.Getenv(v)
	}
}

func addPokemon(name string, types []PokemonType) {
	pokemon := Pokemon{Name: name, Type: types}
	allPokemon = append(allPokemon, pokemon)
}
