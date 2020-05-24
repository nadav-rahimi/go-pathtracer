package pathtracer

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Returns the colour for the given ray in the world using
// c1 and c2 as the gradient for the background colours
func ColourScene(r Ray, w *World, depth int, rnd *rand.Rand, c1, c2 Colour) Colour {
	hit, rec := w.Hit(r, 0.001, math.MaxFloat64)
	backgroundColour := Gradient(c1, c2, r.Direction.MakeUnitVec().Y)

	if hit {
		bounced, bouncedRay := rec.Bounce(r, rec, rnd)

		if depth < 50 { // Depth is the max number of possible bounces
			if bounced {
				newColour := ColourScene(bouncedRay, w, depth+1, rnd, c1, c2)
				if newColour == backgroundColour {
					return rec.Material.Colour()
				}
				return rec.Material.Colour().Mul(newColour)

			}
		}
		return Black
	}


	return backgroundColour
}

// Adapted from: https://github.com/markphelps/go-trace
// Super samples the given pixel at ns times.
func SuperSample(nx, ny, ns, i, j int, rnd *rand.Rand, cam *Camera, w *World, c1, c2 Colour) Colour {
	c := Black
	for s := 0; s < ns; s++ {
		u := (float64(i) + rnd.Float64()) / float64(nx)
		v := (float64(j) + rnd.Float64()) / float64(ny)
		ray := cam.RayAt(u, v)
		c = c.Add(ColourScene(ray, w, 0, rnd, c1, c2))
	}

	c = c.DivFloat(float64(ns))
	c = Colour{math.Sqrt(c.R), math.Sqrt(c.G), math.Sqrt(c.B)} // Gamma Correction

	return c
}

// Adapted from: https://github.com/markphelps/go-trace
// Renders the given image of width nx, and height ny. Samples each pixel ns time.
// Uses the viewpoint from cam and elements in the world "w". Uses colours c1 and c2
// to generate the background gradient
func Render(nx, ny, ns int, cam *Camera, w *World, c1, c2 Colour) image.Image{
	img := image.NewNRGBA(image.Rect(0, 0, nx, ny))
	numCpus := runtime.NumCPU()
	var wg sync.WaitGroup

	for i := 0; i < numCpus; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

			for row := i; row < ny; row += numCpus {
				for col := 0; col < nx; col++ {
					c := SuperSample(nx, ny, ns, col, row, rnd, cam, w, c1, c2)
					ir := c.R256()
					ig := c.G256()
					ib := c.B256()

					img.Set(col, ny-row-1, color.RGBA{ir, ig, ib, 0xff})
				}
			}
		}(i)
	}

	wg.Wait()
	return img
}