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

	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/geo"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/gfx"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/hit"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/mat"
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
	const max_depth = 50

	var world hit.HittableList

	material_ground := mat.Lambertian{Albedo: vec.Color{X: 0.8, Y: 0.8, Z: 0.0}}
	material_center := mat.Lambertian{Albedo: vec.Color{X: 0.7, Y: 0.3, Z: 0.3}}
	material_left := mat.Metal{Albedo: vec.Color{X: 0.8, Y: 0.8, Z: 0.0}, Fuzz: 0.3}
	material_right := mat.Metal{Albedo: vec.Color{X: 0.8, Y: 0.6, Z: 0.2}, Fuzz: 1.0}

	world.Add(geo.Sphere{Center: vec.Point{X: 0, Y: -100.5, Z: -1}, Radius: 100, Mat: material_ground})
	world.Add(geo.Sphere{Center: vec.Point{X: 0, Y: 0, Z: -1}, Radius: 0.5, Mat: material_center})
	world.Add(geo.Sphere{Center: vec.Point{X: -1, Y: 0, Z: -1}, Radius: 0.5, Mat: material_left})
	world.Add(geo.Sphere{Center: vec.Point{X: 1, Y: 0, Z: -1}, Radius: 0.5, Mat: material_right})

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
					pixel_color = pixel_color.Add(ray_color(r, world, max_depth))
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

func ray_color(r vec.Ray, world hit.Hittable, depth int) vec.Color {
	if depth <= 0 {
		return vec.Color{}
	}

	var rec hit.HitRecord
	if world.Hit(r, 0.001, math.Inf(1), &rec) {
		scatter, scattered, attenuation := rec.Mat.Scatter(r, rec)
		if scatter {
			return attenuation.MultiplyVec(ray_color(scattered, world, depth-1))
		}
		return vec.Color{}
	}

	unit_dir := r.Direction.Unit()
	t := 0.5 * (unit_dir.Y + 1)
	return vec.Add(
		vec.Color{X: 1, Y: 1, Z: 1}.Multiply(1.0-t),
		vec.Color{X: 0.5, Y: 0.7, Z: 1}.Multiply(t))
}
