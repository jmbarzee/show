package span

import (
	"encoding/json"
	"time"

	"github.com/jmbarzee/show/common"
)

// Span represents anything that starts and Ends
type Span struct {
	start time.Time
	end   time.Time
}

// New is the primary way to make a new span
func New(start, end time.Time) *Span {
	return &Span{
		start: start,
		end:   end,
	}
}

var _ common.Spanner = (*Span)(nil)

// Start returns the Start time
func (s Span) Start() time.Time { return s.start }

// SetStart sets the Start time
func (s *Span) SetStart(start time.Time) { s.start = start }

// End returns the End time
func (s Span) End() time.Time { return s.end }

// SetEnd sets the End time
func (s *Span) SetEnd(end time.Time) { s.end = end }

// SetSpan sets the start and end of a span from the provided span
func (s *Span) SetSpan(ss common.Spanner) { s.end = ss.End(); s.start = ss.Start() }

func (s *Span) UnmarshalJSON(b []byte) error {
	ss := struct {
		Start time.Time
		End   time.Time
	}{}

	err := json.Unmarshal(b, &ss)
	if err != nil {
		return err
	}

	s.start = ss.Start
	s.end = ss.End
	return nil
}

func (s Span) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Start time.Time
		End   time.Time
	}{
		Start: s.start,
		End:   s.end,
	})
}
