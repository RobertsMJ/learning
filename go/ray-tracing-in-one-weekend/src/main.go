package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
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
	aspect_ratio := 16.0 / 9.0
	image_width, _ := strconv.Atoi(getEnv("WIDTH", "500"))
	image_height := int(float64(image_width) / aspect_ratio)
	const samples_per_pixel = 500
	const max_depth = 50

	world := random_scene()

	// Camera
	lookfrom := vec.Point{X: 13, Y: 2, Z: 3}
	lookat := vec.Point{X: 0, Y: 0, Z: 0}
	vup := vec.Vec3{X: 0, Y: 1, Z: 0}
	// dist_to_focus := lookfrom.Subtract(lookat).Length()
	dist_to_focus := 10.0
	vfov := 20.0
	aperture := 0.1
	camera := gfx.NewCamera(lookfrom, lookat, vup, vfov, aspect_ratio, aperture, dist_to_focus)

	// Render
	bar := progressbar.Default(int64(image_height * image_width))
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
				bar.Add(1)
			}
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

func random_scene() hit.HittableList {
	var world hit.HittableList

	ground_material := mat.Lambertian{Albedo: vec.Color{X: 0.5, Y: 0.5, Z: 0.5}}
	world.Add(geo.Sphere{Center: vec.Point{X: 0, Y: -1000, Z: 0}, Radius: 1000, Mat: ground_material})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			center := vec.Point{X: float64(a) + 0.9*util.Random(), Y: 0.2, Z: float64(b) + 0.9*util.Random()}
			mat_type := util.Random()
			switch {
			case mat_type < 0.8:
				albedo := vec.Random().MultiplyVec(vec.Random())
				mat := mat.Lambertian{Albedo: albedo}
				world.Add(geo.Sphere{Center: center, Radius: 0.2, Mat: mat})
			case mat_type < 0.95:
				albedo := vec.RandomInRange(0.5, 1)
				fuzz := util.RandomInRange(0, 0.5)
				mat := mat.Metal{Albedo: albedo, Fuzz: fuzz}
				world.Add(geo.Sphere{Center: center, Radius: 0.2, Mat: mat})
			default:
				mat := mat.Dielectric{IR: 1.5}
				world.Add(geo.Sphere{Center: center, Radius: 0.2, Mat: mat})
			}
		}
	}

	mat_lambertian := mat.Lambertian{Albedo: vec.Color{X: 0.4, Y: 0.2, Z: 0.1}}
	mat_dielectric := mat.Dielectric{IR: 1.5}
	mat_metal := mat.Metal{Albedo: vec.Color{X: 0.7, Y: 0.6, Z: 0.5}, Fuzz: 0.0}

	world.Add(geo.Sphere{Center: vec.Point{X: -4, Y: 1, Z: 0}, Radius: 1, Mat: mat_lambertian})
	world.Add(geo.Sphere{Center: vec.Point{X: 0, Y: 1, Z: 0}, Radius: 1, Mat: mat_dielectric})
	world.Add(geo.Sphere{Center: vec.Point{X: 4, Y: 1, Z: 0}, Radius: 1, Mat: mat_metal})

	return world
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
