package span

import (
	"time"

	"github.com/jmbarzee/show/common"
)

// Span represents anything that starts and Ends
type Span struct {
	StartTime time.Time
	EndTime   time.Time
}

var _ common.Spanner = (*Span)(nil)

// Start returns the Start time
func (s Span) Start() time.Time { return s.StartTime }

// End returns the End time
func (s Span) End() time.Time { return s.EndTime }

// Seed is represents anything that starts and Ends
type Seed struct {
	Span

	count int // incremented to change seed
}

var _ common.Seeder = (*Seed)(nil)

func NewSeed(start, end time.Time) *Seed {
	return &Seed{
		count: 0,
		Span: Span{
			StartTime: start,
			EndTime:   end,
		},
	}
}

func (s *Seed) NextSeed() time.Time {
	s.count++
	return s.Start().Add(time.Second * time.Duration(s.count))
}
