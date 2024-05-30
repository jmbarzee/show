package effect

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/effect/painter"
	"github.com/jmbarzee/show/common/vibe/effect/shifter"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestSolidEffect(t *testing.T) {
	aTime := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aSecond := time.Second
	aFloat := 1.0
	numLights := 5
	cases := []EffectTest{
		{
			Name: "Solid Effect with Static Painter",
			Effect: &Solid{
				Painter: &painter.Static{
					Color: color.Blue,
				},
			},
			IntialLights: GetLights(numLights, color.Black),
			Instants: []Instant{
				{
					Time:           aTime,
					ExpectedLights: GetLights(numLights, color.Blue),
				},
				{
					Time:           aTime.Add(time.Millisecond * 1),
					ExpectedLights: GetLights(numLights, color.Blue),
				},
				{
					Time:           aTime.Add(time.Second * 1),
					ExpectedLights: GetLights(numLights, color.Blue),
				},
				{
					Time:           aTime.Add(time.Minute * 1),
					ExpectedLights: GetLights(numLights, color.Blue),
				},
				{
					Time:           aTime.Add(time.Hour * 1),
					ExpectedLights: GetLights(numLights, color.Blue),
				},
			},
		},
		{
			Name: "Solid Effect with Moving Painter",
			Effect: &Solid{
				Painter: &painter.Move{
					ColorStart: color.Blue,
					HueShifter: &shifter.Temporal{
						Start:    &aTime,
						Interval: &aSecond,
						Bender: &bender.Linear{
							Coefficient: aFloat,
						},
					},
				},
			},
			IntialLights: GetLights(numLights, color.Black),
			Instants: []Instant{
				{
					Time:           aTime.Add(time.Second * 0 / 24),
					ExpectedLights: GetLights(numLights, color.Blue),
				},
				{
					Time:           aTime.Add(time.Second * 4 / 24),
					ExpectedLights: GetLights(numLights, color.Magenta),
				},
				{
					Time:           aTime.Add(time.Second * 8 / 24),
					ExpectedLights: GetLights(numLights, color.Red),
				},
				{
					Time:           aTime.Add(time.Second * 12 / 24),
					ExpectedLights: GetLights(numLights, color.Yellow),
				},
				{
					Time:           aTime.Add(time.Second * 16 / 24),
					ExpectedLights: GetLights(numLights, color.Green),
				},
				{
					Time:           aTime.Add(time.Second * 20 / 24),
					ExpectedLights: GetLights(numLights, color.Cyan),
				},
				{
					Time:           aTime.Add(time.Second * 0 / 24),
					ExpectedLights: GetLights(numLights, color.Blue),
				},
			},
		},
	}
	RunEffectTests(t, cases)
}

func TestSolidGetStabilizeFuncs(t *testing.T) {
	c := testutil.StabilizerTest{
		Stabilizer: &Solid{},
		ExpectedVersions: []common.Stabilizer{
			&Solid{
				Painter: &painter.Static{},
			},
			&Solid{
				Painter: &painter.Static{
					Color: color.Blue,
				},
			},
		},
		Palette: testutil.TestPalette{
			Color:   color.Blue,
			Painter: &painter.Static{},
		},
	}
	testutil.RunStabilizerTest(t, c)
}
