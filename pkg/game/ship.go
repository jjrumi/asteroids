package game

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const shipSize = 20
const thrust = 0.2
const shipFriction = 0.99

type Ship struct {
	*Object
}

func NewShip(pos pixel.Vec) *Ship {
	points := []pixel.Vec{
		pixel.V(-shipSize, -shipSize),
		pixel.V(shipSize, -shipSize),
		pixel.V(0, shipSize),
	}

	ship := &Ship{
		Object: &Object{
			IMDraw:       imdraw.New(nil),
			points:       points,
			heading:      math.Pi / 2,
			position:     pixel.ZV,
			velocity:     pixel.ZV,
			acceleration: pixel.ZV,
		},
	}

	ship.moveBy(pos)

	return ship
}

func (s *Ship) Update(winWidth float64, winHeight float64) {
	s.velocity = s.velocity.Add(s.acceleration)
	s.velocity = s.velocity.Scaled(shipFriction)
	s.acceleration = pixel.ZV

	s.updatePosition(winWidth, winHeight)
	s.redrawShip()
}

func (s *Ship) redrawShip() {
	s.Clear()
	s.Reset()

	s.Color = pixel.RGB(1, 1, 1)
	s.Push(s.points[0])
	s.Push(s.points[1])
	s.Color = pixel.RGB(1, 0, 0)
	s.Push(s.points[2])
	s.Polygon(1)
}

func (s *Ship) Thrust() {
	s.acceleration.X += thrust * math.Cos(s.heading)
	s.acceleration.Y += thrust * math.Sin(s.heading)
}
