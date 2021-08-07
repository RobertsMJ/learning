package vector

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
