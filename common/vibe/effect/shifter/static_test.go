package shifter

import (
	"testing"

	"github.com/jmbarzee/show/common"
	helper "github.com/jmbarzee/show/common/vibe/testhelper"
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
	c := helper.StabilizeableTest{
		Stabilizer: &Static{},
		ExpectedVersions: []common.Stabilizer{
			&Static{
				TheShift: &aFloat,
			},
		},
		Palette: helper.TestPalette{
			Shift: aFloat,
		},
	}
	helper.RunStabilizeableTest(t, c)
}
