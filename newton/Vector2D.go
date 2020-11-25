package newton

import (
	"math"
)

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

// XComponent : get x component of vec1
func (vec1 *Vector2D) XComponent() float64 {
	return vec1.magnitude * math.Cos(vec1.direction)
}

// YComponent : get y component of vec1
func (vec1 *Vector2D) YComponent() float64 {
	return vec1.magnitude * math.Sin(vec1.direction)
}

// Add : add vec2 to vec1
func (vec1 *Vector2D) Add(vec2 Vector2D) {
	newXComponent := vec1.XComponent() + vec2.XComponent()
	newYComponent := vec1.YComponent() + vec2.YComponent()

	vec1.magnitude = math.Sqrt(newXComponent*newXComponent + newYComponent*newYComponent)
	vec1.direction = math.Atan(math.Abs(newYComponent) / math.Abs(newXComponent))

	// Reflect angle to appropriate quadrant
	if newXComponent < 0 && newYComponent > 0 {
		vec1.direction = math.Pi - vec1.direction
	} else if newXComponent < 0 && newYComponent < 0 {
		vec1.direction = math.Pi + vec1.direction
	} else if newXComponent > 0 && newYComponent < 0 {
		vec1.direction = math.Pi*2 - vec1.direction
	}
}
