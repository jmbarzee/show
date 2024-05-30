package bender

import (
	"math"
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestExponentialBend(t *testing.T) {
	exponent := 2.4
	coefficient := 1.1
	cases := []BenderTest{
		{
			Name: "Paint Black",
			Bender: &Exponential{
				Exponent:    exponent,
				Coefficient: coefficient,
			},
			Instants: []Instant{
				{
					Input:        -2.0,
					ExpectedBend: coefficient * math.Pow(2.0, exponent),
				},
				{
					Input:        -1.0,
					ExpectedBend: coefficient * math.Pow(1.0, exponent),
				},
				{
					Input:        0.0,
					ExpectedBend: coefficient * math.Pow(0.0, exponent),
				},
				{
					Input:        1.0,
					ExpectedBend: coefficient * math.Pow(1.0, exponent),
				},
				{
					Input:        2.0,
					ExpectedBend: coefficient * math.Pow(2.0, exponent),
				},
			},
		},
	}
	RunBenderTests(t, cases)
}

func TestExponentialGetStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := testutil.StabilizerTest{
		Stabilizer: &Exponential{
			Exponent:    aFloat,
			Coefficient: aFloat,
		},
		ExpectedVersions: []common.Stabilizer{},
		Palette:          testutil.TestPalette{},
	}
	testutil.RunStabilizerTest(t, c)
}
