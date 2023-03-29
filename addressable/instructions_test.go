package addressable

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestLightPlan(t *testing.T) {
	time0 := time.Date(2009, 11, 17, 20, 34, 50, 651387237, time.UTC)
	time1 := time.Date(2009, 11, 17, 20, 34, 51, 651387237, time.UTC)
	time2 := time.Date(2009, 11, 17, 20, 34, 52, 651387237, time.UTC)
	iMin := Instruction{Time: time0}
	iMid := Instruction{Time: time1}
	iMax := Instruction{Time: time2}

	cases := []struct {
		Name       string
		Initial    Instructions
		Operations func(h Instructions) Instructions
		Expected   []Instruction
	}{
		{
			Name:    "Template",
			Initial: NewInstructions(),
			Operations: func(p Instructions) Instructions {
				return p
			},
			Expected: []Instruction{},
		},
		{
			Name:    "Single Instruction",
			Initial: NewInstructions(),
			Operations: func(p Instructions) Instructions {
				p.Add(iMin)
				return p
			},
			Expected: []Instruction{
				iMin,
			},
		},
		{
			Name:    "Tripple Instruction",
			Initial: NewInstructions(),
			Operations: func(p Instructions) Instructions {
				p.Add(iMin)
				p.Add(iMid)
				p.Add(iMax)
				return p
			},
			Expected: []Instruction{
				iMin,
				iMid,
				iMax,
			},
		},
		{
			Name:    "Tripple Instruction (out of order)",
			Initial: NewInstructions(),
			Operations: func(p Instructions) Instructions {
				p.Add(iMax)
				p.Add(iMid)
				p.Add(iMin)
				return p
			},
			Expected: []Instruction{
				iMin,
				iMid,
				iMax,
			},
		},
		{
			Name:    "Duplicate Instruction",
			Initial: NewInstructions(),
			Operations: func(p Instructions) Instructions {
				p.Add(iMid)
				p.Add(iMid)
				return p
			},
			Expected: []Instruction{
				iMid,
			},
		},
	}
	for i, c := range cases {
		t.Run("case_"+strconv.Itoa(i)+"_"+c.Name, func(t *testing.T) {
			actual := c.Operations(c.Initial)
			if len(c.Expected) > actual.heap.Len() {
				t.Fatalf("LightPlan length was not long enough to match\nExpect %v\nActual %v", len(c.Expected), actual.heap.Len())
			}
			for _, expectedChange := range c.Expected {
				timeAfter := expectedChange.GetTime().Add(time.Millisecond)
				actualChange := actual.Advance(timeAfter)
				if !reflect.DeepEqual(&expectedChange, actualChange) {
					t.Fatalf("LightPlan did not pop the expected Instruction\nExpect %v\nActual %v", &expectedChange, actualChange)
				}
			}
		})
	}
}
