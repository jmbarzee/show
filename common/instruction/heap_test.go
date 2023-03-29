package instruction

import (
	"container/heap"
	"reflect"
	"strconv"
	"testing"
	"time"
)

type realInstruction struct {
	t time.Time
}

func (i realInstruction) GetTime() time.Time { return i.t }

func TestHeap(t *testing.T) {
	atime := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	riMin := realInstruction{
		t: atime.Add(time.Hour * -1),
	}
	riMid := realInstruction{
		t: atime,
	}
	riMax := realInstruction{
		t: atime.Add(time.Hour),
	}

	cases := []struct {
		Name       string
		Initial    Heap
		Operations func(h *Heap) *Heap
		Expected   Heap
	}{
		{
			Name:    "Template",
			Initial: Heap([]Instruction{}),
			Operations: func(h *Heap) *Heap {
				return h
			},
			Expected: Heap{},
		},
		{
			Name:    "Empty push",
			Initial: Heap([]Instruction{}),
			Operations: func(h *Heap) *Heap {
				heap.Push(h, riMin)
				return h
			},
			Expected: Heap{
				riMin,
			},
		},
		{
			Name:    "Ordered push",
			Initial: Heap([]Instruction{}),
			Operations: func(h *Heap) *Heap {
				heap.Push(h, riMin)
				heap.Push(h, riMid)
				heap.Push(h, riMax)
				return h
			},
			Expected: Heap{
				riMin,
				riMid,
				riMax,
			},
		},
		{
			Name:    "Unordered push",
			Initial: Heap([]Instruction{}),
			Operations: func(h *Heap) *Heap {
				heap.Push(h, riMax)
				heap.Push(h, riMid)
				heap.Push(h, riMin)
				return h
			},
			Expected: Heap{
				riMin,
				riMid,
				riMax,
			},
		},
		{
			Name: "Unordered push on filled heap",
			Initial: Heap([]Instruction{
				riMin,
				riMid,
				riMax}),
			Operations: func(h *Heap) *Heap {
				heap.Push(h, riMax)
				heap.Push(h, riMid)
				heap.Push(h, riMin)
				return h
			},
			Expected: Heap{
				riMin,
				riMin,
				riMid,
				riMid,
				riMax,
				riMax,
			},
		},
	}
	for i, c := range cases {
		t.Run("case_"+strconv.Itoa(i)+"_"+c.Name, func(t *testing.T) {
			actual := c.Operations(&c.Initial)
			if len(c.Expected) != len(*actual) {
				t.Fatalf("Heap lengths did not match\nExpect %v\nActual %v", len(c.Expected), len(*actual))
			}
			for _, expectedChange := range c.Expected {
				peekedChange := c.Initial.Peek()
				if !reflect.DeepEqual(expectedChange, peekedChange) {
					t.Fatalf("Heap did not pop the expected realInstruction\nExpect %v\nPeeked %v", expectedChange, peekedChange)
				}
				actualChange := heap.Pop(&c.Initial)
				if !reflect.DeepEqual(&expectedChange, actualChange) {
					t.Fatalf("Heap did not pop the expected realInstruction\nExpect %v\nActual %v", &expectedChange, actualChange)
				}

			}
		})
	}
}
