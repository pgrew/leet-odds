package cardtable

// Suit represents a color suit in cards
type Suit string

func (s Suit) String() string {
	return string(s)
}

const (
	Clubs    Suit = "CLUBS"
	Diamonds Suit = "DIAMONDS"
	Hearts   Suit = "HEARTS"
	Spades   Suit = "SPADES"
	Wild     Suit = "WILD"
)

// Card represents a standard (or wild) playing card
type Card struct {
	isHidden bool
	rank     int
	wildcard bool
	suit     Suit
	ace      bool
}

// NewCard returns a standard playing card
func NewCard(rank int, suit Suit, isAce bool) *Card {
	return &Card{
		rank:     rank,
		suit:     suit,
		wildcard: false,
		ace:      isAce,
	}
}

// NewWildCard returns a wildcard
func NewWildCard() *Card {
	return &Card{
		rank:     33,
		suit:     Wild,
		wildcard: true,
		ace:      true,
	}
}

// GetRank returns the rank of the card
func (c Card) GetRank() int {
	return c.rank
}

// IsWildcard returns true if this is a wildcard, otherwise false
func (c Card) IsWildcard() bool {
	return c.wildcard
}

// IsAce returns true if this is an ace, otherwise false
func (c Card) IsAce() bool {
	return c.ace
}

// SetAsWild changes this card instance into a wildcard
func (c *Card) SetAsWild() {
	c = NewWildCard()
}

// NewSuit generates an entire standard suit of cards (13 in total)
func NewSuit(suit Suit) []*Card {
	cards := make([]*Card, 13)
	cards = append(cards, NewCard(2, suit, false))
	cards = append(cards, NewCard(3, suit, false))
	cards = append(cards, NewCard(4, suit, false))
	cards = append(cards, NewCard(5, suit, false))
	cards = append(cards, NewCard(6, suit, false))
	cards = append(cards, NewCard(7, suit, false))
	cards = append(cards, NewCard(8, suit, false))
	cards = append(cards, NewCard(9, suit, false))
	cards = append(cards, NewCard(10, suit, false))
	cards = append(cards, NewCard(11, suit, false))
	cards = append(cards, NewCard(12, suit, false))
	cards = append(cards, NewCard(13, suit, false))
	cards = append(cards, NewCard(14, suit, true))
	return cards
}
