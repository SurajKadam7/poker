package main

import (
	"crypto/rand"
)

// will write here the gamelogic of the game

var reader = rand.Reader

const cards = 52

func main() {
	var i int8
	allCards := make([]int8, cards)
	for i = 1; i <= cards; i++ {
		allCards[i-1] = i
	}

	cd := CardsDeck{
		cards: allCards,
		size:  cards,
	}
	// need to start game on a table
	card, suite := cd.cardDistributer()

}
