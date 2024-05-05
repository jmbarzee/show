package common

import (
	"time"

	"github.com/jmbarzee/show/common/color"
)

// CopyEffect returns a deep copy of an Effect if its not nil
// Used by the Copy functions of palette providers
func CopyEffect(e Effect) Effect {
	if e == nil {
		return nil
	}
	return e.Copy()
}

// CopyPainter returns a deep copy of an Painter if its not nil
// Used by the Copy functions of palette providers
func CopyPainter(p Painter) Painter {
	if p == nil {
		return nil
	}
	return p.Copy()
}

// CopyShifter returns a deep copy of an Shifter if its not nil
// Used by the Copy functions of palette providers
func CopyShifter(s Shifter) Shifter {
	if s == nil {
		return nil
	}
	return s.Copy()
}

// CopyBender returns a deep copy of an Bender if its not nil
// Used by the Copy functions of palette providers
func CopyBender(b Bender) Bender {
	if b == nil {
		return nil
	}
	return b.Copy()
}

// CopyColor returns a deep copy of an color.Color if its not nil
// Used by the Copy functions of palette providers
func CopyColor(c color.Color) color.Color {
	if c == nil {
		return nil
	}
	return c.Copy()
}

// CopyDuration returns a deep copy of a *time.Duration if its not nil
// Used by the Copy functions of palette providers
func CopyDuration(d *time.Duration) *time.Duration {
	var dP *time.Duration
	if d != nil {
		dTemp := *d
		dP = &dTemp
	}

	return dP
}

// CopyTime returns a deep copy of a *time.Time if its not nil
// Used by the Copy functions of palette providers
func CopyTime(t *time.Time) *time.Time {
	var tP *time.Time
	if t != nil {
		tTemp := *t
		tP = &tTemp
	}

	return tP
}

// CopyFloat64 returns a deep copy of a *float64 if its not nil
// Used by the Copy functions of palette providers
func CopyFloat64(f *float64) *float64 {
	var fP *float64
	if f != nil {
		fTemp := *f
		fP = &fTemp
	}

	return fP
}

// CopyBool returns a deep copy of a *bool if its not nil
// Used by the Copy functions of palette providers
func CopyBool(b *bool) *bool {
	var bP *bool
	if b != nil {
		bTemp := *b
		bP = &bTemp
	}

	return bP
}
