package repeat

import (
	"hash/fnv"
	"math"
	"time"

	"github.com/jmbarzee/show/common"
)

// Seed is represents anything that starts and Ends
type Seed struct {
	seed  time.Time
	count int // incremented to change seed
}

var _ common.Seeder = (*Seed)(nil)

// NewSeed creates a new seed and specifies the seed
// Used to replay Past Patterns
func NewSeed(seed time.Time) *Seed {
	return &Seed{
		count: 0,
		seed:  seed,
	}
}

// NextSeed provides a unique time for seeding repeatable, random operations
func (s *Seed) NextSeed() time.Time {
	s.count++
	return s.seed.Add(time.Millisecond * time.Duration(s.count))
}

// Option is an idempotent call based on a time and a number of options
// If the result of a fast hash of the time is within a certain
// option-range of possible hash values, it will return that option
func (s *Seed) Option(options int) int {
	return option(s.NextSeed(), options)
}

// Chance is an idempotent call based on a time and a prospect in the range (1, 0)
// it will return true if the result of a fast hash of the time is above
// a certain percentage (chance) of possible hash values
func (s *Seed) Chance(prospect float64) bool {
	return chance(s.NextSeed(), prospect)
}

// RandDuration is an idempotent call based on a time and a range of durations
// it will return a random duration from within that range
func (s *Seed) RandDuration(min, max time.Duration) time.Duration {
	return randDuration(s.NextSeed(), min, max)
}

// RandShift is an idempotent call based on a time and a range of shifts
// it will return a random shift from within that range
// unit provides the granularity of the randomness
func (s *Seed) RandShift(min, max, unit float64) float64 {
	return randShift(s.NextSeed(), min, max, unit)
}

// option is an idempotent call based on a time and a number of options
// If the result of a fast hash of the time is within a certain
// option-range of possible hash values, it will return that option
func option(t time.Time, options int) int {
	sum := float32(hashSum(t))
	max := float32(math.MaxUint32) // maximum possible hash
	portion := max / float32(options)

	return int(sum / portion)
}

// chance is an idempotent call based on a time and a chance in the range (1, 0)
// it will return true if the result of a fast hash of the time is above
// a certain percentage (chance) of possible hash values
func chance(t time.Time, chance float64) bool {
	sum := float64(hashSum(t))
	max := float64(math.MaxUint32) // maximum possible hash

	return sum <= max*chance
}

// randDuration is an idempotent call based on a time and a range of durations
// it will return a random duration from within that range
func randDuration(t time.Time, min, max time.Duration) time.Duration {
	sum := time.Duration(hashSum(t))
	diff := max - min
	return sum%(diff) + min
}

// randShift is an idempotent call based on a time and a range of shifts
// it will return a random shift from within that range
// unit provides the granularity of the randomness
func randShift(t time.Time, min, max, unit float64) float64 {
	diff := max - min
	units := int(diff / unit)
	// fmt.Println("units:", units)
	option := option(t, units)
	// fmt.Println("Option:", option)
	return float64(option)*unit + min
}

func hashSum(t time.Time) uint32 {
	hash := fnv.New32()
	timeBinary, err := t.MarshalBinary()
	if err != nil {
		panic(err)
	}

	_, err = hash.Write(timeBinary)
	if err != nil {
		panic(err)
	}

	return hash.Sum32()
}
