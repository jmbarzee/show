package shifter

import (
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/testutil"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	helper "github.com/jmbarzee/show/common/vibe/testhelper"
)

func TestPositionalShift(t *testing.T) {
	aPosition := 5
	numPositions := 25
	aFloat := 1.1
	cases := []ShiftTest{
		{
			Name: "One shift per second",
			Shifter: &Positional{
				Bender: &bender.Static{
					TheBend: &aFloat,
				},
			},
			Instants: []Instant{
				{
					Light: &testutil.Light{
						Position:     aPosition,
						NumPositions: numPositions,
					},
					ExpectedShift: aFloat,
				},
			},
		},
		{
			Name: "One shift per second",
			Shifter: &Positional{
				Bender: &bender.Linear{
					Interval: &aFloat,
				},
			},
			Instants: []Instant{
				{
					Light: &testutil.Light{
						Position:     aPosition,
						NumPositions: numPositions,
					},
					ExpectedShift: float64(aPosition) / aFloat / float64(numPositions),
				},
			},
		},
	}
	RunShifterTests(t, cases)
}
func TestPositionalGetStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := helper.StabilizerTest{
		Stabilizer: &Positional{},
		ExpectedVersions: []common.Stabilizer{
			&Positional{
				Bender: &bender.Static{},
			},
			&Positional{
				Bender: &bender.Static{
					TheBend: &aFloat,
				},
			},
		},
		Palette: helper.TestPalette{
			Bender: &bender.Static{},
			Shift:  aFloat,
		},
	}
	helper.RunStabilizerTest(t, c)
}
