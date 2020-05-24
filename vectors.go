package pathtracer

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

// Default methods no pointers
func (v1 Vec3) Add(v2 Vec3) Vec3 {
	v1.X = v1.X + v2.X
	v1.Y = v1.Y + v2.Y
	v1.Z = v1.Z + v2.Z

	return v1
}

func (v1 Vec3) Sub(v2 Vec3) Vec3{
	v1.X = v1.X - v2.X
	v1.Y = v1.Y - v2.Y
	v1.Z = v1.Z - v2.Z

	return v1
}

func (v1 Vec3) Mul(v2 Vec3) Vec3{
	v1.X = v1.X * v2.X
	v1.Y = v1.Y * v2.Y
	v1.Z = v1.Z * v2.Z

	return v1
}

func (v1 Vec3) MulFloat(f float64) Vec3{
	v1.X = v1.X * f
	v1.Y = v1.Y * f
	v1.Z = v1.Z * f

	return v1
}

func (v1 Vec3) Div(v2 Vec3) Vec3{
	v1.X = v1.X / v2.X
	v1.Y = v1.Y / v2.Y
	v1.Z = v1.Z / v2.Z

	return v1
}

func (v1 Vec3) DivFloat(f float64) Vec3{
	v1.X = v1.X / f
	v1.Y = v1.Y / f
	v1.Z = v1.Z / f

	return v1
}

func (v1 Vec3) Cross(v2 Vec3) Vec3{
	return Vec3{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: -1 * (v1.X*v2.Z - v1.Z*v2.X),
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}

func Dot(v1, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
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
var UnitVector = Vec3{X: 1, Y: 1, Z: 1,}
