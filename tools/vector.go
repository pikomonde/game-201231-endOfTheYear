package tools

import (
	"math"
)

type Vector2 struct {
	X, Y float64
}

func (v *Vector2) Rotate(deg float64) {
	v.X, v.Y = v.X*math.Cos(deg)-v.Y*math.Sin(deg), v.X*math.Sin(deg)+v.Y*math.Cos(deg)
}
