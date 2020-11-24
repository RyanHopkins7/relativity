package newton

import "math"

// Vector2D : a 2d vector containing magnitude and direction
type Vector2D struct {
	magnitude float64
	direction float64
}

// NewVector2D : construct a Vector2D with magnitude and direction
func NewVector2D(magnitude float64, direction float64) *Vector2D {
	v := new(Vector2D)
	v.magnitude = magnitude
	v.direction = direction

	return v
}

// Add : add vec2 to vec1
func (vec1 Vector2D) Add(vec2 Vector2D) {
	vec1.magnitude += vec2.magnitude
	vec1.direction += vec2.direction
}

// XComponent : get x component of vec1
func (vec1 Vector2D) XComponent() float64 {
	return vec1.magnitude * math.Cos(vec1.direction)
}

// yComponent : get y component of vec1
func (vec1 Vector2D) yComponent() float64 {
	return vec1.magnitude * math.Sin(vec1.direction)
}
