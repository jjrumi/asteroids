package game

import (
	"math"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type World struct {
	win       *pixelgl.Window
	width     float64
	height    float64
	ship      *Ship
	asteroids []*Asteroid
}

func NewWorld(title string, width float64, height float64) *World {
	rand.Seed(time.Now().UnixNano())

	return &World{
		win:       mustBuildWindow(title, width, height),
		width:     width,
		height:    height,
		ship:      NewShip(pixel.V(width/2, height/2)),
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

func spawnFlockOfAsteroids(width float64, height float64) []*Asteroid {
	asteroids := make([]*Asteroid, 10)
	for i := 0; i < 10; i++ {
		asteroid := NewAsteroid(pixel.V(rand.Float64()*width, rand.Float64()*height))
		asteroids[i] = asteroid
	}

	return asteroids
}

func (w *World) GameLoop() {
	for !w.win.Closed() {
		w.processInput()
		w.updateGame()
		w.render()
	}
}

func (w *World) processInput() {
	if w.win.Pressed(pixelgl.KeyLeft) {
		w.ship.Rotate(math.Pi / 20)
	}
	if w.win.Pressed(pixelgl.KeyRight) {
		w.ship.Rotate(-math.Pi / 20)
	}
	if w.win.Pressed(pixelgl.KeyUp) {
		w.ship.Thrust()
	}
}

func (w *World) updateGame() {
	w.ship.Update(w.width, w.height)

	for _, asteroid := range w.asteroids {
		asteroid.Update(w.width, w.height)
	}
}

func (w *World) render() {
	w.win.Clear(colornames.Black)

	w.ship.Draw(w.win)
	for _, asteroid := range w.asteroids {
		asteroid.Draw(w.win)
	}

	w.win.Update()
}
