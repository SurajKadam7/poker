package main

import (
	"crypto/rand"
)

// will write here the gamelogic of the game

var reader = rand.Reader

const cards = 52

func main() {
	t := Table{
		minPlayers:   2,
		maxPlayers:   6,
		minBuyin:     10,
		totalPlayers: 0,
		deler:        0,
		bbAmount:     25,
		sbAmount:     10,
	}
	// this will start adding the players on the table

	for i := 0; i < 10; i++ {
		t.addPlayers()
	}

	t.startNewGame()

}
