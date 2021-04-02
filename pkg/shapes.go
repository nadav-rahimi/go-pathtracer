package pathtracer

import (
	"math"
)

// Records a hit
type Hit struct {
	// Time for motion blur
	T float64
	// Coordinate where the hit occurred
	Point Vec3
	// Normal to the point
	Normal Vec3
	// Material of the object hit
	Material
}

// Represents object which can be hit
type Hitable interface {
	Hit(r Ray, tMin, tMax float64) (bool, Hit)
}

// World - holds all objects in the world
type World struct {
	Elem []Hitable
}

func NewWorld() *World {
	elem := make([]Hitable, 0, 5)
	return &World{Elem: elem}
}

func (w *World) Hit(r Ray, tMin, tMax float64) (bool, Hit) {
	hitAnything := false
	closest_so_far := tMax
	rec_to_return := Hit{}

	for _, elem := range w.Elem {
		hit, rec := elem.Hit(r, tMin, closest_so_far)

		if hit {
			hitAnything = true
			closest_so_far = rec.T
			rec_to_return = rec
		}
	}

	return hitAnything, rec_to_return
}

func (w *World) Add(h Hitable) {
	w.Elem = append(w.Elem, h)
}

// Sphere
type Sphere struct {
	Centre Vec3
	Radius float64
	Material
}

func NewSphere(c Vec3, r float64, m Material) *Sphere {
	return &Sphere{
		Centre:   c,
		Radius:   r,
		Material: m,
	}
}

func (s *Sphere) Hit(r Ray, tMin, tMax float64) (bool, Hit) {
	var oc = r.Origin.Sub(s.Centre)
	var a = Dot(r.Direction, r.Direction)
	var b = Dot(oc, r.Direction)
	var c = Dot(oc, oc) - s.Radius*s.Radius
	var discrimnant = b*b - a*c

	hit := Hit{Material: s.Material}

	if discrimnant > 0 {
		var temp = (-b - math.Sqrt(discrimnant)) / a
		if temp > tMin && temp < tMax {
			hit.T = temp
			hit.Point = r.PointAtParameter(hit.T)
			hit.Normal = hit.Point.Sub(s.Centre).MakeUnitVec()
			return true, hit
		}
		temp = (-b + math.Sqrt(discrimnant)) / a
		if temp > tMin && temp < tMax {
			hit.T = temp
			hit.Point = r.PointAtParameter(temp)
			hit.Normal = hit.Point.Sub(s.Centre).MakeUnitVec()
			return true, hit
		}
	}

	return false, Hit{}
}
