package main

import (
	"fmt"
	"math"
	"relativity/newton"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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

	// Initialize gravitational bodies
	body1InitialVelocity := *newton.NewVector2D(0.5, math.Pi)
	body1 := *newton.NewGravitationalBody(1.0e12, 10.0, 50.0, 50.0, body1InitialVelocity)

	body2InitialVelocity := *newton.NewVector2D(0.5, 0.0)
	body2 := *newton.NewGravitationalBody(1.0e12, 10.0, -50.0, -50.0, body2InitialVelocity)

	// Initialize sprites
	body1Sprite := imdraw.New(nil)
	body2Sprite := imdraw.New(nil)
	body1Sprite.Color = pixel.RGB(0, 0, 0)
	body2Sprite.Color = pixel.RGB(0, 0, 0)

	// Center sprites at origin
	body1Sprite.SetMatrix(pixel.IM.Moved(cfg.Bounds.Center()))
	body2Sprite.SetMatrix(pixel.IM.Moved(cfg.Bounds.Center()))

	// Measure FPS
	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {
		body1.Update(body1.Gravity(body2))
		body2.Update(body2.Gravity(body1))

		win.Clear(colornames.White)

		// Draw sprites
		body1Sprite.Clear()
		body1Sprite.Push(pixel.V(body1.GetXPosition(), body1.GetYPosition()))
		body1Sprite.Circle(body1.GetRadius(), 2)
		body1Sprite.Draw(win)

		body2Sprite.Clear()
		body2Sprite.Push(pixel.V(body2.GetXPosition(), body2.GetYPosition()))
		body2Sprite.Circle(body2.GetRadius(), 2)
		body2Sprite.Draw(win)

		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
