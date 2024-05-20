package effect

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/effect/painter"
	"github.com/jmbarzee/show/common/vibe/effect/shifter"
	"github.com/jmbarzee/show/common/vibe/span"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

func TestFutureEffect(t *testing.T) {
	aTime := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aSpan := span.New(aTime, aTime.Add(time.Hour))
	aFloat := 1.0
	aSecond := time.Second
	a24thSecond := time.Second / 24
	numLights := 5
	cases := []EffectTest{
		{
			Name: "Future Effect with Static Painter",
			Effect: &Future{
				BasicEffect:  BasicEffect{Spanner: aSpan},
				TimePerLight: &aSecond,
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
			Name: "Future Effect with Moving Painter",
			Effect: &Future{
				BasicEffect:  BasicEffect{Spanner: aSpan},
				TimePerLight: &a24thSecond,
				Painter: &painter.Move{
					ColorStart: color.Blue,
					Shifter: &shifter.Temporal{
						Start:    &aTime,
						Interval: &aSecond,
						Bender: &bender.Linear{
							Interval: &aFloat,
						},
					},
				},
			},
			IntialLights: GetLights(3, color.Black),
			Instants: []Instant{
				{
					Time: aTime.Add(time.Second * 0 / 24),
					ExpectedLights: []common.Renderable{
						&testutil.Light{Color: color.Blue},
						&testutil.Light{Color: color.WarmBlue},
						&testutil.Light{Color: color.Violet},
					},
				},
				{
					Time: aTime.Add(time.Second * 1 / 24),
					ExpectedLights: []common.Renderable{
						&testutil.Light{Color: color.WarmBlue},
						&testutil.Light{Color: color.Violet},
						&testutil.Light{Color: color.CoolMagenta},
					},
				},
				{
					Time: aTime.Add(time.Second * 2 / 24),
					ExpectedLights: []common.Renderable{
						&testutil.Light{Color: color.Violet},
						&testutil.Light{Color: color.CoolMagenta},
						&testutil.Light{Color: color.Magenta},
					},
				},
			},
		},
	}
	RunEffectTests(t, cases)
}

func TestFutureGetStabilizeFuncs(t *testing.T) {
	aSecond := time.Second
	c := testutil.StabilizerTest{
		Stabilizer: &Future{},
		ExpectedVersions: []common.Stabilizer{
			&Future{
				TimePerLight: &aSecond,
			},
			&Future{
				TimePerLight: &aSecond,
				Painter:      &painter.Static{},
			},
			&Future{
				TimePerLight: &aSecond,
				Painter: &painter.Static{
					Color: color.Blue,
				},
			},
		},
		Palette: testutil.TestPalette{
			Color:    color.Blue,
			Painter:  &painter.Static{},
			Duration: aSecond,
		},
	}
	testutil.RunStabilizerTest(t, c)
}
