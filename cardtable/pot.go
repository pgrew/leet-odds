package cardtable

import "errors"

// Pot is the pile of chips up for grab in a card game
type Pot struct {
	chips       []*Chip
	minBetValue int32
}

// NewPot returns a new Pot instance
func NewPot() *Pot {
	return &Pot{make([]*Chip, 0), 0}
}

// Bet adds a pile of chips to the pot and sets the min bet after
func (p *Pot) Bet(theBet []*Chip) error {
	var total int64 = 0
	for _, c := range theBet {
		p.chips = append(p.chips, c)
		total += int64(c.GetValue())
	}
	if int64(p.GetMinimumBetAmount()) > total {
		return errors.New("Error: bet did not exceed the minimum allowed by pot")
	}
	// TODO: what if total exceeds int32 size?
	p.SetMinimumBetAmout(int32(total))
	return nil
}

// add puts more chips in the pot
func (p *Pot) add(c *Chip) {
	p.chips = append(p.chips, c)
}

// SetMinimumBetAmout sets the minimum value a bet can be
func (p *Pot) SetMinimumBetAmout(minBet int32) {
	p.minBetValue = minBet
}

// GetMinimumBetAmount returns the minimum bet value allowed by this pot
func (p *Pot) GetMinimumBetAmount() int32 {
	return p.minBetValue
}

// GetTotal returns the total value of this pot
func (p *Pot) GetTotal() (int64, error) {
	if len(p.chips) <= 0 {
		return 0, errors.New("Error: no chips in the pot!")
	}

	var total int64
	total = 0
	for _, c := range p.chips {
		total += int64(c.GetValue())
	}
	return total, nil
}

// Clear empties this pot
func (p *Pot) Clear() {
	p.chips = make([]*Chip, 0)
	p.minBetValue = 0
}
