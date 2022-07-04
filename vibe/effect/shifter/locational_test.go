package shifter

import (
	"testing"

	"github.com/jmbarzee/show/ifaces"
	"github.com/jmbarzee/show/light"
	"github.com/jmbarzee/show/vibe/effect/bender"
	helper "github.com/jmbarzee/show/vibe/testhelper"
	"github.com/jmbarzee/space"
)

func TestLocationalShift(t *testing.T) {
	aFloat := 1.1
	cases := []ShiftTest{
		{
			Name: "One shift per second",
			Shifter: &Locational{
				XBender: &bender.Static{
					TheBend: &aFloat,
				},
				YBender: &bender.Static{
					TheBend: &aFloat,
				},
				ZBender: &bender.Static{
					TheBend: &aFloat,
				},
			},
			Instants: []Instant{
				{
					Light: &light.Basic{
						Location: space.Cartesian{
							X: 1,
							Y: 2,
							Z: 3,
						},
					},
					ExpectedShift: aFloat * 3,
				},
				{
					Light: &light.Basic{
						Location: space.Cartesian{
							X: -1,
							Y: -2,
							Z: -3,
						},
					},
					ExpectedShift: aFloat * 3,
				},
				{
					Light: &light.Basic{
						Location: space.Cartesian{
							X: 0,
							Y: 0,
							Z: 0,
						},
					},
					ExpectedShift: aFloat * 3,
				},
			},
		},
	}
	RunShifterTests(t, cases)
}
func TestLocationalGetStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := helper.StabilizeableTest{
		Stabalizable: &Locational{},
		ExpectedVersions: []ifaces.Stabalizable{
			&Locational{
				XBender: &bender.Static{},
			},
			&Locational{
				XBender: &bender.Static{
					TheBend: &aFloat,
				},
			},
			&Locational{
				XBender: &bender.Static{
					TheBend: &aFloat,
				},
				YBender: &bender.Static{
					TheBend: &aFloat, // This is a little dirty. The Benders are both/all pointing to the same struct, so TheBend is set with the first bender
				},
			},
			&Locational{
				XBender: &bender.Static{
					TheBend: &aFloat,
				},
				YBender: &bender.Static{
					TheBend: &aFloat,
				},
				ZBender: &bender.Static{
					TheBend: &aFloat,
				},
			},
		},
		Palette: helper.TestPalette{
			Bender: &bender.Static{},
			Shift:  aFloat,
		},
	}
	helper.RunStabilizeableTest(t, c)
}
