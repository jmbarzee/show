package shifter

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/span"
	helper "github.com/jmbarzee/show/common/vibe/testhelper"
)

func TestTemporalShift(t *testing.T) {
	aTime := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aSecond := time.Second
	aFloat := 1.1
	cases := []ShiftTest{
		{
			Name: "One shift per second",
			Shifter: &Temporal{
				Start:    &aTime,
				Interval: &aSecond,
				Bender: &bender.Static{
					TheBend: &aFloat,
				},
			},
			Instants: []Instant{
				{
					Time:          aTime.Add(0 * time.Second),
					ExpectedShift: aFloat,
				},
				{
					Time:          aTime.Add(1 * time.Second),
					ExpectedShift: aFloat,
				},
				{
					Time:          aTime.Add(1 * time.Hour),
					ExpectedShift: aFloat,
				},
			},
		},
		{
			Name: "One shift per second",
			Shifter: &Temporal{
				Start:    &aTime,
				Interval: &aSecond,
				Bender: &bender.Linear{
					Interval: &aFloat,
				},
			},
			Instants: []Instant{
				{
					Time:          aTime.Add(0 * time.Second),
					ExpectedShift: 0 / aFloat,
				},
				{
					Time:          aTime.Add(1 * time.Second),
					ExpectedShift: 1 / aFloat,
				},
				{
					Time:          aTime.Add(1 * time.Hour),
					ExpectedShift: 3600 / aFloat,
				},
			},
		},
	}
	RunShifterTests(t, cases)
}

func TestTemporalGetStabilizeFuncs(t *testing.T) {
	aTime := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	aSecond := time.Second
	aFloat := 1.1
	c := helper.StabilizeableTest{
		Stabilizer: &Temporal{},
		ExpectedVersions: []common.Stabilizer{
			&Temporal{
				Start: &aTime,
			},
			&Temporal{
				Start:    &aTime,
				Interval: &aSecond,
			},
			&Temporal{
				Start:    &aTime,
				Interval: &aSecond,
				Bender:   &bender.Static{},
			},
			&Temporal{
				Start:    &aTime,
				Interval: &aSecond,
				Bender: &bender.Static{
					TheBend: &aFloat,
				},
			},
		},
		Palette: helper.TestPalette{
			Span: span.Span{
				StartTime: aTime,
			},
			Duration: aSecond,
			Bender:   &bender.Static{},
			Shift:    aFloat,
		},
	}
	helper.RunStabilizeableTest(t, c)
}
