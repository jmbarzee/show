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
	durations := []time.Duration{
		time.Second / 10,
		time.Second / 5,
		time.Second / 2,
		time.Second,
		time.Second * 2,
		time.Second * 5,
		time.Second * 10,
	}

	shifts := []float64{
		0.005,
		0.01,
		0.02,
		0.05,
		0.1,
		0.2,
		0.5,
	}

	return &Basic{
		Seeder: s,
		Benders: []common.Bender{
			//&bender.Static{},
			&bender.Linear{},
			&bender.Exponential{},
			&bender.Sinusoidal{},
			&bender.Combo{},
		},
		Colors:    color.AllColors,
		Durations: durations,
		Shifts:    shifts,
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
			&effect.Solid{
				BasicEffect: effect.BasicEffect{Spanner: s},
			},
			&effect.Future{
				BasicEffect: effect.BasicEffect{Spanner: s},
			},
		},
	}

}
