package cardtable

import "image/color"

// Material is the material chips are made from. Will be useful when players
// are splashing the pot.
type Material string

const (
	Clay Material = "CLAY"

	Plastic Material = "PLASTIC"

	Ceramic Material = "CERAMIC"
)

// Chip is a poker chip
type Chip struct {
	denom    int32
	color    color.Color
	material Material
}

// GetValue returns this chip's value
func (c Chip) GetValue() int32 {
	return c.denom
}
