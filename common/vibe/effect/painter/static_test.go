package painter

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	helper "github.com/jmbarzee/show/common/vibe/testhelper"
)

func TestStaticPaint(t *testing.T) {
	aTime := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	cases := []PainterTest{
		{
			Name: "Paint Black",
			Painter: &Static{
				Color: color.Black,
			},
			Instants: []Instant{
				{
					Time:          aTime,
					ExpectedColor: color.Black,
				},
				{
					Time:          aTime.Add(time.Millisecond),
					ExpectedColor: color.Black,
				},
				{
					Time:          aTime.Add(time.Second),
					ExpectedColor: color.Black,
				},
				{
					Time:          aTime.Add(time.Minute),
					ExpectedColor: color.Black,
				},
				{
					Time:          aTime.Add(time.Hour),
					ExpectedColor: color.Black,
				},
			},
		},
		{
			Name: "Paint White",
			Painter: &Static{
				Color: color.White,
			},
			Instants: []Instant{
				{
					Time:          aTime,
					ExpectedColor: color.White,
				},
				{
					Time:          aTime.Add(time.Millisecond),
					ExpectedColor: color.White,
				},
				{
					Time:          aTime.Add(time.Second),
					ExpectedColor: color.White,
				},
				{
					Time:          aTime.Add(time.Minute),
					ExpectedColor: color.White,
				},
				{
					Time:          aTime.Add(time.Hour),
					ExpectedColor: color.White,
				},
			},
		},
		{
			Name: "Paint Blue",
			Painter: &Static{
				Color: color.Blue,
			},
			Instants: []Instant{
				{
					Time:          aTime,
					ExpectedColor: color.Blue,
				},
				{
					Time:          aTime.Add(time.Millisecond),
					ExpectedColor: color.Blue,
				},
				{
					Time:          aTime.Add(time.Second),
					ExpectedColor: color.Blue,
				},
				{
					Time:          aTime.Add(time.Minute),
					ExpectedColor: color.Blue,
				},
				{
					Time:          aTime.Add(time.Hour),
					ExpectedColor: color.Blue,
				},
			},
		},
	}
	RunPainterTests(t, cases)
}

func TestStaticGetStabilizeFuncs(t *testing.T) {
	c := helper.StabilizerTest{
		Stabilizer: &Static{},
		ExpectedVersions: []common.Stabilizer{
			&Static{
				Color: color.Blue,
			},
		},
		Palette: helper.TestPalette{
			Color: color.Blue,
		},
	}
	helper.RunStabilizerTest(t, c)
}
