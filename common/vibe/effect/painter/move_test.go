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

func TestMovePaint(t *testing.T) {
	aTime := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aSecond := time.Second
	aFloat := 1.0

	cases := []PainterTest{
		{
			Name: "Paint all the colors",
			Painter: &Move{
				ColorStart: color.Red,
				Shifter: &shifter.Temporal{
					Start:    &aTime,
					Interval: &aSecond,
					Bender: &bender.Linear{
						Interval: &aFloat,
					},
				},
			},
			Instants: func() []Instant {
				insts := make([]Instant, len(color.AllColors))
				for i := range insts {
					insts[i] = Instant{
						Time:          aTime.Add(aSecond * time.Duration(i) / 24),
						ExpectedColor: color.AllColors[i],
					}
				}
				return insts
			}(),
		},
	}
	RunPainterTests(t, cases)
}

func TestMoveGetStabilizeFuncs(t *testing.T) {
	aTime1 := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aTime2 := time.Date(2009, 11, 17, 20, 34, 50, 652387237, time.UTC)
	aFloat := 1.1
	aDuration := time.Second
	c := helper.StabilizerTest{
		Stabilizer: &Move{},
		ExpectedVersions: []common.Stabilizer{
			&Move{
				ColorStart: color.Red,
			},
			&Move{
				ColorStart: color.Red,
				Shifter:    &shifter.Temporal{},
			},
			&Move{
				ColorStart: color.Red,
				Shifter: &shifter.Temporal{
					Start: &aTime2,
				},
			},
			&Move{
				ColorStart: color.Red,
				Shifter: &shifter.Temporal{
					Start:    &aTime2,
					Interval: &aDuration,
				},
			},
			&Move{
				ColorStart: color.Red,
				Shifter: &shifter.Temporal{
					Start:    &aTime2,
					Interval: &aDuration,
					Bender:   &bender.Static{},
				},
			},
			&Move{
				ColorStart: color.Red,
				Shifter: &shifter.Temporal{
					Start:    &aTime2,
					Interval: &aDuration,
					Bender: &bender.Static{
						TheBend: &aFloat,
					},
				},
			},
		},
		Palette: helper.TestPalette{
			Seed:     repeat.NewSeed(aTime1),
			Bender:   &bender.Static{},
			Duration: aDuration,
			Color:    color.Red,
			Shift:    aFloat,
			Shifter:  &shifter.Temporal{},
		},
	}
	helper.RunStabilizerTest(t, c)
}
