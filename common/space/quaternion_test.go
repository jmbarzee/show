package space

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	QuaternionTest struct {
		initial   Quaternion
		operation QuaternionOperation
		expected  Quaternion
	}
	QuaternionOperation func(Quaternion) Quaternion
)

func RunQuaternionTests(t *testing.T, cases []QuaternionTest) {
	for i, c := range cases {
		actual := c.operation(c.initial)

		assert.Equal(t, c.expected, actual, "test("+strconv.Itoa(i)+"): The two vectors should be the same.")
	}
}

func TestNewQuaternion(t *testing.T) {
	cases := []struct {
		W, X, Y, Z float64
		expected   *Quaternion
	}{
		{
			expected: &Quaternion{},
		},
		{
			W: 1,
			X: 2,
			Y: 3,
			Z: 4,
			expected: &Quaternion{
				W: 1,
				X: 2,
				Y: 3,
				Z: 4,
			},
		},
		{
			W: -1,
			X: -2,
			Y: -3,
			Z: -4,
			expected: &Quaternion{
				W: -1,
				X: -2,
				Y: -3,
				Z: -4,
			},
		},
	}

	for _, c := range cases {
		actual := NewQuaternion(c.W, c.X, c.Y, c.Z)
		assert.Equal(t, c.expected, actual, "The two quaternions should be the same.")
	}
}
func TestNewIdentityQuaternion(t *testing.T) {
	cases := []struct {
		expected *Quaternion
	}{
		{
			expected: &Quaternion{
				W: 1,
			},
		},
	}

	for _, c := range cases {
		actual := NewIdentityQuaternion()
		assert.Equal(t, c.expected, actual, "The two quaternions should be the same.")
	}
}

func TestNewRotationQuaternion(t *testing.T) {
	cases := []struct {
		radians  float64
		axis     Vector
		expected *Quaternion
	}{
		// {
		// 	expected: NewIdentityQuaternion(),
		// },
		// {
		// 	radians:  0.5,
		// 	expected: &Quaternion{W: 0, X: 1, Y: 0, Z: 0},
		// },
	}

	for _, c := range cases {
		actual := NewRotationQuaternion(c.radians, c.axis)
		assert.Equal(t, c.expected, actual, "The two quaternions should be the same.")
	}
}

func TestNewPointToPointQaternion(t *testing.T) {
	cases := []struct {
		v, u     Vector
		expected *Quaternion
	}{
		// {
		// 	expected: &Quaternion{},
		// },
		// {
		// 	expected: &Quaternion{W: 1, X: 2, Y: 3, Z: 4},
		// },
	}

	for _, c := range cases {
		actual := NewPointToPointQaternion(c.v, c.u)
		assert.Equal(t, c.expected, actual, "The two quaternions should be the same.")
	}
}

func TestQuaternionSetW(t *testing.T) {
	setW := func(w float64) QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.SetW(w)
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: setW(0),
		},
		{
			operation: setW(1),
			expected:  Quaternion{W: 1},
		},
		{
			operation: setW(-1),
			expected:  Quaternion{W: -1},
		},
	}
	RunQuaternionTests(t, cases)
}

func TestQuaternionSetX(t *testing.T) {
	setX := func(x float64) QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.SetX(x)
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: setX(0),
		},
		{
			operation: setX(1),
			expected:  Quaternion{X: 1},
		},
		{
			operation: setX(-1),
			expected:  Quaternion{X: -1},
		},
	}
	RunQuaternionTests(t, cases)
}

func TestQuaternionSetY(t *testing.T) {
	setY := func(y float64) QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.SetY(y)
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: setY(0),
		},
		{
			operation: setY(1),
			expected:  Quaternion{Y: 1},
		},
		{
			operation: setY(-1),
			expected:  Quaternion{Y: -1},
		},
	}
	RunQuaternionTests(t, cases)
}

func TestQuaternionSetZ(t *testing.T) {
	setZ := func(z float64) QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.SetZ(z)
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: setZ(0),
		},
		{
			operation: setZ(1),
			expected:  Quaternion{Z: 1},
		},
		{
			operation: setZ(-1),
			expected:  Quaternion{Z: -1},
		},
	}
	RunQuaternionTests(t, cases)
}

func TestQuaternionSetWXYZ(t *testing.T) {
	setWXYZ := func(w, x, y, z float64) QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.SetWXYZ(w, x, y, z)
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: setWXYZ(0, 0, 0, 0),
		},
		{
			operation: setWXYZ(1, 2, 3, 4),
			expected:  Quaternion{W: 1, X: 2, Y: 3, Z: 4},
		},
		{
			operation: setWXYZ(-1, -2, -3, -4),
			expected:  Quaternion{W: -1, X: -2, Y: -3, Z: -4},
		},
	}
	RunQuaternionTests(t, cases)
}

func TestQuaternionSet(t *testing.T) {
	set := func(q Quaternion) QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.Set(q)
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: set(Quaternion{}),
		},
		{
			operation: set(Quaternion{1, 2, 3, 4}),
			expected:  Quaternion{W: 1, X: 2, Y: 3, Z: 4},
		},
		{
			operation: set(Quaternion{-1, -2, -3, -4}),
			expected:  Quaternion{W: -1, X: -2, Y: -3, Z: -4},
		},
	}
	RunQuaternionTests(t, cases)
}

func TestQuaternionScale(t *testing.T) {
	scale := func(i float64) QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.Scale(i)
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: scale(0),
		},
		{
			operation: scale(1),
		},
		{
			operation: scale(2),
		},
		{
			operation: scale(-1),
		},
		{
			operation: scale(-2),
		},
		{
			initial:   Quaternion{W: 1, X: 2, Y: 3, Z: 4},
			operation: scale(1),
			expected:  Quaternion{W: 1, X: 2, Y: 3, Z: 4},
		},
		{
			initial:   Quaternion{W: 1, X: 2, Y: 3, Z: 4},
			operation: scale(2),
			expected:  Quaternion{W: 2, X: 4, Y: 6, Z: 8},
		},
		{
			initial:   Quaternion{W: -1, X: -2, Y: -3, Z: -4},
			operation: scale(1),
			expected:  Quaternion{W: -1, X: -2, Y: -3, Z: -4},
		},
		{
			initial:   Quaternion{W: -1, X: -2, Y: -3, Z: -4},
			operation: scale(2),
			expected:  Quaternion{W: -2, X: -4, Y: -6, Z: -8},
		},
		{
			initial:   Quaternion{W: 1, X: 2, Y: 3, Z: 4},
			operation: scale(-1),
			expected:  Quaternion{W: -1, X: -2, Y: -3, Z: -4},
		},
		{
			initial:   Quaternion{W: 1, X: 2, Y: 3, Z: 4},
			operation: scale(-2),
			expected:  Quaternion{W: -2, X: -4, Y: -6, Z: -8},
		},
		{
			initial:   Quaternion{W: -1, X: -2, Y: -3, Z: -4},
			operation: scale(-1),
			expected:  Quaternion{W: 1, X: 2, Y: 3, Z: 4},
		},
		{
			initial:   Quaternion{W: -1, X: -2, Y: -3, Z: -4},
			operation: scale(-2),
			expected:  Quaternion{W: 2, X: 4, Y: 6, Z: 8},
		},
		{
			initial:   Quaternion{W: 1, X: 2, Y: 3, Z: 4},
			operation: scale(0),
		},
		{
			initial:   Quaternion{W: -1, X: -2, Y: -3, Z: -4},
			operation: scale(0),
		},
	}
	RunQuaternionTests(t, cases)
}

func TestQuaternionIsIdentity(t *testing.T) {
	cases := []struct {
		initial  Quaternion
		expected bool
	}{
		{
			// Zero value quaternion should not return as an identity
		},
		{
			initial:  Quaternion{W: 1},
			expected: true,
		},
		{
			initial: Quaternion{W: -1},
		},
		{
			initial: Quaternion{X: 2, Y: 3, Z: 4},
		},
		{
			initial: Quaternion{X: -2, Y: -3, Z: -4},
		},
		{
			initial: Quaternion{W: 1, X: 2, Y: 3, Z: 4},
		},
		{
			initial: Quaternion{W: -1, X: -2, Y: -3, Z: -4},
		},
	}

	for _, c := range cases {
		actual := c.initial.IsIdentity()
		assert.Equal(t, c.expected, actual, "The two booleans should be the same.")
	}
}

func TestQuaterionInverse(t *testing.T) {
	inverse := func() QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.Inverse()
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: inverse(),
			expected:  Quaternion{W: 1},
		},
	}
	RunQuaternionTests(t, cases)
}

func TestQuaterionConjugate(t *testing.T) {
	conjugate := func() QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.Conjugate()
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: conjugate(),
		},
	}
	RunQuaternionTests(t, cases)

}

func TestQuaterionNegate(t *testing.T) {
	negate := func() QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.Negate()
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: negate(),
		},
	}
	RunQuaternionTests(t, cases)

}

func TestQuaternionLength(t *testing.T) {
	cases := []struct {
		initial  Quaternion
		expected float64
	}{
		{
			// The Zero Quaternion should have zero length
		},
		{
			initial:  Quaternion{W: 1},
			expected: 1,
		},
		{
			initial:  Quaternion{X: 2},
			expected: 2,
		},
		{
			initial:  Quaternion{Y: 3},
			expected: 3,
		},
		{
			initial:  Quaternion{Z: 4},
			expected: 4,
		},
	}

	for _, c := range cases {
		actual := c.initial.Length()
		assert.Equal(t, c.expected, actual, "The two floats should be the same.")
	}
}

func TestQuaternionDot(t *testing.T) {
	cases := []struct {
		q        Quaternion
		p        Quaternion
		expected float64
	}{
		{
			// The Zero Quaternions should have zero dot result
		},
		{
			q:        Quaternion{W: 2},
			p:        Quaternion{W: 3},
			expected: 6,
		},
		{
			q:        Quaternion{X: 2},
			p:        Quaternion{X: 3},
			expected: 6,
		},
		{
			q:        Quaternion{Y: 2},
			p:        Quaternion{Y: 3},
			expected: 6,
		},
		{
			q:        Quaternion{Z: 2},
			p:        Quaternion{Z: 3},
			expected: 6,
		},
	}

	for _, c := range cases {
		actual := c.p.Dot(c.q)
		assert.Equal(t, c.expected, actual, "The two floats should be the same.")
	}
}

func TestQuaterionCross(t *testing.T) {
	cross := func(q Quaternion) QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.Cross(q)
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: cross(Quaternion{}),
		},
	}
	RunQuaternionTests(t, cases)

}

func TestQuaterionNormalize(t *testing.T) {
	normalize := func() QuaternionOperation {
		return func(initial Quaternion) Quaternion {
			actual := initial
			actual.Normalize()
			return actual
		}
	}
	cases := []QuaternionTest{
		{
			operation: normalize(),
			expected:  Quaternion{W: 1},
		},
		{
			initial:   Quaternion{W: 1},
			operation: normalize(),
			expected:  Quaternion{W: 1},
		},
		{
			initial:   Quaternion{X: 1},
			operation: normalize(),
			expected:  Quaternion{X: 1},
		},
		{
			initial:   Quaternion{Y: 1},
			operation: normalize(),
			expected:  Quaternion{Y: 1},
		},
		{
			initial:   Quaternion{Z: 1},
			operation: normalize(),
			expected:  Quaternion{Z: 1},
		},
	}
	RunQuaternionTests(t, cases)

}
