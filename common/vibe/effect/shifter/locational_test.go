package shifter

import (
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/space"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/testutil"
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
					Light: &testutil.Light{
						Object: space.NewObject(
							space.Vector{X: 1, Y: 2, Z: 3},
							*space.NewIdentityQuaternion(),
						),
					},
					ExpectedShift: aFloat * 3,
				},
				{
					Light: &testutil.Light{
						Object: space.NewObject(
							space.Vector{X: -1, Y: -2, Z: -3},
							*space.NewIdentityQuaternion(),
						),
					},
					ExpectedShift: aFloat * 3,
				},
				{
					Light: &testutil.Light{
						Object: space.NewObject(
							space.Vector{X: 0, Y: 0, Z: 0},
							*space.NewIdentityQuaternion(),
						),
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
	c := testutil.StabilizerTest{
		Stabilizer: &Locational{},
		ExpectedVersions: []common.Stabilizer{
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
					TheBend: &aFloat, // This is a little dirty. The Benders are both/all pointing to the same struct, so Bend is set with the first bender
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
		Palette: testutil.TestPalette{
			Bender: &bender.Static{},
			Shift:  aFloat,
		},
	}
	testutil.RunStabilizerTest(t, c)
}
