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
	win        *pixelgl.Window
	width      float64
	height     float64
	ship       internal.Ship
	asteroids  internal.AsteroidPool
	laserBolts internal.LaserBoltPool
	blasts     internal.AsteroidBlastPool
}

type World interface {
	GameLoop()
}

func NewWorld(title string, width float64, height float64) World {
	rand.Seed(time.Now().UnixNano())

	return &world{
		win:        mustBuildWindow(title, width, height),
		width:      width,
		height:     height,
		ship:       internal.NewShip(pixel.V(width/2, height/2)),
		asteroids:  spawnFlockOfAsteroids(width, height),
		laserBolts: internal.NewLaserBoltPool(),
		blasts:     internal.NewAsteroidBlastPool(),
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

func spawnFlockOfAsteroids(width float64, height float64) internal.AsteroidPool {
	pool := internal.NewAsteroidPool()
	for i := 0; i < numAsteroids; i++ {
		pool.Create(pixel.V(rand.Float64()*width, rand.Float64()*height))
	}

	return pool
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
	if w.win.JustPressed(pixelgl.KeySpace) {
		w.laserBolts.Create(w.ship.Fire())
	}
}

func (w *world) updateGame() {
	w.ship.Update(w.width, w.height)
	w.asteroids.Update(w.width, w.height)
	w.laserBolts.Update(w.width, w.height)
	w.blasts.Update(w.width, w.height)

	if w.asteroids.DetectShipCollision(w.ship) {
		log.Printf("GAME OVER !!")
	}

	destroyed := w.asteroids.HandleBoltCollision(w.laserBolts)
	w.blasts.Create(destroyed)
	// TODO: Create smaller asteroids if possible
	// w.asteroids.CreateFromBlast(destroyed)
}

func (w *world) render() {
	w.win.Clear(colornames.Black)

	w.ship.Render(w.win)
	w.asteroids.Render(w.win)
	w.laserBolts.Render(w.win)
	w.blasts.Render(w.win)

	w.win.Update()
}
