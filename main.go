package main

import (
	"fmt"
	"gravity-simulation/physics"
	"math"
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
	body1InitialVelocity := *physics.NewVector2D(5.6, -math.Pi/12)
	body1 := *physics.NewGravitationalBody(1.0e12, 10.0, 0.0, 200.0, body1InitialVelocity)

	body2InitialVelocity := *physics.NewVector2D(1.0, math.Pi)
	body2 := *physics.NewGravitationalBody(2.0e13, 20.0, 0.0, -100.0, body2InitialVelocity)

	body3InitialVelocity := *physics.NewVector2D(2.5, 0)
	body3 := *physics.NewGravitationalBody(1.0e13, 15.0, 0.0, 100.0, body3InitialVelocity)

	// Initialize sprites
	body1Sprite := imdraw.New(nil)
	body2Sprite := imdraw.New(nil)
	body3Sprite := imdraw.New(nil)
	body1Sprite.Color = pixel.RGB(0, 0, 0)
	body2Sprite.Color = pixel.RGB(0, 0, 0)
	body3Sprite.Color = pixel.RGB(0, 0, 0)

	// Center sprites at origin
	body1Sprite.SetMatrix(pixel.IM.Moved(cfg.Bounds.Center()))
	body2Sprite.SetMatrix(pixel.IM.Moved(cfg.Bounds.Center()))
	body3Sprite.SetMatrix(pixel.IM.Moved(cfg.Bounds.Center()))

	// Measure FPS
	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {
		body1Gravity := body1.Gravity(body2)
		body1Gravity.Add(body1.Gravity(body3))
		body1.Update(body1Gravity)

		body2Gravity := body2.Gravity(body1)
		body2Gravity.Add(body2.Gravity(body3))
		body2.Update(body2Gravity)

		body3Gravity := body3.Gravity(body1)
		body3Gravity.Add(body3.Gravity(body2))
		body3.Update(body3Gravity)

		win.Clear(colornames.White)

		// Draw sprites
		body1Sprite.Clear()
		body1Sprite.Push(pixel.V(body1.XPosition, body1.YPosition))
		body1Sprite.Circle(body1.Radius, 2)
		body1Sprite.Draw(win)

		body2Sprite.Clear()
		body2Sprite.Push(pixel.V(body2.XPosition, body2.YPosition))
		body2Sprite.Circle(body2.Radius, 2)
		body2Sprite.Draw(win)

		body3Sprite.Clear()
		body3Sprite.Push(pixel.V(body3.XPosition, body3.YPosition))
		body3Sprite.Circle(body3.Radius, 2)
		body3Sprite.Draw(win)

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
