package painter

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/repeat"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/effect/shifter"
	helper "github.com/jmbarzee/show/common/vibe/testhelper"
)

func TestBouncePaint(t *testing.T) {
	aTime := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aSecond := time.Second
	theTruth := true
	theFalsehood := false
	aFloat := 1.0

	cases := []PainterTest{
		{
			Name: "Bounce from Green to Cyan ",
			Painter: &Bounce{
				ColorStart: color.Green,
				ColorEnd:   color.Cyan,
				Up:         &theTruth,
				Shifter: &shifter.Temporal{
					Start:    &aTime,
					Interval: &aSecond,
					Bender: &bender.Linear{
						Interval: &aFloat,
					},
				},
			},
			Instants: []Instant{
				{
					Time:          aTime.Add(time.Second * 0 / 24),
					ExpectedColor: color.Green,
				},
				{
					Time:          aTime.Add(time.Second * 1 / 24),
					ExpectedColor: color.CoolGreen,
				},
				{
					Time:          aTime.Add(time.Second * 2 / 24),
					ExpectedColor: color.GreenCyan,
				},
				{
					Time:          aTime.Add(time.Second * 3 / 24),
					ExpectedColor: color.WarmCyan,
				},
				{
					Time:          aTime.Add(time.Second * 4 / 24),
					ExpectedColor: color.Cyan,
				},
				{
					Time:          aTime.Add(time.Second * 5 / 24),
					ExpectedColor: color.WarmCyan,
				},
				{
					Time:          aTime.Add(time.Second * 6 / 24),
					ExpectedColor: color.GreenCyan,
				},
				{
					Time:          aTime.Add(time.Second * 7 / 24),
					ExpectedColor: color.CoolGreen,
				},
				{
					Time:          aTime.Add(time.Second * 8 / 24),
					ExpectedColor: color.Green,
				},
				{
					Time:          aTime.Add(time.Second * 9 / 24),
					ExpectedColor: color.CoolGreen,
				},
			},
		},
		{
			Name: "Bounce from Cyan to Green ",
			Painter: &Bounce{
				ColorStart: color.Cyan,
				ColorEnd:   color.Green,
				Up:         &theFalsehood,
				Shifter: &shifter.Temporal{
					Start:    &aTime,
					Interval: &aSecond,
					Bender: &bender.Linear{
						Interval: &aFloat,
					},
				},
			},
			Instants: []Instant{
				{
					Time:          aTime.Add(time.Second * 0 / 24),
					ExpectedColor: color.Cyan,
				},
				{
					Time:          aTime.Add(time.Second * 1 / 24),
					ExpectedColor: color.WarmCyan,
				},
				{
					Time:          aTime.Add(time.Second * 2 / 24),
					ExpectedColor: color.GreenCyan,
				},
				{
					Time:          aTime.Add(time.Second * 3 / 24),
					ExpectedColor: color.CoolGreen,
				},
				{
					Time:          aTime.Add(time.Second * 4 / 24),
					ExpectedColor: color.Green,
				},
				{
					Time:          aTime.Add(time.Second * 5 / 24),
					ExpectedColor: color.CoolGreen,
				},
				{
					Time:          aTime.Add(time.Second * 6 / 24),
					ExpectedColor: color.GreenCyan,
				},
				{
					Time:          aTime.Add(time.Second * 7 / 24),
					ExpectedColor: color.WarmCyan,
				},
				{
					Time:          aTime.Add(time.Second * 8 / 24),
					ExpectedColor: color.Cyan,
				},
				{
					Time:          aTime.Add(time.Second * 9 / 24),
					ExpectedColor: color.WarmCyan,
				},
			},
		},
		{
			Name: "Bounce from Orange to RedMagenta ",
			Painter: &Bounce{
				ColorStart: color.Orange,
				ColorEnd:   color.RedMagenta,
				Up:         &theFalsehood,
				Shifter: &shifter.Temporal{
					Start:    &aTime,
					Interval: &aSecond,
					Bender: &bender.Linear{
						Interval: &aFloat,
					},
				},
			},
			Instants: []Instant{
				{
					Time:          aTime.Add(time.Second * 0 / 24),
					ExpectedColor: color.Orange,
				},
				{
					Time:          aTime.Add(time.Second * 1 / 24),
					ExpectedColor: color.WarmRed,
				},
				{
					Time:          aTime.Add(time.Second * 2 / 24),
					ExpectedColor: color.Red,
				},
				{
					Time:          aTime.Add(time.Second * 3 / 24),
					ExpectedColor: color.CoolRed,
				},
				{
					Time:          aTime.Add(time.Second * 4 / 24),
					ExpectedColor: color.RedMagenta,
				},
				{
					Time:          aTime.Add(time.Second * 5 / 24),
					ExpectedColor: color.CoolRed,
				},
				{
					Time:          aTime.Add(time.Second * 6 / 24),
					ExpectedColor: color.Red,
				},
				{
					Time:          aTime.Add(time.Second * 7 / 24),
					ExpectedColor: color.WarmRed,
				},
				{
					Time:          aTime.Add(time.Second * 8 / 24),
					ExpectedColor: color.Orange,
				},
				{
					Time:          aTime.Add(time.Second * 9 / 24),
					ExpectedColor: color.WarmRed,
				},
			},
		},
		{
			Name: "Bounce from RedMagenta to Orange ",
			Painter: &Bounce{
				ColorStart: color.RedMagenta,
				ColorEnd:   color.Orange,
				Up:         &theTruth,
				Shifter: &shifter.Temporal{
					Start:    &aTime,
					Interval: &aSecond,
					Bender: &bender.Linear{
						Interval: &aFloat,
					},
				},
			},
			Instants: []Instant{
				{
					Time:          aTime.Add(time.Second * 0 / 24),
					ExpectedColor: color.RedMagenta,
				},
				{
					Time:          aTime.Add(time.Second * 1 / 24),
					ExpectedColor: color.CoolRed,
				},
				{
					Time:          aTime.Add(time.Second * 2 / 24),
					ExpectedColor: color.Red,
				},
				{
					Time:          aTime.Add(time.Second * 3 / 24),
					ExpectedColor: color.WarmRed,
				},
				{
					Time:          aTime.Add(time.Second * 4 / 24),
					ExpectedColor: color.Orange,
				},
				{
					Time:          aTime.Add(time.Second * 5 / 24),
					ExpectedColor: color.WarmRed,
				},
				{
					Time:          aTime.Add(time.Second * 6 / 24),
					ExpectedColor: color.Red,
				},
				{
					Time:          aTime.Add(time.Second * 7 / 24),
					ExpectedColor: color.CoolRed,
				},
				{
					Time:          aTime.Add(time.Second * 8 / 24),
					ExpectedColor: color.RedMagenta,
				},
				{
					Time:          aTime.Add(time.Second * 9 / 24),
					ExpectedColor: color.CoolRed,
				},
			},
		},
	}
	RunPainterTests(t, cases)
}

func TestBounceGetStabilizeFuncs(t *testing.T) {
	aTime1 := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aTime2 := aTime1.Add(time.Millisecond * 2)
	aDuration := time.Second
	theTruth := false
	aFloat := 1.1
	c := helper.StabilizerTest{
		Stabilizer: &Bounce{},
		ExpectedVersions: []common.Stabilizer{
			&Bounce{
				ColorStart: color.Red,
			},
			&Bounce{
				ColorStart: color.Red,
				ColorEnd:   color.Red,
			},
			&Bounce{
				ColorStart: color.Red,
				ColorEnd:   color.Red,
				Up:         &theTruth,
			},
			&Bounce{
				ColorStart: color.Red,
				ColorEnd:   color.Red,
				Up:         &theTruth,
				Shifter:    &shifter.Temporal{},
			},
			&Bounce{
				ColorStart: color.Red,
				ColorEnd:   color.Red,
				Up:         &theTruth,
				Shifter: &shifter.Temporal{
					Start: &aTime2,
				},
			},
			&Bounce{
				ColorStart: color.Red,
				ColorEnd:   color.Red,
				Up:         &theTruth,
				Shifter: &shifter.Temporal{
					Start:    &aTime2,
					Interval: &aDuration,
				},
			},
			&Bounce{
				ColorStart: color.Red,
				ColorEnd:   color.Red,
				Up:         &theTruth,
				Shifter: &shifter.Temporal{
					Start:    &aTime2,
					Interval: &aDuration,
					Bender:   &bender.Linear{},
				},
			},
			&Bounce{
				ColorStart: color.Red,
				ColorEnd:   color.Red,
				Up:         &theTruth,
				Shifter: &shifter.Temporal{
					Start:    &aTime2,
					Interval: &aDuration,
					Bender: &bender.Linear{
						Interval: &aFloat,
					},
				},
			},
		},
		Palette: helper.TestPalette{
			Seed:     repeat.NewSeed(aTime1),
			Bender:   &bender.Linear{},
			Duration: aDuration,
			Color:    color.Red,
			Shift:    aFloat,
			Shifter:  &shifter.Temporal{},
		},
	}
	helper.RunStabilizerTest(t, c)
}
