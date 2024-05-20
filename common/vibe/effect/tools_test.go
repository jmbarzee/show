package effect

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

type (
	EffectTest struct {
		Name         string
		Effect       common.Effect
		IntialLights []common.Renderable
		Instants     []Instant
	}

	Instant struct {
		Time           time.Time
		ExpectedLights []common.Renderable
	}
)

func RunEffectTests(t *testing.T, cases []EffectTest) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for i, instant := range c.Instants {
				for _, light := range c.IntialLights {
					c.Effect.Render(instant.Time, light)
				}

				for j, expectedLight := range instant.ExpectedLights {
					actualLight := c.IntialLights[j]
					if !testutil.ColorsEqual(expectedLight.GetColor(), actualLight.GetColor()) {
						t.Fatalf("instant %v, light %v failed:\n\tExpected: %v,\n\tActual: %v", i, j, expectedLight.GetColor(), actualLight.GetColor())
					}
				}
			}
		})
	}
}

func GetLights(length int, c color.HSL) []common.Renderable {
	lights := make([]common.Renderable, length)
	for i := range lights {
		lights[i] = &testutil.Light{
			Position:     i,
			NumPositions: length,
			Color:        c,
		}
	}
	return lights
}
