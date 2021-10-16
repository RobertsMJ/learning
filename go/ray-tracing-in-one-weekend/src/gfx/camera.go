package gfx

import (
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"
)

type Camera struct {
	*camera
}

type camera struct {
	Origin          vec.Point
	LowerLeftCorner vec.Point
	Horizontal      vec.Vec3
	Vertical        vec.Vec3
}

func NewCamera() camera {
	aspect_ratio := 2.0
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

	return camera{
		Origin:          origin,
		Horizontal:      horizontal,
		Vertical:        vertical,
		LowerLeftCorner: lower_left_corner,
	}
}

func (c camera) GetRay(u, v float64) vec.Ray {
	dir := c.LowerLeftCorner.
		Add(c.Horizontal.Multiply(u)).
		Add(c.Vertical.Multiply(v)).
		Subtract(c.Origin)

	return vec.Ray{Origin: c.Origin, Direction: dir}
}
