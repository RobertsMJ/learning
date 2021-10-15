package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"sync"

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

	// Camera
	viewport_height := 2.0
	viewport_width := aspect_ratio * viewport_height
	focal_length := 1.0

	origin := Point{X: 0, Y: 0, Z: 0}
	horizontal := Vec3{X: viewport_width, Y: 0, Z: 0}
	vertical := Vec3{X: 0, Y: viewport_height, Z: 0}
	lower_left_corner := origin.
		Subtract(horizontal.Divide(2)).
		Subtract(vertical.Divide(2)).
		Subtract(Vec3{X: 0, Y: 0, Z: focal_length})

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
				ray := Ray{Origin: origin, Direction: direction}
				pixel := ray_color(ray)
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

func hit_sphere(center Point, radius float64, r Ray) float64 {
	oc := r.Origin.Subtract(center)

	a := r.Direction.LengthSquared()
	half_b := oc.Dot(r.Direction)
	c := oc.LengthSquared() - radius*radius
	discriminant := half_b*half_b - a*c

	if discriminant < 0 {
		return -1.0
	} else {
		return (-half_b - math.Sqrt(discriminant)) / a
	}
}

func ray_color(r Ray) Color {
	t := hit_sphere(Point{X: 0, Y: 0, Z: -1}, 0.5, r)
	if t > 0.0 {
		n := r.At(t).Subtract(Vec3{X: 0, Y: 0, Z: -1})
		return Color{X: n.X + 1, Y: n.Y + 1, Z: n.Z + 1}.Multiply(0.5)
	}

	unit_dir := r.Direction.Unit()
	t = 0.5 * (unit_dir.Y + 1.0)
	return Add(
		Color{X: 1, Y: 1, Z: 1}.Multiply(1.0-t),
		Color{X: 0.5, Y: 0.7, Z: 1}.Multiply(t))
}
