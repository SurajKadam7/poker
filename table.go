package main

import "fmt"

type Table struct {
	players      *[]Player
	minPlayers   int8
	maxPlayers   int8
	deler        int8
	sbAmount     int64
	bbAmount     int64
	totalPlayers int8
	minBuyin     int64
}

func (t *Table) addPlayers() bool {
	// minimum players required on the table to start the game
	if t.totalPlayers > t.maxPlayers {
		fmt.Println("The table is already full ...")
		return false
	}

	// initializing the playyers array
	if t.totalPlayers == 0 {
		players := make([]Player, t.maxPlayers)
		t.players = &players
	}

	var name string
	var amount int64
	fmt.Println("Enter your Name")
	fmt.Scanln(&name)
	fmt.Println("Add your buyin amount")
	fmt.Scanln(&amount)

	p := Player{name: name, amount: amount}

	// minimum amount required to play on the table
	if t.minBuyin > p.amount {
		fmt.Print("minimum buyin amount should be : ", t.minBuyin)
		return false
	}

	(*t.players) = append((*t.players), p)
	t.totalPlayers++
	return true
}

func (t *Table) startNewGame() {
	if t.totalPlayers < t.minPlayers {
		return
	}

	var sb int8

	if t.totalPlayers == 2 {
		sb = t.deler
	} else {
		sb = (t.deler + 1) % t.totalPlayers
	}

	bb := (sb + 1) % t.totalPlayers

	g := Game{
		players:  *t.players,
		sb:       sb,
		bb:       bb,
		bbAmout:  t.bbAmount,
		sbAmount: t.sbAmount,
	}

	// continue playing the game ..
	g.startGame()
}
