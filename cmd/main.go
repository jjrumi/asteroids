package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/jjrumi/asteroids/pkg/game"
)

const (
	WIDTH  = 1024
	HEIGHT = 768
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Asteroids!",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	ship := game.NewShip(pixel.V(WIDTH/2, HEIGHT/2))

	// TODO:
	// 2- Rotation: left & right arrows -> Apply Rotation matrix to ship points

	for !win.Closed() {
		if win.Pressed(pixelgl.KeyLeft) {
			ship.RotateLeft()
		}
		if win.Pressed(pixelgl.KeyRight) {
			ship.RotateRight()
		}
		if win.Pressed(pixelgl.KeyUp) {
			ship.Accelerate()
		}

		win.Clear(colornames.Black)
		ship.ReDraw()
		ship.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
