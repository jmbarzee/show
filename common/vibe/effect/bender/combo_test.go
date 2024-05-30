package bender

import (
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestComboBend(t *testing.T) {
	aFloat1 := 1.1
	aFloat2 := 2.3
	cases := []BenderTest{
		{
			Name: "Paint Black",
			Bender: &Combo{
				A: &Static{
					TheBend: aFloat1,
				},
				B: &Linear{
					Coefficient: aFloat2,
				},
			},
			Instants: []Instant{
				{
					Input:        0.0,
					ExpectedBend: aFloat1 + 0.0*aFloat2,
				},
				{
					Input:        1.0,
					ExpectedBend: aFloat1 + 1.0*aFloat2,
				},
			},
		},
	}
	RunBenderTests(t, cases)
}

func TestComboStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := testutil.StabilizerTest{
		Stabilizer: &Combo{},
		ExpectedVersions: []common.Stabilizer{
			&Combo{
				A: &Static{
					TheBend: aFloat,
				},
			},
			&Combo{
				A: &Static{
					TheBend: aFloat,
				},
				B: &Static{
					TheBend: aFloat,
				},
			},
		},
		Palette: testutil.TestPalette{
			Bender: &Static{TheBend: aFloat},
		},
	}
	testutil.RunStabilizerTest(t, c)
}
