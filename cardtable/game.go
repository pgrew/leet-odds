package cardtable

type game struct {
	rounds      []Round
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
