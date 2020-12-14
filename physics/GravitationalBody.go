// Package newton implements a two body Newtonian gravitational simulation using pixel
package physics

import (
	"math"
)

// G : the Gravitational Constant in m3kg-1s2
const G = 6.674e-11

// GravitationalBody : a circular body which interacts gravitationally with other massive bodies
type GravitationalBody struct {
	Mass         float64  // mass in kg
	Radius       float64  // radius in m (purely graphical)
	XPosition    float64  // x position in m
	YPosition    float64  // y position in m
	Velocity     Vector2D // velocity vector of form (speed, angle) where speed in m/s and angle in radians
	Acceleration Vector2D // acceleration vector of form (acceleration, angle) where acceleration in m/s2 and angle in radians
}

// NewGravitationalBody : construct a GravitationalBody
func NewGravitationalBody(mass float64, radius float64, xPosition float64, yPosition float64, velocity Vector2D) *GravitationalBody {
	body1 := new(GravitationalBody)
	body1.Mass = mass
	body1.Radius = radius
	body1.XPosition = xPosition
	body1.YPosition = yPosition
	body1.Velocity = velocity

	return body1
}

// Distance : calculate distance between two gravitational bodies
func (body1 *GravitationalBody) Distance(body2 GravitationalBody) float64 {
	return math.Sqrt(math.Pow(body2.XPosition-body1.XPosition, 2) + math.Pow(body2.YPosition-body1.YPosition, 2))
}

// Gravity : calculate gravitational force vector between two gravitational bodies
func (body1 *GravitationalBody) Gravity(body2 GravitationalBody) Vector2D {
	r := body1.Distance(body2)
	force := G * (body1.Mass * body2.Mass) / math.Pow(r, 2)
	angle := math.Atan(math.Abs(body1.YPosition-body2.YPosition) / math.Abs(body1.XPosition-body2.XPosition))

	// Reflect angle to appropriate quadrant
	if body1.XPosition-body2.XPosition > 0 && body1.YPosition-body2.YPosition < 0 {
		angle = math.Pi - angle
	} else if body1.XPosition-body2.XPosition > 0 && body1.YPosition-body2.YPosition >= 0 {
		angle = math.Pi + angle
	} else if body1.XPosition-body2.XPosition < 0 && body1.YPosition-body2.YPosition > 0 {
		angle = math.Pi*2 - angle
	}

	return *NewVector2D(force, angle)
}

// Update : update xPosition, yPosition, velocity, and acceleration given netForce acting on body1
func (body1 *GravitationalBody) Update(netForce Vector2D) {
	body1.Acceleration = *NewVector2D(netForce.magnitude/body1.Mass, netForce.direction)

	// dt is 1 frame
	body1.Velocity.Add(body1.Acceleration)

	body1.XPosition += body1.Velocity.XComponent()
	body1.YPosition += body1.Velocity.YComponent()
}
