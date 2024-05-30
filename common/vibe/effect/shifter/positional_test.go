package shifter

import (
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/testutil"
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
					TheBend: aFloat,
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
					Coefficient: aFloat,
				},
			},
			Instants: []Instant{
				{
					Light: &testutil.Light{
						Position:     aPosition,
						NumPositions: numPositions,
					},
					ExpectedShift: float64(aPosition) * aFloat / float64(numPositions),
				},
			},
		},
	}
	RunShifterTests(t, cases)
}
func TestPositionalGetStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := testutil.StabilizerTest{
		Stabilizer: &Positional{},
		ExpectedVersions: []common.Stabilizer{
			&Positional{
				Bender: &bender.Static{
					TheBend: aFloat,
				},
			},
		},
		Palette: testutil.TestPalette{
			Bender: &bender.Static{TheBend: aFloat},
		},
	}
	testutil.RunStabilizerTest(t, c)
}
