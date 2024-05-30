package palette

import (
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/vibe/effect"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/effect/painter"
	"github.com/jmbarzee/show/common/vibe/effect/shifter"
)

// NewRandom Provides a Palette which selects from all
// possible colors, effects, etc.
func NewRandom(s common.Seeder) *Basic {
	return &Basic{
		Seeder: s,
		Benders: []common.Bender{
			//&bender.Static{},
			&bender.Linear{Coefficient: 1.1},
			&bender.Exponential{Exponent: 2.0, Coefficient: 1.1},
			&bender.Sinusoidal{Offset: 0.5, Period: 1.0, Amplitude: 1.4},
			&bender.Combo{},
		},
		Colors: color.AllColors,
		Durations: []time.Duration{
			time.Second / 10,
			time.Second / 5,
			time.Second / 2,
			time.Second,
			time.Second * 2,
			time.Second * 5,
			time.Second * 10,
		},
		Shifters: []common.Shifter{
			//&shifter.Static{},
			&shifter.Positional{},
			&shifter.Locational{},
			// &shifter.Directional{},
			&shifter.Temporal{},
			&shifter.Combo{},
		},
		Painters: []common.Painter{
			//&painter.Static{},
			&painter.Move{},
			&painter.Bounce{},
		},
		Effects: []common.Effect{
			&effect.Solid{},
			&effect.Future{},
		},
	}

}
