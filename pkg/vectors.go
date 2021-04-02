package pathtracer

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

// Default methods no pointers
func (v Vec3) Add(u Vec3) Vec3 {
	v.X = v.X + u.X
	v.Y = v.Y + u.Y
	v.Z = v.Z + u.Z

	return v
}

func (v Vec3) Sub(u Vec3) Vec3 {
	v.X = v.X - u.X
	v.Y = v.Y - u.Y
	v.Z = v.Z - u.Z

	return v
}

func (v Vec3) Mul(u Vec3) Vec3 {
	v.X = v.X * u.X
	v.Y = v.Y * u.Y
	v.Z = v.Z * u.Z

	return v
}

func (v Vec3) MulFloat(f float64) Vec3 {
	v.X = v.X * f
	v.Y = v.Y * f
	v.Z = v.Z * f

	return v
}

func (v Vec3) Div(u Vec3) Vec3 {
	v.X = v.X / u.X
	v.Y = v.Y / u.Y
	v.Z = v.Z / u.Z

	return v
}

func (v Vec3) DivFloat(f float64) Vec3 {
	v.X = v.X / f
	v.Y = v.Y / f
	v.Z = v.Z / f

	return v
}

func (v Vec3) Cross(u Vec3) Vec3 {
	return Vec3{
		X: v.Y*u.Z - v.Z*u.Y,
		Y: -1 * (v.X*u.Z - v.Z*u.X),
		Z: v.X*u.Y - v.Y*u.X,
	}
}

func Dot(v, u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

// Miscellaneous
func (v Vec3) At(i int) float64 {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	default:
		panic("Vector index out of range")
	}
}

func (v Vec3) Len() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func (v Vec3) SquaredLen() float64 {
	return math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2)
}

func (v Vec3) MakeUnitVec() Vec3 {
	return v.DivFloat(v.Len())
}

func (v Vec3) String() string {
	return fmt.Sprintf("(%v, %v, %v)", v.X, v.Y, v.Z)
}

// Specific Vectors
var UnitVector = Vec3{X: 1, Y: 1, Z: 1}
