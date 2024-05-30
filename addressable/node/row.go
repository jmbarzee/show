package node

import (
	"sort"
	"time"

	"github.com/jmbarzee/show/addressable"
	"github.com/jmbarzee/show/common"
)

// row holds basic functionality for any neoPixelDevice
// row implements effect.Allocater
// row partially implements effect.Device with PruneEffect
type row struct {
	// total is the number of LEDs in the row
	total int
	// spacing is the distance between each LED
	spacing addressable.Spacing

	// effects is the array of effects from materializing Vibes in calls to Allocate
	effects []common.Effect

	// lightsCache is a cached version of the lights
	// lightsCache is only read in row, but can be written by scructs which embed row
	lightsCache []*addressable.Light
}

// NewRow creates a new Row
func NewRow(spacing addressable.Spacing, count int) *row {
	return &row{
		total:   count,
		spacing: spacing,
		effects: make([]common.Effect, 0),
	}
}

// Allocate takes Vibes and Materializes them into effects
// This is the bottom of the Allocater hiarchy for neoPixels
func (r *row) Allocate(feeling common.Vibe) {
	newEffects := feeling.Materialize()
	r.effects = append(r.effects, newEffects...)
	sort.Sort(byStartTime(r.effects))
}

// Clean removes all effects which have ended before a time t
func (r *row) Clean(t time.Time) {
	unEndedEffects := make([]common.Effect, 0, len(r.effects))
	for _, e := range r.effects {
		if e.End().Before(t) {
			continue
		}

		unEndedEffects = append(unEndedEffects, e)
	}

	r.effects = unEndedEffects
}

// Render uses the stored effects from allocate(feeling)s to produce an array of lights
func (r *row) Render(t time.Time) []*addressable.Light {
	runningEffects := make([]common.Effect, 0)
	for _, f := range r.effects {
		if f.Start().After(time.Now()) {
			// since the list is ordered we can assume there are no more running effects
			break
		} else if f.End().After(time.Now()) {
			runningEffects = append(runningEffects, f)
		}
	}

	sort.Sort(byPriority(runningEffects))

	lights := r.getLights()

	for _, e := range runningEffects {
		for j := 0; j < len(lights); j++ {
			e.Render(t, lights[j])
		}
	}

	return lights
}

type byStartTime []common.Effect

func (st byStartTime) Len() int           { return len(st) }
func (st byStartTime) Swap(i, j int)      { st[i], st[j] = st[j], st[i] }
func (st byStartTime) Less(i, j int) bool { return st[i].Start().Before(st[j].Start()) }

type byPriority []common.Effect

func (p byPriority) Len() int           { return len(p) }
func (p byPriority) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byPriority) Less(i, j int) bool { return p[i].Priority() < p[j].Priority() }

func (r *row) getLights() []*addressable.Light {
	newLights := make([]*addressable.Light, len(r.lightsCache))
	for i, light := range r.lightsCache {
		newLight := *light
		newLights[i] = &newLight
	}
	return newLights
}
