package main

import (
	"fmt"
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

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	body1InitialVelocity := *newton.NewVector2D(0.0, 0.0)
	body1 := *newton.NewGravitationalBody(0.0, 0.0, 0.0, 0.0, body1InitialVelocity)

	fmt.Print(body1)

	pixelgl.Run(run)
}
