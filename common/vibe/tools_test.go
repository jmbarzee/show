package vibe

import (
	"testing"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/testutil"
)

type StabilizeTest struct {
	Name          string
	ActualVibe    common.Vibe
	ExpectedVibes []common.Vibe
}

func RunStabilizeTests(t *testing.T, cases []StabilizeTest) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			actualVibe := c.ActualVibe
			for i, expectedVibe := range c.ExpectedVibes {

				actualVibe = actualVibe.Stabilize()
				if !testutil.StructsEqual(actualVibe, expectedVibe) {
					t.Fatalf("Stabilize %v failed. Vibes were not equal:\n\tExpected: %+v,\n\tActual: %+v", i, expectedVibe, actualVibe)
				}
			}
		})
	}
}

type MaterializeTest struct {
	Name         string
	ActualVibe   common.Vibe
	ExpectedVibe common.Vibe
}

func RunMaterializeTests(t *testing.T, cases []MaterializeTest) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			c.ActualVibe.Materialize()
			if !testutil.StructsEqual(c.ActualVibe, c.ExpectedVibe) {
				t.Fatalf("Materialize failed. Vibes were not equal:\n\tExpected: %+v,\n\tActual: %+v", c.ExpectedVibe, c.ActualVibe)
			}
		})
	}
}
