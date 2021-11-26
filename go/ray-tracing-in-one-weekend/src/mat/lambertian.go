package mat

import (
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/hit"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"
)

type Lambertian struct {
	Albedo vec.Color
}

func (l Lambertian) Scatter(rayIn vec.Ray, rec hit.HitRecord) (_ bool, scattered vec.Ray, attenuation vec.Color) {
	scatter_direction := rec.Normal.Add(vec.Random().Unit())

	if scatter_direction.NearZero() {
		scatter_direction = rec.Normal
	}

	scattered = vec.Ray{
		Origin:    rec.P,
		Direction: scatter_direction,
	}
	attenuation = l.Albedo
	return true, scattered, attenuation
}
