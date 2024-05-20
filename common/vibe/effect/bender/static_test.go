package bender

import (
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestStaticBend(t *testing.T) {
	aFloat := 1.1
	cases := []BenderTest{
		{
			Name: "Paint Black",
			Bender: &Static{
				TheBend: &aFloat,
			},
			Instants: []Instant{
				{
					Input:        -2.0,
					ExpectedBend: aFloat,
				},
				{
					Input:        -1.0,
					ExpectedBend: aFloat,
				},
				{
					Input:        0.0,
					ExpectedBend: aFloat,
				},
				{
					Input:        1.0,
					ExpectedBend: aFloat,
				},
				{
					Input:        2.0,
					ExpectedBend: aFloat,
				},
			},
		},
	}
	RunBenderTests(t, cases)
}

func TestStaticGetStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := testutil.StabilizerTest{
		Stabilizer: &Static{},
		ExpectedVersions: []common.Stabilizer{
			&Static{
				TheBend: &aFloat,
			},
		},
		Palette: testutil.TestPalette{
			Shift: aFloat,
		},
	}
	testutil.RunStabilizerTest(t, c)
}
