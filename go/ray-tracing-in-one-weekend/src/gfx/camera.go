package gfx

import (
	"math"

	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/util"
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

func NewCamera(lookfrom vec.Point, lookat vec.Point, vup vec.Vec3, vfov float64, aspect_ratio float64, aperture float64, focus_dist float64) camera {
	theta := util.DegreesToRadius(vfov)
	h := math.Tan(theta / 2.0)
	viewport_height := 2.0 * h
	viewport_width := aspect_ratio * viewport_height

	w := lookfrom.Subtract(lookat).Unit()
	u := vec.Cross(vup, w).Unit()
	v := vec.Cross(w, u)

	origin := lookfrom
	horizontal := u.Multiply(viewport_width)
	vertical := v.Multiply(viewport_height)
	lower_left_corner := origin.Subtract(horizontal.Divide(2)).Subtract(vertical.Divide(2)).Subtract(w)

	return camera{
		Origin:          origin,
		Horizontal:      u.Multiply(viewport_width),
		Vertical:        v.Multiply(viewport_height),
		LowerLeftCorner: lower_left_corner,
	}
}

func (c camera) GetRay(s, t float64) vec.Ray {
	dir := c.LowerLeftCorner.
		Add(c.Horizontal.Multiply(s)).
		Add(c.Vertical.Multiply(t)).
		Subtract(c.Origin)

	return vec.Ray{Origin: c.Origin, Direction: dir}
}
