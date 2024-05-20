package bender

import (
	"math"
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestSinusoidalBend(t *testing.T) {
	offset := 2.4
	period := 1.1
	amplitude := 3.7
	cases := []BenderTest{
		{
			Name: "Paint Black",
			Bender: &Sinusoidal{
				Offset:    &offset,
				Period:    &period,
				Amplitude: &amplitude,
			},
			Instants: []Instant{
				{
					Input:        -2.5,
					ExpectedBend: amplitude * math.Sin(offset+2*math.Pi*(-2.5/period)),
				},
				{
					Input:        -1.0,
					ExpectedBend: amplitude * math.Sin(offset+2*math.Pi*(-1.0/period)),
				},
				{
					Input:        0.0,
					ExpectedBend: amplitude * math.Sin(offset+2*math.Pi*(0.0/period)),
				},
				{
					Input:        1.0,
					ExpectedBend: amplitude * math.Sin(offset+2*math.Pi*(1.0/period)),
				},
				{
					Input:        2.5,
					ExpectedBend: amplitude * math.Sin(offset+2*math.Pi*(2.5/period)),
				},
			},
		},
	}
	RunBenderTests(t, cases)
}

func TestSinusoidalGetStabilizeFuncs(t *testing.T) {
	aFloat := 1.1
	c := testutil.StabilizerTest{
		Stabilizer: &Sinusoidal{},
		ExpectedVersions: []common.Stabilizer{
			&Sinusoidal{
				Offset: &aFloat,
			},
			&Sinusoidal{
				Offset: &aFloat,
				Period: &aFloat,
			},
			&Sinusoidal{
				Offset:    &aFloat,
				Period:    &aFloat,
				Amplitude: &aFloat,
			},
		},
		Palette: testutil.TestPalette{
			Shift: aFloat,
		},
	}
	testutil.RunStabilizerTest(t, c)
}
