package shifter

import (
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestStaticShift(t *testing.T) {
	aFloat := 1.1
	cases := []ShiftTest{
		{
			Name: "One shift per second",
			Shifter: &Static{
				TheShift: &aFloat,
			},
			Instants: []Instant{
				{
					ExpectedShift: aFloat,
				},
			},
		},
	}
	RunShifterTests(t, cases)
}
func TestStaticGetStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := testutil.StabilizerTest{
		Stabilizer: &Static{},
		ExpectedVersions: []common.Stabilizer{
			&Static{
				TheShift: &aFloat,
			},
		},
		Palette: testutil.TestPalette{
			Shift: aFloat,
		},
	}
	testutil.RunStabilizerTest(t, c)
}
