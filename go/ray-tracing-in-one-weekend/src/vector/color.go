package vector

import "image/color"

type Color = Vec3

func (c Color) R() float64 {
	return c.X
}

func (c Color) G() float64 {
	return c.Y
}

func (c Color) B() float64 {
	return c.Z
}

func (c Color) ToNRGBA() color.NRGBA {
	return color.NRGBA{
		R: uint8(c.R() * 255.99),
		G: uint8(c.G() * 255.99),
		B: uint8(c.B() * 255.99),
		A: 255,
	}
}
