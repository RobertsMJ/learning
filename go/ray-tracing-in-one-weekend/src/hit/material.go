package hit

import (
	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/vec"
)

type Material interface {
	Scatter(rayIn vec.Ray, rec HitRecord) (bool, vec.Ray, vec.Color)
}
