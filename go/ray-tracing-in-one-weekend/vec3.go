package main

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}
type Color Vec3
type Point Vec3

func (v *Vec3) Log() string {
	return fmt.Sprintf("%f %f %f", v.X, v.Y, v.Z)
}

func (v *Vec3) Negate() *Vec3 {
	return &Vec3{-v.X, -v.Y, -v.Z}
}

func (v1 *Vec3) Add(v2 *Vec3) *Vec3 {
	return &Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 *Vec3) Subtract(v2 *Vec3) *Vec3 {
	return &Vec3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v *Vec3) Multiply(t float64) *Vec3 {
	return &Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v *Vec3) Divide(t float64) *Vec3 {
	return v.Multiply(1 / t)
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) LengthSquared() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}
