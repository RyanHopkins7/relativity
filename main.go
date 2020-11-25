package main

import (
	"math"
	"relativity/newton"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Two body gravitational simulation",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	body1InitialVelocity := *newton.NewVector2D(1.0, math.Pi/2)
	body1 := *newton.NewGravitationalBody(100.0, 0.0, 1.0e5, -10.0, body1InitialVelocity)

	body2InitialVelocity := *newton.NewVector2D(0.0, 0.0)
	body2 := *newton.NewGravitationalBody(5.972e24, 0.0, 0.0, 0.0, body2InitialVelocity)

	for !win.Closed() {
		body1.Update(body1.Gravity(body2))
		// body2.Update(body2.Gravity(body1))

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
