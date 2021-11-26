package geo

import (
	"math"

	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/hit"
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"
)

type Sphere struct {
	Center vec.Point
	Radius float64
	Mat    hit.Material
}

func (s Sphere) Hit(r vec.Ray, t_min float64, t_max float64, rec *hit.HitRecord) bool {
	oc := r.Origin.Subtract(s.Center)

	a := r.Direction.LengthSquared()
	half_b := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := half_b*half_b - a*c

	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range
	root := (-half_b - sqrtd) / a
	if root < t_min || root > t_max {
		root = (-half_b + sqrtd) / a

		if root < t_min || root > t_max {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	outward_normal := rec.P.Subtract(s.Center).Divide(s.Radius)
	rec.SetFaceNormal(r, outward_normal)
	rec.Mat = s.Mat

	return true
}
