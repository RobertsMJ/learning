package hit

import (
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"
)

type Hittable interface {
	Hit(r vec.Ray, t_min float64, t_max float64, rec *HitRecord) bool
}

type HitRecord struct {
	P         vec.Point
	Normal    vec.Vec3
	T         float64
	FrontFace bool
	Mat       Material
}

func (hr *HitRecord) SetFaceNormal(r vec.Ray, outward_normal vec.Vec3) {
	hr.FrontFace = r.Direction.Dot(outward_normal) < 0
	if hr.FrontFace {
		hr.Normal = outward_normal
	} else {
		hr.Normal = outward_normal.Negate()
	}
}
