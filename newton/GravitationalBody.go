// Package newton implements a two body Newtonian gravitational simulation using pixel
package newton

import (
	"fmt"
	"math"
)

// G : the Gravitational Constant in m3kg-1s2
const G = 6.674e-11

// GravitationalBody : a circular body which interacts gravitationally with other massive bodies
type GravitationalBody struct {
	mass         float64  // mass in kg
	radius       float64  // radius in m (purely graphical)
	xPosition    float64  // x position in m
	yPosition    float64  // y position in m
	velocity     Vector2D // velocity vector of form (speed, angle) where speed in m/s and angle in radians
	acceleration Vector2D // acceleration vector of form (acceleration, angle) where acceleration in m/s2 and angle in radians
}

// NewGravitationalBody : construct a GravitationalBody
func NewGravitationalBody(mass float64, radius float64, xPosition float64, yPosition float64, velocity Vector2D) *GravitationalBody {
	body1 := new(GravitationalBody)
	body1.mass = mass
	body1.radius = radius
	body1.xPosition = xPosition
	body1.yPosition = yPosition
	body1.velocity = velocity

	return body1
}

// Distance : calculate distance between two gravitational bodies
func (body1 *GravitationalBody) Distance(body2 GravitationalBody) float64 {
	return math.Sqrt(math.Pow(body2.xPosition-body1.xPosition, 2) + math.Pow(body2.yPosition-body1.yPosition, 2))
}

// Gravity : calculate gravitational force vector between two gravitational bodies
func (body1 *GravitationalBody) Gravity(body2 GravitationalBody) Vector2D {
	r := body1.Distance(body2)
	force := G * (body1.mass * body2.mass) / math.Pow(r, 2)
	angle := math.Atan(math.Abs(body1.yPosition-body2.yPosition) / math.Abs(body1.xPosition-body2.xPosition))

	// Reflect angle to appropriate quadrant
	if body1.xPosition-body2.xPosition > 0 && body1.yPosition-body2.yPosition < 0 {
		angle = math.Pi - angle
	} else if body1.xPosition-body2.xPosition > 0 && body1.yPosition-body2.yPosition > 0 {
		angle = math.Pi + angle
	} else if body1.xPosition-body2.xPosition < 0 && body1.yPosition-body2.yPosition > 0 {
		angle = math.Pi*2 - angle
	}

	return *NewVector2D(force, angle)
}

// Update : update xPosition, yPosition, velocity, and acceleration given netForce acting on body1
func (body1 *GravitationalBody) Update(netForce Vector2D) {
	body1.acceleration = *NewVector2D(netForce.magnitude/body1.mass, netForce.direction)

	// dt is 1 frame
	body1.velocity.Add(body1.acceleration)

	fmt.Print("X position: ")
	fmt.Print(body1.xPosition)
	fmt.Print(" Y position: ")
	fmt.Print(body1.yPosition)
	fmt.Println()

	body1.xPosition += body1.velocity.XComponent()
	body1.yPosition += body1.velocity.YComponent()
}
