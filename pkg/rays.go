package pathtracer

// Ray of "light"
type Ray struct {
	Origin, Direction Vec3
}

// Creates a Ray
func NewRay(o, d Vec3) Ray {
	return Ray{o, d}
}

// Point at "time" T calculated by Point = O + T*D
func (r *Ray) PointAtParameter(t float64) Vec3 {
	return r.Origin.Add(r.Direction.MulFloat(t))
}
