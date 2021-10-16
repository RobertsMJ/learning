package util

import (
	"math"
	"math/rand"
)

func DegreesToRadius(n float64) float64 {
	return n * math.Pi / 180.0
}

func Random() float64 {
	return rand.Float64()
}

func RandomInRange(min, max float64) float64 {
	return min + (max-min)*Random()
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}
