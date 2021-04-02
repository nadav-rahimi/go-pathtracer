package pathtracer

import (
	"math"
)

// The up orientation defined for the camera
var vUp = Vec3{0, 1, 0}

// Camera object which emits the rays
type Camera struct {
	Origin          Vec3 // Vector where the camera looks from
	LowerLeftCorner Vec3 // Vector defined as point (0, 0) of the camera's screen
	Horizontal      Vec3 // Horizontal width of the screen
	Vertical        Vec3 // Vertical height of the screen
}

// Generates a new camera object. The VFOV should be specified in degrees
func NewCamera(lookfrom, lookat Vec3, VFOV, aspectR float64) *Camera {
	theta := VFOV * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspectR * halfHeight

	// The orthogonal vectors of the plane
	w := lookfrom.Sub(lookat).MakeUnitVec()
	u := vUp.Cross(w).MakeUnitVec()
	v := w.Cross(u)

	return &Camera{
		Origin:          lookfrom,
		LowerLeftCorner: lookfrom.Sub(u.MulFloat(halfWidth)).Sub(v.MulFloat(halfHeight)).Sub(w),
		Horizontal:      u.MulFloat(halfWidth * 2.0),
		Vertical:        v.MulFloat(halfHeight * 2.0),
	}
}

// Returns the ray pointing at the specific x and y offset
func (c *Camera) RayAt(s, t float64) Ray {
	x_offset := c.Horizontal.MulFloat(s)
	y_offset := c.Vertical.MulFloat(t)
	direction := c.LowerLeftCorner.Add(x_offset).Add(y_offset).Sub(c.Origin)

	return NewRay(c.Origin, direction)
}

//func NewCamera(VFOV, aspectR float64) *Camera { // takes in degrees but converts to radians
//	theta := VFOV * math.Pi / 180
//	halfHeight := math.Tan(theta/2)
//	halfWidth := aspectR * halfHeight
//
//	return &Camera{
//		Origin:            ZeroVec(),
//		LowerLeftCorner: Vec3{-halfWidth, -halfHeight, -1.0},
//		Horizontal:        Vec3{2*halfWidth, 0, 0},
//		Vertical:          Vec3{0, 2*halfHeight, 0},
//	}
//}
//
//func (c *Camera) RayAt(u, v float64) Ray {
//	x_offset := c.Horizontal.MulFloat(u)
//	y_offset := c.Vertical.MulFloat(v)
//	direction := c.LowerLeftCorner.Add(x_offset).Add(y_offset).Sub(c.Origin)
//
//	return NewRay(c.Origin, direction)
//}
