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
	LensRadius      float64
	u, v, w         vec.Vec3
}

func NewCamera(
	lookfrom vec.Point,
	lookat vec.Point,
	vup vec.Vec3,
	vfov float64,
	aspect_ratio float64,
	aperture float64,
	focus_dist float64) camera {

	theta := util.DegreesToRadius(vfov)
	h := math.Tan(theta / 2.0)
	viewport_height := 2.0 * h
	viewport_width := aspect_ratio * viewport_height

	w := lookfrom.Subtract(lookat).Unit()
	u := vec.Cross(vup, w).Unit()
	v := vec.Cross(w, u)

	origin := lookfrom
	horizontal := u.Multiply(viewport_width).Multiply(focus_dist)
	vertical := v.Multiply(viewport_height).Multiply(focus_dist)
	lower_left_corner := origin.
		Subtract(horizontal.Divide(2)).
		Subtract(vertical.Divide(2)).
		Subtract(w.Multiply(focus_dist))

	return camera{
		Origin:          origin,
		Horizontal:      horizontal,
		Vertical:        vertical,
		LowerLeftCorner: lower_left_corner,
		LensRadius:      aperture / 2.0,
		w:               w,
		u:               u,
		v:               v,
	}
}

func (c camera) GetRay(s, t float64) vec.Ray {
	rd := vec.RandomInUnitDisk().Multiply(c.LensRadius)
	offset := vec.Add(c.u.Multiply(rd.X), c.v.Multiply(rd.Y))

	dir := c.LowerLeftCorner.
		Add(c.Horizontal.Multiply(s)).
		Add(c.Vertical.Multiply(t)).
		Subtract(c.Origin).
		Subtract(offset)

	return vec.Ray{Origin: c.Origin.Add(offset), Direction: dir}
}
