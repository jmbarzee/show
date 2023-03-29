package space

import (
	"fmt"
	"math"
)

// Quaternion is quaternion with X,Y,Z and W components.
type Quaternion struct {
	W, X, Y, Z float64
}

// NewQuaternion builds a new quaternion
func NewQuaternion(w, x, y, z float64) *Quaternion {
	return &Quaternion{W: w, X: x, Y: y, Z: z}
}

// NewIdentityQuaternion returns the identity quaternion.
func NewIdentityQuaternion() *Quaternion {
	return &Quaternion{W: 1}
}

// NewRotationQuaternion returns a rotation quaternion.
// Radians should be given in terms of pi. i.e. a full rotation is 2
// If axis is the null vector, rotation will be about (1, 0, 0)
func NewRotationQuaternion(radians float64, axis Vector) *Quaternion {
	scale, w := math.Sincos(radians * math.Pi)
	axis.Normalize()
	if axis.X == 0 && axis.Y == 0 && axis.Z == 0 {
		axis.SetX(1)
	}
	axis.Scale(scale)
	return &Quaternion{W: w, X: axis.X, Y: axis.Y, Z: axis.Z}
}

// SetFromUnitVectors sets this quaternion to the rotation from vector vFrom to vTo.
// The vectors must be normalized.
// Returns pointer to this updated quaternion.
func NewPointToPointQaternion(vFrom, vTo Vector) *Quaternion {
	var v1 *Vector
	var EPS float64 = 0.000001

	r := vFrom.Dot(vTo) + 1
	if r < EPS {

		r = 0
		if math.Abs(vFrom.X) > math.Abs(vFrom.Z) {
			v1.SetXYZ(-vFrom.Y, vFrom.X, 0)
		} else {
			v1.SetXYZ(0, -vFrom.Z, vFrom.Y)
		}

	} else {

		v1 = vFrom.Cross(vTo)

	}
	q := &Quaternion{
		X: v1.X,
		Y: v1.Y,
		Z: v1.Z,
		W: r,
	}

	q.Normalize()

	return q
}

// ToVector returns the X, Y, Z vector form of the quaternion
func (q Quaternion) ToVector() *Vector {
	return &Vector{q.X, q.Y, q.Z}
}

// SetW sets the quaternion's W component.
func (q *Quaternion) SetW(w float64) { q.W = w }

// SetX sets the Vector's X component.
func (q *Quaternion) SetX(x float64) { q.X = x }

// SetY sets the Vector's Y component.
func (q *Quaternion) SetY(y float64) { q.Y = y }

// SetZ sets the Vector's Z component.
func (q *Quaternion) SetZ(z float64) { q.Z = z }

// Set sets this vector's components.
func (q *Quaternion) SetWXYZ(w, x, y, z float64) {
	q.W = w
	q.X = x
	q.Y = y
	q.Z = z
}

// Set sets this quaternion's components.
func (q *Quaternion) Set(p Quaternion) {
	q.W = p.W
	q.X = p.X
	q.Y = p.Y
	q.Z = p.Z
}

// Scale scales a Quaternion by i
func (q *Quaternion) Scale(i float64) {
	q.W *= i
	q.X *= i
	q.Y *= i
	q.Z *= i
}

// IsIdentity returns it this is an identity quaternion.
func (q Quaternion) IsIdentity() bool {
	return q.X == 0 && q.Y == 0 && q.Z == 0 && q.W == 1
}

// Inverse sets q to its inverse.
func (q *Quaternion) Inverse() {
	q.Conjugate()
	q.Normalize()
}

// Conjugate sets q to the conjugate of itself
func (q *Quaternion) Conjugate() {
	q.X *= -1
	q.Y *= -1
	q.Z *= -1
}

// Negate negates q
func (q *Quaternion) Negate() {
	q.Scale(-1)
}

// Length returns the length of this quaternion
func (q Quaternion) Length() float64 {
	return math.Sqrt(q.Dot(q))
}

// Dot returns the dot products of this quaternion with p.
func (q Quaternion) Dot(p Quaternion) float64 {
	return q.X*p.X + q.Y*p.Y + q.Z*p.Z + q.W*p.W
}

// Cross returns q cross p (qxp).
func (q Quaternion) Cross(p Quaternion) *Quaternion {
	w := q.W*p.W - q.X*p.X - q.Y*p.Y - q.Z*p.Z
	x := q.X*p.W + q.W*p.X - q.Z*p.Y + q.Y*p.Z
	y := q.Y*p.W + q.Z*p.X + q.W*p.Y - q.X*p.Z
	z := q.Z*p.W - q.Y*p.X + q.X*p.Y + q.W*p.Z
	return NewQuaternion(w, x, y, z)
}

// Normalize normalizes this quaternion.
func (q *Quaternion) Normalize() {
	l := q.Length()
	if l == 0 {
		q.SetWXYZ(1, 0, 0, 0)
	} else {
		q.Scale(1 / l)
	}
}

// Slerp returns a quaternion which is the spherically linear interpolation
// from this q to p using t.
func (q Quaternion) Slerp(p Quaternion, t float64) *Quaternion {

	if t == 0 {
		return q.Clone()
	}
	if t == 1 {
		return p.Clone()
	}

	cosHalfTheta := q.Dot(p)
	if cosHalfTheta >= 1.0 || cosHalfTheta <= -1.0 {
		return q.Clone()
	}

	ret := p.Clone()

	if cosHalfTheta < 0 {
		ret.Negate()
		cosHalfTheta = -cosHalfTheta
	}

	sqrSinHalfTheta := 1.0 - cosHalfTheta*cosHalfTheta
	if sqrSinHalfTheta < 0.001 {
		s := 1 - t
		ret.W = s*q.W + t*ret.W
		ret.X = s*q.X + t*ret.X
		ret.Y = s*q.Y + t*ret.Y
		ret.Z = s*q.Z + t*ret.Z
		ret.Normalize()
		return ret
	}

	sinHalfTheta := math.Sqrt(sqrSinHalfTheta)
	halfTheta := math.Atan2(sinHalfTheta, cosHalfTheta)
	ratioA := math.Sin((1-t)*halfTheta) / sinHalfTheta
	ratioB := math.Sin(t*halfTheta) / sinHalfTheta

	ret.W = q.W*ratioA + ret.W*ratioB
	ret.X = q.X*ratioA + ret.X*ratioB
	ret.Y = q.Y*ratioA + ret.Y*ratioB
	ret.Z = q.Z*ratioA + ret.Z*ratioB

	return ret
}

// Equals returns if this quaternion is equal to p.
func (q Quaternion) Equals(p Quaternion) bool {
	return (p.X == q.X) && (p.Y == q.Y) && (p.Z == q.Z) && (p.W == q.W)
}

// Clone returns a copy of this quaternion
func (q Quaternion) Clone() *Quaternion {
	return NewQuaternion(q.W, q.X, q.Y, q.Z)
}

func (q Quaternion) String() string {
	return fmt.Sprintf("{X:%4.2f, Y:%4.2f, Z:%4.2f, W: %4.2f}", q.X, q.Y, q.Z, q.W)
}
