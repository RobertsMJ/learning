package vector

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

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

func Divide(v Vec3, scalar float64) Vec3 {
	return v.Divide(scalar)
}

func Dot(v1 Vec3, v2 Vec3) float64 {
	return v1.Dot(v2)
}

func Cross(v1 Vec3, v2 Vec3) Vec3 {
	return v1.Cross(v2)
}
