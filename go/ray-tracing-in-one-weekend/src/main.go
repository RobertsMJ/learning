package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/gfx"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/hit"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/util"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"
	progressbar "github.com/schollz/progressbar/v3"
)

func init() {
	log.SetOutput(os.Stdout)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Image
	aspect_ratio := 2.0
	image_width := 500
	image_height := int(float64(image_width) / aspect_ratio)
	const samples_per_pixel = 100

	var world hit.HittableList
	world.Add(hit.Sphere{Center: vec.Point{X: 0, Y: 0, Z: -1}, Radius: 0.5})
	world.Add(hit.Sphere{Center: vec.Point{X: 2, Y: 0, Z: -1}, Radius: 0.5})
	world.Add(hit.Sphere{Center: vec.Point{X: -2, Y: 0, Z: -1}, Radius: 0.5})
	world.Add(hit.Sphere{Center: vec.Point{X: 0, Y: -100.5, Z: -1}, Radius: 100})

	// Camera
	camera := gfx.NewCamera()

	bar := progressbar.Default(int64(image_height))
	img := image.NewNRGBA(image.Rect(0, 0, image_width, image_height))

	var waitGroup sync.WaitGroup
	waitGroup.Add(image_height)

	for y := image_height - 1; y >= 0; y-- {
		go func(y int) {
			for x := 0; x < image_width; x++ {
				pixel_color := vec.Color{X: 0, Y: 0, Z: 0}
				for s := 0; s < samples_per_pixel; s++ {
					u := (float64(x) + util.Random()) / float64(image_width-1.0)
					v := (float64(y) + util.Random()) / float64(image_height-1.0)
					r := camera.GetRay(u, v)
					pixel_color = pixel_color.Add(ray_color(r, world))
				}
				img.Set(x, image_height-y, pixel_color.ToNRGBA(samples_per_pixel))
			}
			bar.Add(1)
			waitGroup.Done()
		}(y)
	}

	waitGroup.Wait()

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func ray_color(r vec.Ray, world hit.Hittable) vec.Color {
	var rec hit.HitRecord
	if world.Hit(r, 0, math.Inf(1), &rec) {
		return rec.Normal.Add(vec.Color{X: 1, Y: 1, Z: 1}).Multiply(0.5)
	}

	unit_dir := r.Direction.Unit()
	t := 0.5 * (unit_dir.Y + 1)
	return vec.Add(
		vec.Color{X: 1, Y: 1, Z: 1}.Multiply(1.0-t),
		vec.Color{X: 0.5, Y: 0.7, Z: 1}.Multiply(t))
}
