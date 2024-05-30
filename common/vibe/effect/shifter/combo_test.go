package shifter

import (
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestComboShift(t *testing.T) {
	aPosition := 5
	numPositions := 25
	aFloat := 1.1
	cases := []ShiftTest{
		{
			Name: "One shift per second",
			Shifter: &Combo{
				A: &Static{
					TheShift: aFloat,
				},
				B: &Positional{
					Bender: &bender.Linear{
						Coefficient: aFloat,
					},
				},
			},
			Instants: []Instant{
				{
					Light: &testutil.Light{
						Position:     aPosition,
						NumPositions: numPositions,
					},
					ExpectedShift: aFloat + float64(aPosition)*aFloat/float64(numPositions),
				},
			},
		},
	}
	RunShifterTests(t, cases)
}
func TestComboGetStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := testutil.StabilizerTest{
		Stabilizer: &Combo{},
		ExpectedVersions: []common.Stabilizer{
			&Combo{
				A: &Static{
					TheShift: aFloat,
				},
			},
			&Combo{
				A: &Static{
					TheShift: aFloat,
				},
				B: &Static{
					TheShift: aFloat,
				},
			},
		},
		Palette: testutil.TestPalette{
			Shifter: &Static{TheShift: aFloat},
		},
	}
	testutil.RunStabilizerTest(t, c)
}
