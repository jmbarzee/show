package instruction

import (
	"time"
)

// An Heap is a min-heap of LightChange.
type (
	Heap []Instruction

	// Instruction is a single update to a device
	Instruction interface {
		GetTime() time.Time
	}
)

// Len implements the golang heap interface
func (h Heap) Len() int { return len(h) }

// Less implements the golang heap interface
func (h Heap) Less(i, j int) bool { return h[i].GetTime().Before(h[j].GetTime()) }

// Swap implements the golang heap interface
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push implements the golang heap interface
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Instruction)) }

// Pop implements the golang heap interface
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return &x
}

// Peek returns the next item in the Heap without removing it
func (h Heap) Peek() interface{} {
	if len(h) > 0 {
		return h[0]
	}
	return nil
}
