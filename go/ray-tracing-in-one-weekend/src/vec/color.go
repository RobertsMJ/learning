package vec

import (
	"image/color"
	"math"

	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/util"
)

type Color = Vec3

func (c Color) ToNRGBA(samples_per_pixel int) color.NRGBA {
	r := c.X
	g := c.Y
	b := c.Z

	scale := 1.0 / float64(samples_per_pixel)
	r = math.Sqrt(scale * r)
	g = math.Sqrt(scale * g)
	b = math.Sqrt(scale * b)

	return color.NRGBA{
		R: uint8(256 * util.Clamp(r, 0.0, 0.999)),
		G: uint8(256 * util.Clamp(g, 0.0, 0.999)),
		B: uint8(256 * util.Clamp(b, 0.0, 0.999)),
		A: 255,
	}
}
