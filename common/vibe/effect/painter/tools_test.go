package painter

import (
	"testing"
	"time"

	"github.com/jmbarzee/color"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/testutil"
	helper "github.com/jmbarzee/show/common/vibe/testhelper"
)

type (
	PainterTest struct {
		Name     string
		Painter  common.Painter
		Instants []Instant
	}

	Instant struct {
		Time          time.Time
		Light         common.Renderable
		ExpectedColor color.Color
	}
)

func RunPainterTests(t *testing.T, cases []PainterTest) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for i, instant := range c.Instants {
				instant.Light = &testutil.Light{}
				c.Painter.Paint(instant.Time, instant.Light)
				actualColor := instant.Light.GetColor()
				if !helper.ColorsEqual(instant.ExpectedColor, actualColor) {
					t.Fatalf("instant %v failed:\n\tExpected: %v,\n\tActual: %v", i, instant.ExpectedColor, actualColor)
				}
			}
		})
	}
}
