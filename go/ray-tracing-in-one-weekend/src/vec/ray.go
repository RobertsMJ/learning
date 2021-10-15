package vec

import "fmt"

type Ray struct {
	Origin    Point
	Direction Vec3
}

func (r Ray) Log() string {
	return fmt.Sprintf("Origin: %s, Direction: %s", r.Origin.ToString(), r.Direction.ToString())
}

func (r Ray) At(t float64) Point {
	return r.Origin.Add(r.Direction.Multiply(t))
}
