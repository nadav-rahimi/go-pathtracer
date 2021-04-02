package pathtracer

import (
	"math"
	"math/rand"
)

// Materials define an interface which return the ray bounced
// off the object and the colour of the material
type Material interface {
	Bounce(r Ray, h Hit, rnd *rand.Rand) (bool, Ray)
	Colour() Colour
}

// Returns a random vector in a unit sphere
func VectorInUnitSphere(rnd *rand.Rand) Vec3 {
	for {
		temp := Vec3{rnd.Float64(), rnd.Float64(), rnd.Float64()}
		p := temp.MulFloat(2.0).Sub(UnitVector)
		if p.SquaredLen() >= 1.0 {
			return p
		}
	}
}

// Matte objects use Lambertian reflectance
type Matte struct {
	Attenuation Colour
}

func (m Matte) Bounce(input_ray Ray, h Hit, rnd *rand.Rand) (bool, Ray) {
	direction := h.Normal.Add(VectorInUnitSphere(rnd))
	return true, Ray{h.Point, direction}
}

func (m Matte) Colour() Colour {
	return m.Attenuation
}

// Metal objects
type Metal struct {
	Attenuation Colour
	Fuzz        float64
}

func (m Metal) Bounce(input_ray Ray, h Hit, rnd *rand.Rand) (bool, Ray) {
	direction := m.reflect(input_ray.Direction, h.Normal)
	bounced := Dot(direction, h.Normal) > 0
	return bounced, Ray{h.Point, direction.Add(VectorInUnitSphere(rnd).MulFloat(m.Fuzz))}
}

func (m Metal) Colour() Colour {
	return m.Attenuation
}

func (m Metal) reflect(v, n Vec3) Vec3 {
	return v.Sub(n.MulFloat(2 * Dot(v, n)))
}

// Dielectrics (e.g. Glass)
type Dielectric struct {
	RefractiveIndex float64
	Attenuation     Colour
}

const (
	Air     = 1.0
	Glass   = 1.5
	Diamond = 2.4
)

func (d Dielectric) Bounce(input_ray Ray, h Hit, rnd *rand.Rand) (bool, Ray) {
	var outwardNormal Vec3
	var niOverNt, cosine, reflectProb float64

	bounced := Dot(input_ray.Direction, h.Normal) > 0
	if bounced {
		outwardNormal = h.Normal.MulFloat(-1.0)
		niOverNt = d.RefractiveIndex
		cosine = d.RefractiveIndex * Dot(input_ray.Direction, h.Normal) / input_ray.Direction.Len()
	} else {
		outwardNormal = h.Normal
		niOverNt = 1.0 / d.RefractiveIndex
		cosine = -1 * d.RefractiveIndex * Dot(input_ray.Direction, h.Normal) / input_ray.Direction.Len()
	}

	refracted, refraction := d.refract(input_ray.Direction, outwardNormal, niOverNt)
	if refracted {
		reflectProb = d.schlick(cosine)
	} else {
		reflectProb = 1.0
	}

	if rnd.Float64() < reflectProb {
		reflection := d.reflect(input_ray.Direction, h.Normal)
		return true, NewRay(h.Point, reflection)
	}

	return true, NewRay(h.Point, refraction)
}

func (d Dielectric) Colour() Colour {
	return d.Attenuation
}

func (d Dielectric) refract(v, n Vec3, niOverNt float64) (bool, Vec3) {
	uv := v.MakeUnitVec()
	un := n.MakeUnitVec()
	dt := Dot(uv, un)
	discrimnant := 1.0 - (niOverNt * niOverNt * (1 - dt*dt))
	if discrimnant > 0 {
		refraction := (uv.Sub(un.MulFloat(dt))).MulFloat(niOverNt).Sub(un.MulFloat(math.Sqrt(discrimnant)))
		return true, refraction
	} else {
		return false, Vec3{}
	}
}

func (d Dielectric) reflect(v, n Vec3) Vec3 {
	return v.Sub(n.MulFloat(2 * Dot(v, n)))
}

func (d Dielectric) schlick(cosine float64) float64 {
	r0 := (1 - d.RefractiveIndex) / (1 + d.RefractiveIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
