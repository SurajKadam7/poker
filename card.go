package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type CardsDeck struct {
	cards []int8
	size  int64
}

// mapping 1 --> diamonds, 2 --> spades, 3 --> hearts, 4 --> clubs

// randome card shuffler function
func (cd *CardsDeck) cardDistributer() (int8, int8) {
	maxSize := big.NewInt(cd.size)

	randCard, err := rand.Int(reader, maxSize)

	if err != nil {
		fmt.Println("problem in generating randome card")
	}

	randomeIndex := randCard.Int64()

	cardValue := cd.cards[randomeIndex]

	// swping the selected element to last and then reducing the size of the array will help to remove the distributed card
	cd.cards[randomeIndex], cd.cards[cd.size-1] = cd.cards[cd.size-1], cd.cards[randomeIndex]

	cd.size--

	// fmt.Printf("randome int : %d\nrandomeIndex : %d\ncards : %v\n\n", randomeIndex, cardValue, cd.cards)

	if cardValue > 0 && cardValue < 14 {
		return cardValue + 1, 1
	}
	if cardValue > 13 && cardValue < 27 {
		return cardValue - 12, 2
	}
	if cardValue > 26 && cardValue < 40 {
		return cardValue - 25, 3
	}

	return cardValue - 38, 4

}

// types of suits
// suits ---> spades, clubs, hearts, and diamonds

// types of the cards

// cards --> jack, queen, king, ace (11, 12, 13, 14)

// .......................................

// types of hands in poker

// royal flush

// stright flush --> flush + stright

// four of a kind

// full house --> three of a kind + two of kind

// flush

// stright

// three of a kind

// two pair

// pair

// high card

// tested the values
// sort.Ints(d["C"])
// sort.Ints(d["S"])
// sort.Ints(d["H"])
// sort.Ints(d["D"])

// fmt.Printf("c : %+v \n", d["C"])
// fmt.Printf("d : %+v \n", d["D"])
// fmt.Printf("h : %+v \n", d["H"])
// fmt.Printf("s : %+v \n", d["S"])
