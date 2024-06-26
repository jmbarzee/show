package bender

import (
	"fmt"
	"math"

	"github.com/jmbarzee/show/common"
)

// Exponential is a Bender which provides a single unchanging bend
type Exponential struct {
	Exponent    float64
	Coefficient float64
}

var _ common.Bender = (*Exponential)(nil)

// Bend returns a value representing some change or bend
func (b Exponential) Bend(f float64) float64 {
	bend := b.Coefficient * math.Pow(math.Abs(f), b.Exponent)
	return bend
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (b *Exponential) GetStabilizeFuncs() []func(p common.Palette) {
	return []func(p common.Palette){}
}

// Copy returns a deep copy of the Bender
func (b Exponential) Copy() common.Bender {
	return &Exponential{
		Exponent:    b.Exponent,
		Coefficient: b.Coefficient,
	}
}

// String returns a string representation of the Bender
func (b Exponential) String() string {
	return fmt.Sprintf("shifter.Exponential{Exponent:%v, Coefficient:%v}", b.Exponent, b.Coefficient)
}
