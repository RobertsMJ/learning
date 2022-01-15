package vec

import (
	"fmt"
	"math"

	"github.com/robertsmj1/learning/go/ray-tracing-in-one-weekend/src/util"
)

type Vec3 struct {
	X, Y, Z float64
}

type Point = Vec3

func (v Vec3) ToString() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}

func (v Vec3) Negate() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vec3) Subtract(v2 Vec3) Vec3 {
	return Vec3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v Vec3) Multiply(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v Vec3) MultiplyVec(v2 Vec3) Vec3 {
	return Vec3{X: v.X * v2.X, Y: v.Y * v2.Y, Z: v.Z * v2.Z}
}

func (v Vec3) Divide(t float64) Vec3 {
	return v.Multiply(1.0 / t)
}

func (v1 Vec3) Dot(v2 Vec3) float64 {
	return v1.X*v2.X +
		v1.Y*v2.Y +
		v1.Z*v2.Z
}

func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

func (v Vec3) Unit() Vec3 {
	return v.Divide(v.Length())
}

func (v Vec3) NearZero() bool {
	s := 1e-8
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}

// TODO:MJR is this better?
func Add(v1 Vec3, v2 Vec3) Vec3 {
	return v1.Add(v2)
}

func Subtract(v1 Vec3, v2 Vec3) Vec3 {
	return v1.Subtract(v2)
}

func Multiply(v Vec3, scalar float64) Vec3 {
	return v.Multiply(scalar)
}

func MultiplyVec(v1 Vec3, v2 Vec3) Vec3 {
	return v1.MultiplyVec(v2)
}

func Divide(v Vec3, scalar float64) Vec3 {
	return v.Divide(scalar)
}

func Dot(v1 Vec3, v2 Vec3) float64 {
	return v1.Dot(v2)
}

func Cross(v1 Vec3, v2 Vec3) Vec3 {
	return v1.Cross(v2)
}

func Reflect(v Vec3, n Vec3) Vec3 {
	return v.Subtract(n.Multiply(Dot(v, n) * 2))
}

func Refract(uv Vec3, n Vec3, etai_over_etat float64) Vec3 {
	cos_theta := math.Min(Dot(uv.Negate(), n), 1.0)
	r_out_perp := uv.Add(n.Multiply(cos_theta)).Multiply(etai_over_etat)
	r_out_parallel := n.Multiply(-math.Sqrt(math.Abs(1.0 - r_out_perp.LengthSquared())))
	return r_out_perp.Add(r_out_parallel)
}

func Random() Vec3 {
	return Vec3{X: util.Random(), Y: util.Random(), Z: util.Random()}
}

func RandomInRange(min, max float64) Vec3 {
	return Vec3{X: util.RandomInRange(min, max), Y: util.RandomInRange(min, max), Z: util.RandomInRange(min, max)}
}

func RandomInUnitSphere() Vec3 {
	var p Vec3
	for p = Random(); p.LengthSquared() >= 1; p = Random() {
		return p
	}
	return p
}

func RandomInUnitDisk() Vec3 {
	for {
		p := Vec3{X: util.RandomInRange(-1, 1), Y: util.RandomInRange(-1, 1), Z: 0}
		if p.LengthSquared() < 1 {
			return p
		}
	}
}

func RandomUnitVector() Vec3 {
	return RandomInUnitSphere().Unit()
}

func RandomInHemisphere(normal Vec3) Vec3 {
	inUnitSphere := RandomInUnitSphere()
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	}
	return inUnitSphere.Negate()
}
