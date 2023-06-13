package main

type Game struct {
	players      []Player
	totalPlayers int8

	sb          int8
	bb          int8
	firstPlayer int8

	round         int8
	pot           int64
	currentPlayer int8

	bbAmout  int64
	sbAmount int64

	cd           *CardsDeck
	raisedPlayer int8
	raiseAmount  int64
	equalPot     int64

	nextToDeler int8 // index of the player from the players array
	stage       int8 // 0-preflop, 1-postflop, 2-turn, 3-river

	communityCards [][]int8

	winners  [][]int8
	totalPot int64

	playersInGame int8
}

func (g *Game) deltCards() {
	var i int8
	for i = 0; i < g.totalPlayers; i++ {
		// first card
		card, suit := g.cd.cardDistributer()
		g.players[i].cards = append(g.players[i].cards, []int8{card, suit})

		// second card
		card, suit = g.cd.cardDistributer()
		g.players[i].cards = append(g.players[i].cards, []int8{card, suit})

	}

}

func (g *Game) newCardDeck() *CardsDeck {
	var i int8
	allCards := make([]int8, cards)
	for i = 1; i <= cards; i++ {
		allCards[i-1] = i
	}

	cd := CardsDeck{
		cards: allCards,
		size:  cards,
	}

	return &cd
}

func (g *Game) startGame() {
	// create new card deck
	g.cd = g.newCardDeck()

	// cutting amount from bb and sb
	g.players[g.bb].amount -= g.bbAmout
	g.players[g.sb].amount -= g.sbAmount

	// first player is the player next to the deler all the rounds will start form him
	g.firstPlayer = g.sb

	// bb will raise the amound first time in the game
	g.raisedPlayer = (g.bb + 1) % g.totalPlayers

	// card distribution will start here
	g.deltCards()

	// this will handle all the four rounds in the poker
	g.HandleRounds()

	// finding the winner of the current game ..
	g.findWinners()

	// distribute the winnings to every player
	g.distributeWinnings()

}

func (g *Game) distributeWinnings() {
	totalWinners := len(g.winners)

	winAmount := g.totalPot / int64(totalWinners)

	for _, winner := range g.winners {
		w, _ := winner[0], winner[1]
		g.players[w].amount += winAmount
	}
}

// I can pass the winners by referance also need to think about memeory optimizations
func (g *Game) findWinners() {
	var winners [][]int8

	for i, player := range g.players {
		score := player.maxHand()

		if score > winners[0][1] {
			winners = [][]int8{}
			winners = append(winners, []int8{int8(i), score})
		} else if score == winners[0][1] {
			winners = append(winners, []int8{int8(i), score})
		}
	}
	g.winners = winners
}

func (g *Game) HandleRounds() {

	for round := 1; round <= 4; round++ {
		if g.stage > 0 {
			if g.stage == 1 {
				g.openCommunityCards(3)
			} else {
				g.openCommunityCards(1)
			}
		}
		g.startRound()
		// this is for all fold case situation ...
		if g.playersInGame == 1 {
			return
		}
		g.stage += 1
		g.raisedPlayer = g.firstPlayer
	}
}

// improve the logic for first round like which value should I pass to raisedPlayer
// because my bb also need chance to make decision in first round if I make is raisePlayer
func (g *Game) startRound() {
	i := g.raisedPlayer + 1

	for i != g.raisedPlayer {
		if !g.players[i].isFold {
			g.players[i].Play()

			if g.players[i].isRaised {
				g.raisedPlayer = i
				g.players[i].isRaised = false
			}

			if g.players[i].isFold {
				g.playersInGame--
			}
		}

		i++
		i %= g.totalPlayers
	}

}

func (g *Game) openCommunityCards(numberOfCards int8) {
	var i int8
	for i = 0; i < numberOfCards; i++ {
		cardValue, suit := g.cd.cardDistributer()
		g.communityCards = append(g.communityCards, []int8{cardValue, suit})
	}
}
