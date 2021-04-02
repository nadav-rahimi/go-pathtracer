package main

import (
	"flag"
	"fmt"
	pt "github.com/fiwippi/go-pathtracer/pkg"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func main() {
	nx := flag.Int("width", 800, "Width of each example image")
	ny := flag.Int("height", 800, "Height of each example image")
	ns := flag.Int("samples", 100, "How many times to render each pixel")
	flag.Parse()

	fmt.Printf("Running with Width=%v Height=%v Samples=%v\n", *nx, *ny, *ns)

	// Example 1
	f, _ := os.Create("ex_1.png")
	err := png.Encode(f, Example1(*nx, *ny, *ns))
	if err != nil {
		log.Println(err)
	}

	// Example 2
	f, _ = os.Create("ex_2.png")
	err = png.Encode(f, Example2(*nx, *ny, *ns))
	if err != nil {
		log.Println(err)
	}

	// Example 3
	f, _ = os.Create("ex_3.png")
	err = png.Encode(f, Example3(*nx, *ny, *ns))
	if err != nil {
		log.Println(err)
	}
}

func Example1(nx, ny, ns int) image.Image {
	w := pt.NewWorld()
	w.Add(pt.NewSphere(pt.Vec3{0, 0, -1}, 0.5, pt.Metal{pt.Orange, 0}))
	w.Add(pt.NewSphere(pt.Vec3{0.5 * 0.05 * 50, 0, -1}, 0.5, pt.Metal{pt.Yellow, 0}))
	w.Add(pt.NewSphere(pt.Vec3{0, -1000, 0}, 999.5, pt.Metal{pt.SortaBlue, 0}))
	w.Add(pt.NewSphere(pt.Vec3{-0.5 * 0.05 * 50, 0, -1}, 0.5, pt.Metal{pt.Red, 0}))

	cam := pt.NewCamera(pt.Vec3{0, 0, 1}, pt.Vec3{0, 0, -1}, 90, float64(nx)/float64(ny))

	return pt.Render(nx, ny, ns, cam, w, pt.OffWhite, pt.Salmon)
}

func Example2(nx, ny, ns int) image.Image {
	w := pt.NewWorld()
	w.Add(pt.NewSphere(pt.Vec3{0, -1000, 0}, 999, pt.Metal{pt.SortaBlue, 0}))

	height := 0.06363636363
	w.Add(pt.NewSphere(pt.Vec3{0, float64(height), -1}, 1, pt.Dielectric{1.5, pt.White}))
	w.Add(pt.NewSphere(pt.Vec3{1.05, float64(height), -2.5}, 1, pt.Dielectric{1.5, pt.LightRed}))
	w.Add(pt.NewSphere(pt.Vec3{-1.05, float64(height), -2.5}, 1, pt.Dielectric{1.5, pt.SpotifyGreen}))

	cam := pt.NewCamera(pt.Vec3{0, 0, 3}, pt.Vec3{0, height, -1}, 80, (float64(nx) / float64(ny)))

	return pt.Render(nx, ny, ns, cam, w, pt.Orange, pt.DarkBlue)
}

func Example3(nx, ny, ns int) image.Image {
	w := pt.NewWorld()
	w.Add(pt.NewSphere(pt.Vec3{0, -1000, 0}, 1000, pt.Matte{pt.Grey}))
	for a := -42; a < 7; a++ {
		for b := -42; b < 11; b++ {
			chooseMat := rand.Float64()
			centre := pt.Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if centre.Sub(pt.Vec3{4, 0.2, 0}).Len() > 0.9 {
				if chooseMat < 0.6 {
					w.Add(pt.NewSphere(centre, rand.Float64()*0.25+0.05, pt.Matte{pt.Colour{rand.Float64(), rand.Float64(), rand.Float64()}}))
				} else if chooseMat < 0.95 {
					w.Add(pt.NewSphere(centre, rand.Float64()*0.25+0.05, pt.Metal{pt.Colour{rand.Float64(), rand.Float64(), rand.Float64()}, rand.Float64() * 0.5}))
				} else {
					w.Add(pt.NewSphere(centre, rand.Float64()*0.25+0.05, pt.Dielectric{pt.Glass, pt.White}))
				}
			}

		}
	}
	w.Add(pt.NewSphere(pt.Vec3{-25, 7, -15}, 7, pt.Dielectric{pt.Glass, pt.White}))
	w.Add(pt.NewSphere(pt.Vec3{-2, 1, -4.2}, 1, pt.Metal{pt.Colour{0.7, 0.6, 0.5}, 0}))
	w.Add(pt.NewSphere(pt.Vec3{-2 - 10, 1, -4.2 + 5}, 1, pt.Metal{pt.Colour{0.7, 0.6, 0.5}, 0}))
	w.Add(pt.NewSphere(pt.Vec3{-3, 1, -4.2 - 10}, 1, pt.Metal{pt.Colour{0.7, 0.6, 0.5}, 0}))
	cam := pt.NewCamera(pt.Vec3{13, 5, 3}, pt.Vec3{-25, 12, -15}, 70.35, (float64(nx) / float64(ny)))
	return pt.Render(nx, ny, ns, cam, w, pt.LightGreen, pt.LightPink)
}
