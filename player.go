package main

import (
	"fmt"
	"sort"
)

// if four of a kind the it will not become straight

type custome struct {
	ok    bool
	score int8
}

type Hand struct {
	Straight      custome
	Flush         custome
	Pair          custome
	ThreeOfAKind  custome
	straightArray []int8
}

type Player struct {
	cards    [][]int8
	amount   int64
	hand     Hand
	isBB     bool
	isSB     bool
	isDeler  bool
	isAllIn  bool
	isFold   bool
	canCheck bool
	isRaised bool
}

func (p *Player) Raise() {

}

func (p *Player) Check() {

}

func (p *Player) Fold() {
	p.isFold = true

}

func (p *Player) Call() {

}

func (p *Player) AllIn() {
	p.isAllIn = true
}

// taking input from the players ...
func (p *Player) Play() {
	var move int

	if p.canCheck {
		fmt.Printf("Options : \n  allIn : 1 \n  raise : 2 \n  call : 3 \n  check : 4 \n  fold : 5 \n\nEnter your move : ")
	} else {
		fmt.Printf("Options : \n  allIn : 1 \n  raise : 2 \n  call : 3 \n  fold : 4 \n\nEnter your move : ")
	}

	switch move {
	case 1:
		fmt.Println("allIn")

		p.Raise()
	case 2:
		fmt.Println("raise")
		p.Raise()
	case 3:
		fmt.Println("call")
		p.Call()
	case 4:
		fmt.Println("check")
		p.Check()
	default:
		fmt.Println("fold")
		p.Fold()
	}

}

func (p *Player) sortCards() {
	sort.Slice(p.cards, func(i, j int) bool {
		return p.cards[i][0] < p.cards[j][0]
	})
}

func (p *Player) sortSuits() {
	sort.Slice(p.cards, func(i, j int) bool {
		return p.cards[0][i] < p.cards[0][j]
	})
}

// need to impliment this
func (p *Player) maxHand() (score int8) {
	return
}

// for royal flush condition is straight + flush and
// both should be for the higher cards values

func (p *Player) isRoyalFlush() (ok bool) {
	ok, score := p.isStraight()

	if !ok || score != 14 {
		return false
	}

	ok, score = p.isFlush()

	if ok && score == 14 {
		return true
	}

	return false

}

func (p *Player) isFlush() (ok bool, score int8) {
	p.sortSuits()
	for i := 0; i < 3; i++ {
		if p.cards[i][0] == p.cards[i+4][0] {
			ok = true
			score = p.cards[i][0]
		}
	}
	return ok, score
}

func (p *Player) isStraight() (ok bool, score int8) {
	p.sortCards()
	for i := 0; i < 3; i++ {
		if (p.cards[i][0] + 4) == p.cards[i+4][0] {
			ok = true
			score = p.cards[i+4][0]
		}
	}
	return ok, score
}

// 5 out of 7
func (p *Player) isFullHouse() (ok bool, score int8) {
	// need to add weightage according to the three of a kind and one pair
	// three of a kind should get more weightage
	if p.hand.Pair.ok && p.hand.ThreeOfAKind.ok {
		ok = true
		score = p.hand.Pair.score + (p.hand.ThreeOfAKind.score * 5)
	}

	return ok, score
}

func (p *Player) isStraightFlush() (ok bool, score int8) {
	if p.hand.Flush.ok {

		s := p.hand.Flush.score

		for sc := range p.hand.straightArray {
			if sc == int(s) {
				ok = true
				score = s
			}
		}
	}

	return ok, score
}

func (p *Player) isKind(rng int16, kind int16) (ok bool, score int8) {
	// need sorted card array
	var i int16
	for i = 0; i <= rng; i++ {
		if p.cards[i][0] == p.cards[i+kind][0] {
			ok = true
			score = p.cards[i][0]
			i += 2
		}
	}

	return ok, score
}

func (p *Player) isFourOfAKind() (ok bool, score int8) {
	// need sorted card array
	ok, score = p.isKind(3, 3)
	return ok, score
}

func (p *Player) isThreeOfAKind() (ok bool, score int8) {
	// need sorted card array
	ok, score = p.isKind(4, 2)
	return ok, score

}

func (p *Player) isTwoOfAKind() (ok bool, score int8) {
	// need sorted card array
	ok, score = p.isKind(5, 1)
	return ok, score
}

func (p *Player) HighCard() (ok bool, score int8) {
	// need sorted card array
	ok = true
	score = p.cards[6][0]
	return ok, score
}
