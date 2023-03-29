package space

import (
	"fmt"
	"math"
)

// Vector represents a point in space in Vector coordinates
type Vector struct {
	X, Y, Z float64
}

// NewVector produces a new Vector from spherical coordinates
func NewVector(x, y, z float64) *Vector {
	return &Vector{X: x, Y: y, Z: z}
}

func (v Vector) ToQuaternion() *Quaternion {
	return &Quaternion{0, v.X, v.Y, v.Z}
}

// SetX sets the Vector's X component.
func (v *Vector) SetX(x float64) { v.X = x }

// SetY sets the Vector's Y component.
func (v *Vector) SetY(y float64) { v.Y = y }

// SetZ sets the Vector's Z component.
func (v *Vector) SetZ(z float64) { v.Z = z }

// Set sets this vector's components.
func (v *Vector) SetXYZ(x, y, z float64) {
	v.X = x
	v.Y = y
	v.Z = z
}

// Set sets this vector's to the given vector's components.
func (v *Vector) Set(u Vector) {
	v.X = u.X
	v.Y = u.Y
	v.Z = u.Z
}

// Length returns the length of this Vector
func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

// Translate shifts a Vector by a Vector (addition)
func (v *Vector) Translate(u Vector) {
	v.X = v.X + u.X
	v.Y = v.Y + u.Y
	v.Z = v.Z + u.Z
}

// Scale scales a Vector by i
func (v *Vector) Scale(i float64) {
	v.X = v.X * i
	v.Y = v.Y * i
	v.Z = v.Z * i
}

// Normalize normalizes this Vector.
// The zero vector will not be altered
func (v *Vector) Normalize() {
	l := v.Length()
	if l == 0 {
		v.SetXYZ(0, 0, 0)
	} else {
		v.Scale(1 / l)
	}
}

// Dot returns the dot products of this Vector with u.
func (v Vector) Dot(u Vector) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

// Cross returns v cross u (vxu).
func (v Vector) Cross(u Vector) *Vector {
	x := v.Y*u.Z - v.Z*u.Y
	y := v.Z*u.X - v.X*u.Z
	z := v.X*u.Y - v.Y*u.X
	return NewVector(x, y, z)
}

// Rotate spins v around the origin based on q
func (v *Vector) Rotate(q Quaternion) {
	p := v.ToQuaternion()
	qi := q.Clone()
	qi.Conjugate()

	qp := q.Cross(*p)
	qpqi := qp.Cross(*qi)
	v.Set(*qpqi.ToVector())
}

// Clone returns a copy of this Vector
func (v Vector) Clone() *Vector {
	return NewVector(v.X, v.Y, v.Z)
}

func (c Vector) String() string {
	return fmt.Sprintf("{X:%4.2f, Y:%4.2f, Z:%4.2f}", c.X, c.Y, c.Z)
}
