package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"sync"

	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/hit"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"
	progressbar "github.com/schollz/progressbar/v3"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	// Image
	aspect_ratio := 2.0
	image_width := 500
	image_height := int(float64(image_width) / aspect_ratio)

	var world hit.HittableList
	world.Objects = append(world.Objects, hit.Sphere{Center: vec.Point{X: 0, Y: 0, Z: -1}, Radius: 0.5})
	world.Objects = append(world.Objects, hit.Sphere{Center: vec.Point{X: 0, Y: -100.5, Z: -1}, Radius: 100})

	// Camera
	viewport_height := 2.0
	viewport_width := aspect_ratio * viewport_height
	focal_length := 1.0

	origin := vec.Point{X: 0, Y: 0, Z: 0}
	horizontal := vec.Vec3{X: viewport_width, Y: 0, Z: 0}
	vertical := vec.Vec3{X: 0, Y: viewport_height, Z: 0}
	lower_left_corner := origin.
		Subtract(horizontal.Divide(2)).
		Subtract(vertical.Divide(2)).
		Subtract(vec.Vec3{X: 0, Y: 0, Z: focal_length})

	bar := progressbar.Default(int64(image_height))
	img := image.NewNRGBA(image.Rect(0, 0, image_width, image_height))

	var waitGroup sync.WaitGroup
	waitGroup.Add(image_height)

	for y := image_height - 1; y >= 0; y-- {
		go func(y int) {
			for x := 0; x < image_width; x++ {
				u := float64(x) / (float64(image_width))
				v := float64(y) / (float64(image_height))
				direction := lower_left_corner.
					Add(horizontal.Multiply(u)).
					Add(vertical.Multiply(v)).
					Subtract(origin)
				ray := vec.Ray{Origin: origin, Direction: direction}
				pixel := ray_color(ray, world)
				img.Set(x, image_height-y, pixel.ToNRGBA())
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
