package main

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

}
