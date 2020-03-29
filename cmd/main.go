package main

import (
	"github.com/faiface/pixel/pixelgl"

	"github.com/jjrumi/asteroids/pkg/game"
)

const (
	WinWidth  = 1024
	WinHeight = 768
	WinTitle = "Asteroids!"
)

func run() {
	world := game.NewWorld(WinTitle, WinWidth, WinHeight)
	world.GameLoop()
}

func main() {
	pixelgl.Run(run)
}
