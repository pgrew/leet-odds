package main

type Table struct {
	players []Player
	mainPot Pot
	sidePot Pot
	badBeat Pot
	board Board
	deck Deck
}
