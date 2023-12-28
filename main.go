package main

import (
	"fmt"
	"net/http"
	"os"

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
)

func main() {
	port := getEnv("PORT")

	fmt.Println("running server on port", port)

	http.ListenAndServe(":"+port, nil)
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
