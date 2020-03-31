package game

import (
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/jjrumi/asteroids/pkg/game/internal"
)

const numAsteroids = 10

type world struct {
	win       *pixelgl.Window
	width     float64
	height    float64
	ship      internal.Ship
	asteroids []internal.Asteroid
}

type World interface {
	GameLoop()
}

func NewWorld(title string, width float64, height float64) World {
	rand.Seed(time.Now().UnixNano())

	return &world{
		win:       mustBuildWindow(title, width, height),
		width:     width,
		height:    height,
		ship:      internal.NewShip(pixel.V(width/2, height/2)),
		asteroids: spawnFlockOfAsteroids(width, height),
	}
}

func mustBuildWindow(title string, width float64, height float64) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}

func spawnFlockOfAsteroids(width float64, height float64) []internal.Asteroid {
	asteroids := make([]internal.Asteroid, numAsteroids)
	for i := 0; i < numAsteroids; i++ {
		asteroid := internal.NewAsteroid(pixel.V(rand.Float64()*width, rand.Float64()*height))
		asteroids[i] = asteroid
	}

	return asteroids
}

func (w *world) GameLoop() {
	for !w.win.Closed() {
		w.processInput()
		w.updateGame()
		w.render()
	}
}

func (w *world) processInput() {
	if w.win.Pressed(pixelgl.KeyLeft) {
		w.ship.RotateLeft()
	}
	if w.win.Pressed(pixelgl.KeyRight) {
		w.ship.RotateRight()
	}
	if w.win.Pressed(pixelgl.KeyUp) {
		w.ship.Thrust()
	}
}

func (w *world) updateGame() {
	w.ship.Update(w.width, w.height)

	for _, asteroid := range w.asteroids {
		asteroid.Update(w.width, w.height)

		if w.ship.DetectCollision(asteroid) {
			log.Printf("COLLISION DETECTED!!!")
		}
	}
}

func (w *world) render() {
	w.win.Clear(colornames.Black)

	w.ship.Render(w.win)
	for _, asteroid := range w.asteroids {
		asteroid.Render(w.win)
	}

	w.win.Update()
}
