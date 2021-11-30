package mat

import (
	"math"

	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/hit"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/util"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"
)

type Dielectric struct {
	IR float64 // Index of Refraction
}

func (d Dielectric) Scatter(rayIn vec.Ray, rec hit.HitRecord) (_ bool, scattered vec.Ray, attenuation vec.Color) {
	attenuation = vec.Color{X: 1, Y: 1, Z: 1}
	var refraction_ratio float64
	if rec.FrontFace {
		refraction_ratio = 1.0 / d.IR
	} else {
		refraction_ratio = d.IR
	}
	unit_direction := rayIn.Direction.Unit()

	cos_theta := math.Min(vec.Dot(unit_direction.Negate(), rec.Normal), 1.0)
	sin_theta := math.Sqrt(1.0 - cos_theta*cos_theta)
	cannot_refract := refraction_ratio*sin_theta > 1.0

	var direction vec.Vec3
	if cannot_refract || reflectance(cos_theta, refraction_ratio) > util.Random() {
		direction = vec.Reflect(unit_direction, rec.Normal)
	} else {
		direction = vec.Refract(unit_direction, rec.Normal, refraction_ratio)
	}

	scattered = vec.Ray{
		Origin:    rec.P,
		Direction: direction,
	}
	return true, scattered, attenuation
}

func reflectance(cosine float64, ref_idx float64) float64 {
	r0 := math.Pow((1-ref_idx)/(1+ref_idx), 2)
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
