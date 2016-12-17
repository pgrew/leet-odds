package cardtable

import "errors"

// Deck represents a deck of cards. Decks are assembled, shuffled, and dealt.
type Deck struct {
	// All cards held in this deck
	deck []*Card `json:"deck"`
}

// Clear removes all cards from the deck
func (d *Deck) Clear() {
	d.deck = []*Card{}
}

// Assemble builds a deck of cards from a given game
func (d *Deck) Assemble(game Game) {
	d.Clear()
	d.assembleStandardDeck()
	// if game.Type == "Standard" {
	// 	d.assembleStandardDeck()
	// }
}

// addCards adds cards to the deck
func (d *Deck) addCards(moreCards []*Card) {
	for _, c := range moreCards {
		d.deck = append(d.deck, c)
	}
}

// assembleStandardDeck makes a 52 card, regular deck
func (d *Deck) assembleStandardDeck() {
	d.addCards(NewSuit(Clubs))
	d.addCards(NewSuit(Diamonds))
	d.addCards(NewSuit(Hearts))
	d.addCards(NewSuit(Spades))
}

// Shuffle randomly rearranges the cards in this deck
func (d *Deck) Shuffle() {
	// TODO shuffle
}

// GetCard gets a card off the top of the deck
func (d *Deck) GetCard() (*Card, error) {
	if len(d.deck) <= 0 {
		return nil, errors.New("Deck is out of cards!")
	}
	// Give away the "top" card in the deck (array position 0)
	c := d.deck[0]
	// Remove that card
	d.deck = append(d.deck[:0], d.deck[1:]...)
	// append(d.deck[:], d.deck[1:]...)
	return c, nil
}
