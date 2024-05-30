package bender

import (
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestLinearBend(t *testing.T) {
	aFloat := 1.1
	cases := []BenderTest{
		{
			Name: "Paint Black",
			Bender: &Linear{
				Coefficient: aFloat,
			},
			Instants: []Instant{
				{
					Input:        -2.0,
					ExpectedBend: -2.0 * aFloat,
				},
				{
					Input:        -1.0,
					ExpectedBend: -1.0 * aFloat,
				},
				{
					Input:        0.0,
					ExpectedBend: 0.0 * aFloat,
				},
				{
					Input:        1.0,
					ExpectedBend: 1.0 * aFloat,
				},
				{
					Input:        2.0,
					ExpectedBend: 2.0 * aFloat,
				},
			},
		},
	}
	RunBenderTests(t, cases)
}

func TestLinearGetStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := testutil.StabilizerTest{
		Stabilizer: &Linear{
			Coefficient: aFloat,
		},
		ExpectedVersions: []common.Stabilizer{},
		Palette:          testutil.TestPalette{},
	}
	testutil.RunStabilizerTest(t, c)
}
