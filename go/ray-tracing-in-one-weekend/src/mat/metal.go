package mat

import (
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/hit"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"
)

type Metal struct {
	Albedo vec.Color
	Fuzz   float64
}

func (m Metal) Scatter(rayIn vec.Ray, rec hit.HitRecord) (_ bool, scattered vec.Ray, attenuation vec.Color) {
	reflected := vec.Reflect(rayIn.Direction.Unit(), rec.Normal)
	scattered = vec.Ray{
		Origin:    rec.P,
		Direction: reflected.Add(vec.RandomInUnitSphere().Multiply(m.Fuzz)),
	}
	attenuation = m.Albedo

	return vec.Dot(scattered.Direction, rec.Normal) > 0, scattered, attenuation
}
