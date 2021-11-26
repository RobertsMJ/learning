package hit

import "github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"

type HittableList struct {
	Objects []Hittable
}

func (hl *HittableList) Add(hittable Hittable) {
	hl.Objects = append(hl.Objects, hittable)
}

func (hl HittableList) Hit(r vec.Ray, t_min float64, t_max float64, rec *HitRecord) bool {
	var temp_rec HitRecord
	hit_anything := false
	closest_so_far := t_max

	for _, hit := range hl.Objects {
		if hit.Hit(r, t_min, closest_so_far, &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.T
			rec.P = temp_rec.P
			rec.Normal = temp_rec.Normal
			rec.T = temp_rec.T
			rec.FrontFace = temp_rec.FrontFace
			rec.Mat = temp_rec.Mat
		}
	}
	return hit_anything
}
