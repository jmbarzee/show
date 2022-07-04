package bender

import (
	"testing"

	"github.com/jmbarzee/show/ifaces"
	helper "github.com/jmbarzee/show/vibe/testhelper"
)

type (
	BenderTest struct {
		Name     string
		Bender   ifaces.Bender
		Instants []Instant
	}

	Instant struct {
		Input        float64
		ExpectedBend float64
	}
)

func RunBenderTests(t *testing.T, cases []BenderTest) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for i, instant := range c.Instants {
				actualBend := c.Bender.Bend(instant.Input)
				if !helper.FloatsEqual(instant.ExpectedBend, actualBend, helper.MinErrColor) {
					t.Fatalf("instant %v failed:\n\tExpected: %v,\n\tActual: %v", i, instant.ExpectedBend, actualBend)
				}
			}
		})
	}
}
