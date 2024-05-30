package bender

import (
	"fmt"

	"github.com/jmbarzee/show/common"
)

// Linear is a Bender which provides a bend which changes linearly
type Linear struct {
	Coefficient float64
}

var _ common.Bender = (*Linear)(nil)

// Bend returns a value representing some change or bend
func (b Linear) Bend(f float64) float64 {
	bend := f * b.Coefficient
	return bend
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (b *Linear) GetStabilizeFuncs() []func(p common.Palette) {
	return []func(p common.Palette){}
}

// Copy returns a deep copy of the Bender
func (b Linear) Copy() common.Bender {
	return &Linear{
		Coefficient: b.Coefficient,
	}
}

// String returns a string representation of the Bender
func (b Linear) String() string {
	return fmt.Sprintf("shifter.Linear{Interval:%v}", b.Coefficient)
}
