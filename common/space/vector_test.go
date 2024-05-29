package space

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	VectorTest struct {
		initial   Vector
		operation VectorOperation
		expected  Vector
	}
	VectorOperation func(Vector) Vector
)

func RunVectorTests(t *testing.T, cases []VectorTest) {
	for i, c := range cases {
		actual := c.operation(c.initial)

		VectorsEqual(t, i, c.expected, actual)
	}
}

// VectorsEqual compares and diffs Vectors
func VectorsEqual(t *testing.T, testNum int, p, q Vector) bool {
	equal := FloatsEqual(p.X, q.X, MinErr) &&
		FloatsEqual(p.Y, q.Y, MinErr) &&
		FloatsEqual(p.Z, q.Z, MinErr)

	if !equal {
		assert.Equal(t, p, q, "test("+strconv.Itoa(testNum)+"): The two Quaternions should be nearly identical.")
	}
	return equal
}

func TestNewVector(t *testing.T) {
	cases := []struct {
		X, Y, Z  float64
		expected *Vector
	}{
		{
			expected: &Vector{},
		},
		{
			X:        1,
			Y:        2,
			Z:        3,
			expected: &Vector{X: 1, Y: 2, Z: 3},
		},
		{
			X:        -1,
			Y:        -2,
			Z:        -3,
			expected: &Vector{X: -1, Y: -2, Z: -3},
		},
	}

	for i, c := range cases {
		actual := NewVector(c.X, c.Y, c.Z)

		VectorsEqual(t, i, *c.expected, *actual)
	}
}

func TestVectorSetX(t *testing.T) {
	setX := func(x float64) VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			actual.SetX(x)
			return actual
		}
	}
	cases := []VectorTest{
		{
			operation: setX(0),
		},
		{
			operation: setX(1),
			expected:  Vector{X: 1},
		},
		{
			operation: setX(-1),
			expected:  Vector{X: -1},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorSetY(t *testing.T) {
	setY := func(x float64) VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			actual.SetY(x)
			return actual
		}
	}
	cases := []VectorTest{
		{
			operation: setY(0),
		},
		{
			operation: setY(1),
			expected:  Vector{Y: 1},
		},
		{
			operation: setY(-1),
			expected:  Vector{Y: -1},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorSetZ(t *testing.T) {
	setZ := func(x float64) VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			actual.SetZ(x)
			return actual
		}
	}
	cases := []VectorTest{
		{
			operation: setZ(0),
		},
		{
			operation: setZ(1),
			expected:  Vector{Z: 1},
		},
		{
			operation: setZ(-1),
			expected:  Vector{Z: -1},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorSetXYZ(t *testing.T) {
	setXYZ := func(x, y, z float64) VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			actual.SetXYZ(x, y, z)
			return actual
		}
	}
	cases := []VectorTest{
		{
			operation: setXYZ(0, 0, 0),
		},
		{
			operation: setXYZ(1, 2, 3),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},
		{
			operation: setXYZ(-1, -2, -3),
			expected:  Vector{X: -1, Y: -2, Z: -3},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorSet(t *testing.T) {
	set := func(u Vector) VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			actual.Set(u)
			return actual
		}
	}
	cases := []VectorTest{
		{
			operation: set(Vector{}),
		},
		{
			operation: set(Vector{X: 1, Y: 1, Z: 1}),
			expected:  Vector{X: 1, Y: 1, Z: 1},
		},
		{
			operation: set(Vector{X: -1, Y: -1, Z: -1}),
			expected:  Vector{X: -1, Y: -1, Z: -1},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorLength(t *testing.T) {
	cases := []struct {
		initial  *Vector
		expected float64
	}{
		{
			initial: &Vector{},
		},
		{
			initial:  &Vector{X: 1, Y: 2, Z: 3},
			expected: math.Sqrt(14),
		},
		{
			initial:  &Vector{X: -1, Y: -2, Z: -3},
			expected: math.Sqrt(14),
		},
		{
			initial:  &Vector{X: 1, Y: -2, Z: -3},
			expected: math.Sqrt(14),
		},
		{
			initial:  &Vector{X: -1, Y: 2, Z: -3},
			expected: math.Sqrt(14),
		},
		{
			initial:  &Vector{X: -1, Y: -2, Z: 3},
			expected: math.Sqrt(14),
		},
	}

	for _, c := range cases {
		actual := c.initial.Length()
		assert.Equal(t, c.expected, actual, "The two vectors should be the same.")
	}
}

func TestVectorTranslate(t *testing.T) {
	translate := func(u Vector) VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			actual.Translate(u)
			return actual
		}
	}
	cases := []VectorTest{
		{
			operation: translate(Vector{}),
		},
		{
			operation: translate(Vector{X: 1, Y: 2, Z: 3}),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},
		{
			operation: translate(Vector{X: -1, Y: -2, Z: -3}),
			expected:  Vector{X: -1, Y: -2, Z: -3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: translate(Vector{X: 1, Y: 2, Z: 3}),
			expected:  Vector{X: 2, Y: 4, Z: 6},
		},
		{
			initial:   Vector{X: -1, Y: -2, Z: -3},
			operation: translate(Vector{X: -1, Y: -2, Z: -3}),
			expected:  Vector{X: -2, Y: -4, Z: -6},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorScale(t *testing.T) {
	scale := func(i float64) VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			actual.Scale(i)
			return actual
		}
	}
	cases := []VectorTest{
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
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: scale(1),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: scale(2),
			expected:  Vector{X: 2, Y: 4, Z: 6},
		},
		{
			initial:   Vector{X: -1, Y: -2, Z: -3},
			operation: scale(1),
			expected:  Vector{X: -1, Y: -2, Z: -3},
		},
		{
			initial:   Vector{X: -1, Y: -2, Z: -3},
			operation: scale(2),
			expected:  Vector{X: -2, Y: -4, Z: -6},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: scale(-1),
			expected:  Vector{X: -1, Y: -2, Z: -3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: scale(-2),
			expected:  Vector{X: -2, Y: -4, Z: -6},
		},
		{
			initial:   Vector{X: -1, Y: -2, Z: -3},
			operation: scale(-1),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},
		{
			initial:   Vector{X: -1, Y: -2, Z: -3},
			operation: scale(-2),
			expected:  Vector{X: 2, Y: 4, Z: 6},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: scale(0),
		},
		{
			initial:   Vector{X: -1, Y: -2, Z: -3},
			operation: scale(0),
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorNormalize(t *testing.T) {
	normalize := func() VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			actual.Normalize()
			return actual
		}
	}
	cases := []VectorTest{
		{
			operation: normalize(),
		},
		{
			initial:   Vector{X: 1},
			operation: normalize(),
			expected:  Vector{X: 1},
		},
		{
			initial:   Vector{Y: 1},
			operation: normalize(),
			expected:  Vector{Y: 1},
		},
		{
			initial:   Vector{Z: 1},
			operation: normalize(),
			expected:  Vector{Z: 1},
		},
		{
			initial:   Vector{X: -1},
			operation: normalize(),
			expected:  Vector{X: -1},
		},
		{
			initial:   Vector{Y: -1},
			operation: normalize(),
			expected:  Vector{Y: -1},
		},
		{
			initial:   Vector{Z: -1},
			operation: normalize(),
			expected:  Vector{Z: -1},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: normalize(),
			expected:  Vector{X: 1 / math.Sqrt(14), Y: 2 / math.Sqrt(14), Z: 3 / math.Sqrt(14)},
		},
		{
			initial:   Vector{X: -1, Y: -2, Z: -3},
			operation: normalize(),
			expected:  Vector{X: -1 / math.Sqrt(14), Y: -2 / math.Sqrt(14), Z: -3 / math.Sqrt(14)},
		},
		{
			initial:   Vector{X: 1, Y: -2, Z: -3},
			operation: normalize(),
			expected:  Vector{X: 1 / math.Sqrt(14), Y: -2 / math.Sqrt(14), Z: -3 / math.Sqrt(14)},
		},
		{
			initial:   Vector{X: -1, Y: 2, Z: -3},
			operation: normalize(),
			expected:  Vector{X: -1 / math.Sqrt(14), Y: 2 / math.Sqrt(14), Z: -3 / math.Sqrt(14)},
		},
		{
			initial:   Vector{X: -1, Y: -2, Z: 3},
			operation: normalize(),
			expected:  Vector{X: -1 / math.Sqrt(14), Y: -2 / math.Sqrt(14), Z: 3 / math.Sqrt(14)},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorDot(t *testing.T) {
	cases := []struct {
		v        *Vector
		u        *Vector
		expected float64
	}{
		{
			v: &Vector{},
			u: &Vector{},
		},
		{
			v:        &Vector{},
			u:        &Vector{X: 1, Y: 2, Z: 3},
			expected: 0,
		},
		{
			v:        &Vector{X: 1, Y: 2, Z: 3},
			u:        &Vector{},
			expected: 0,
		},
		{
			v:        &Vector{},
			u:        &Vector{X: -1, Y: -2, Z: -3},
			expected: 0,
		},
		{
			v:        &Vector{X: 1, Y: 2, Z: 3},
			u:        &Vector{X: 1, Y: 2, Z: 3},
			expected: 14,
		},
		{
			v:        &Vector{X: -1, Y: -2, Z: -3},
			u:        &Vector{X: -1, Y: -2, Z: -3},
			expected: 14,
		},
		{
			v:        &Vector{X: 1, Y: 2, Z: 3},
			u:        &Vector{X: -1, Y: -2, Z: -3},
			expected: -14,
		},
	}

	for _, c := range cases {
		actual := c.v.Dot(*c.u)
		assert.Equal(t, c.expected, actual, "The two vectors should be the same.")
	}
}

func TestVectorCross(t *testing.T) {
	cross := func(v Vector) VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			return *actual.Cross(v)
		}
	}
	cases := []VectorTest{
		{
			operation: cross(Vector{}),
		},
		{
			initial:   Vector{X: 1},
			operation: cross(Vector{X: 1}),
		},
		{
			initial:   Vector{Y: 1},
			operation: cross(Vector{Y: 1}),
		},
		{
			initial:   Vector{Z: 1},
			operation: cross(Vector{Z: 1}),
		},
		{
			initial:   Vector{X: -1},
			operation: cross(Vector{X: -1}),
		},
		{
			initial:   Vector{Y: -1},
			operation: cross(Vector{Y: -1}),
		},
		{
			initial:   Vector{Z: -1},
			operation: cross(Vector{Z: -1}),
		},
		{
			initial:   Vector{X: -1},
			operation: cross(Vector{X: 1}),
		},
		{
			initial:   Vector{Y: -1},
			operation: cross(Vector{Y: 1}),
		},
		{
			initial:   Vector{Z: -1},
			operation: cross(Vector{Z: 1}),
		},
		{
			initial:   Vector{X: 1},
			operation: cross(Vector{X: -1}),
		},
		{
			initial:   Vector{Y: 1},
			operation: cross(Vector{Y: -1}),
		},
		{
			initial:   Vector{Z: 1},
			operation: cross(Vector{Z: -1}),
		},
		{
			initial:   Vector{X: 1},
			operation: cross(Vector{Y: 1}),
			expected:  Vector{Z: 1},
		},
		{
			initial:   Vector{X: 1},
			operation: cross(Vector{Z: 1}),
			expected:  Vector{Y: -1},
		},
		{
			initial:   Vector{X: 1},
			operation: cross(Vector{Y: -1}),
			expected:  Vector{Z: -1},
		},
		{
			initial:   Vector{X: 1},
			operation: cross(Vector{Z: -1}),
			expected:  Vector{Y: 1},
		},
		{
			initial:   Vector{Y: 1},
			operation: cross(Vector{X: 1}),
			expected:  Vector{Z: -1},
		},
		{
			initial:   Vector{Y: 1},
			operation: cross(Vector{Z: 1}),
			expected:  Vector{X: 1},
		},
		{
			initial:   Vector{Y: 1},
			operation: cross(Vector{X: -1}),
			expected:  Vector{Z: 1},
		},
		{
			initial:   Vector{Y: 1},
			operation: cross(Vector{Z: -1}),
			expected:  Vector{X: -1},
		},
		{
			initial:   Vector{Z: 1},
			operation: cross(Vector{X: 1}),
			expected:  Vector{Y: 1},
		},
		{
			initial:   Vector{Z: 1},
			operation: cross(Vector{Y: 1}),
			expected:  Vector{X: -1},
		},
		{
			initial:   Vector{Z: 1},
			operation: cross(Vector{X: -1}),
			expected:  Vector{Y: -1},
		},
		{
			initial:   Vector{Z: 1},
			operation: cross(Vector{Y: -1}),
			expected:  Vector{X: 1},
		},
		{
			initial:   Vector{X: -1},
			operation: cross(Vector{Y: 1}),
			expected:  Vector{Z: -1},
		},
		{
			initial:   Vector{X: -1},
			operation: cross(Vector{Z: 1}),
			expected:  Vector{Y: 1},
		},
		{
			initial:   Vector{X: -1},
			operation: cross(Vector{Y: -1}),
			expected:  Vector{Z: 1},
		},
		{
			initial:   Vector{X: -1},
			operation: cross(Vector{Z: -1}),
			expected:  Vector{Y: -1},
		},
		{
			initial:   Vector{Y: -1},
			operation: cross(Vector{X: 1}),
			expected:  Vector{Z: 1},
		},
		{
			initial:   Vector{Y: -1},
			operation: cross(Vector{Z: 1}),
			expected:  Vector{X: -1},
		},
		{
			initial:   Vector{Y: -1},
			operation: cross(Vector{X: -1}),
			expected:  Vector{Z: -1},
		},
		{
			initial:   Vector{Y: -1},
			operation: cross(Vector{Z: -1}),
			expected:  Vector{X: 1},
		},
		{
			initial:   Vector{Z: -1},
			operation: cross(Vector{X: 1}),
			expected:  Vector{Y: -1},
		},
		{
			initial:   Vector{Z: -1},
			operation: cross(Vector{Y: 1}),
			expected:  Vector{X: 1},
		},
		{
			initial:   Vector{Z: -1},
			operation: cross(Vector{X: -1}),
			expected:  Vector{Y: 1},
		},
		{
			initial:   Vector{Z: -1},
			operation: cross(Vector{Y: -1}),
			expected:  Vector{X: -1},
		},
		{
			initial:   Vector{X: 5, Y: 6, Z: 7},
			operation: cross(Vector{X: 2, Y: 3, Z: 4}),
			expected:  Vector{X: 3, Y: -6, Z: 3},
		},
		{
			initial:   Vector{X: -5, Y: -6, Z: -7},
			operation: cross(Vector{X: 2, Y: 3, Z: 4}),
			expected:  Vector{X: -3, Y: 6, Z: -3},
		},
		{
			initial:   Vector{X: 5, Y: 6, Z: 7},
			operation: cross(Vector{X: -2, Y: -3, Z: -4}),
			expected:  Vector{X: -3, Y: 6, Z: -3},
		},
		{
			initial:   Vector{X: 2, Y: 3, Z: 4},
			operation: cross(Vector{X: 5, Y: 6, Z: 7}),
			expected:  Vector{X: -3, Y: 6, Z: -3},
		},
		{
			initial:   Vector{X: 2, Y: 3, Z: 4},
			operation: cross(Vector{X: -5, Y: -6, Z: -7}),
			expected:  Vector{X: 3, Y: -6, Z: 3},
		},
		{
			initial:   Vector{X: -2, Y: -3, Z: -4},
			operation: cross(Vector{X: 5, Y: 6, Z: 7}),
			expected:  Vector{X: 3, Y: -6, Z: 3},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorRotate(t *testing.T) {
	rotate := func(q Quaternion) VectorOperation {
		return func(initial Vector) Vector {
			actual := initial
			actual.Rotate(q)
			return actual
		}
	}
	cases := []VectorTest{
		{
			operation: rotate(Quaternion{W: 1}),
		},

		// X-axis
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(0.0, Vector{1, 0, 0})),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(0.5, Vector{1, 0, 0})),
			expected:  Vector{X: 1, Y: -3, Z: 2},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(1.0, Vector{1, 0, 0})),
			expected:  Vector{X: 1, Y: -2, Z: -3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(1.5, Vector{1, 0, 0})),
			expected:  Vector{X: 1, Y: 3, Z: -2},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(2.0, Vector{1, 0, 0})),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},

		// Y-axis
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(0.0, Vector{0, 1, 0})),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(0.5, Vector{0, 1, 0})),
			expected:  Vector{X: 3, Y: 2, Z: -1},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(1.0, Vector{0, 1, 0})),
			expected:  Vector{X: -1, Y: 2, Z: -3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(1.5, Vector{0, 1, 0})),
			expected:  Vector{X: -3, Y: 2, Z: 1},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(2.0, Vector{0, 1, 0})),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},

		// Z-axis
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(0.0, Vector{0, 0, 1})),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(0.5, Vector{0, 0, 1})),
			expected:  Vector{X: -2, Y: 1, Z: 3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(1.0, Vector{0, 0, 1})),
			expected:  Vector{X: -1, Y: -2, Z: 3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(1.5, Vector{0, 0, 1})),
			expected:  Vector{X: 2, Y: -1, Z: 3},
		},
		{
			initial:   Vector{X: 1, Y: 2, Z: 3},
			operation: rotate(*NewRotationQuaternion(2.0, Vector{0, 0, 1})),
			expected:  Vector{X: 1, Y: 2, Z: 3},
		},
	}
	RunVectorTests(t, cases)
}
