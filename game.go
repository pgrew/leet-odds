package main

import (
	"fmt"
)

type game struct {
	players     []Player
	board       Board
	button      Player
	currentAnte Chip
	pot         Pot
	maxPlayers  int
}

type Game interface {
	Deal() error
}

func main() {
	fmt.Println("Hello, playground")
}
