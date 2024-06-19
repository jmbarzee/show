package addressable

import (
	"container/heap"
	"sync"
	"time"

	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/instruction"
)

type Instruction struct {
	Time   time.Time
	Colors []color.Color
}

func (i Instruction) GetTime() time.Time { return i.Time }

// Instructions is an set of Instructions which are ordered by their Time
type Instructions struct {
	*sync.Mutex
	heap *instruction.Heap
}

// NewInstructions initializes and returns a Instructions
func NewInstructions() Instructions {
	instHeap := instruction.Heap{}
	heap.Init(&instHeap)
	return Instructions{
		Mutex: &sync.Mutex{},
		heap:  &instHeap,
	}
}

// Add inserts a Instruction into a Instructions
func (p Instructions) Add(c Instruction) {
	p.Lock()
	defer p.Unlock()
	heap.Push(p.heap, c)
}

// Advance drops all Instructions before t and returns the most recent Instruction
func (p Instructions) Advance(t time.Time) *Instruction {
	var change *Instruction
	p.Lock()
	defer p.Unlock()

	if p.heap.Len() > 0 {
		var past Instruction
		var next Instruction
		next = p.heap.Peek().(Instruction)

		for next.GetTime().Before(t) {
			past = next
			heap.Pop(p.heap)
			change = &past

			if p.heap.Len() > 0 {
				next = p.heap.Peek().(Instruction)
			} else {
				break
			}
		}
	}

	return change
}
