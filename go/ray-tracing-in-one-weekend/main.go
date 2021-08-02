package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	progressbar "github.com/schollz/progressbar/v3"
)

func main() {
	const width, height = 100, 100
	bar := progressbar.Default(height)
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		bar.Add(1)
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	vec := Vec3{1.0, 2.0, 3.0}
	vec2 := vec.Multiply(2)
	vec3 := vec.Multiply(10).Divide(2)
	println(vec.Log())
	println(vec2.Log())
	println(vec3.Log())
	println(vec2.Add(vec3).Log())

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
