package space

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewObject(t *testing.T) {
	cases := []struct {
		location    Vector
		orientation Quaternion
		expected    *Object
	}{
		{
			expected: &Object{},
		},
	}

	for _, c := range cases {
		actual := NewObject(c.location, c.orientation)
		assert.Equal(t, actual, c.expected, "The two Objects should be the same.")
	}
}
