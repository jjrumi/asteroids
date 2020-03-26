package game

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const shipSize = 20
const maxSpeed = 10
const thrust = 0.2

type Ship struct {
	*imdraw.IMDraw
	points   []pixel.Vec
	pos      pixel.Vec
	dir      pixel.Vec
	velocity pixel.Vec
}

func NewShip(pos pixel.Vec) *Ship {
	points := []pixel.Vec{
		pixel.V(-shipSize, -shipSize),
		pixel.V(shipSize, -shipSize),
		pixel.V(0, shipSize),
	}

	ship := &Ship{
		imdraw.New(nil),
		points,
		pixel.ZV,
		pixel.V(0, 1),
		pixel.ZV,
	}

	ship.moveShipBy(pos)

	return ship
}

func (s *Ship) Update(screenWidth int, screenHeight int) {
	s.updatePosition(float64(screenWidth), float64(screenHeight))
	s.redrawShip()
}

func (s *Ship) moveShipBy(v pixel.Vec) {
	tm := pixel.IM.Moved(v)
	s.pos = tm.Project(s.pos)

	newPoints := s.points[:0]
	for _, point := range s.points {
		newPoints = append(newPoints, tm.Project(point))
	}
}

func (s *Ship) updatePosition(screenWidth float64, screenHeight float64) {
	s.moveShipBy(pixel.V(s.velocity.X, s.velocity.Y))

	// go over the edge
	if s.pos.Y > screenHeight+shipSize {
		s.moveShipBy(pixel.V(0, -screenHeight-shipSize))
	}
	if s.pos.X > screenWidth+shipSize {
		s.moveShipBy(pixel.V(-screenWidth-shipSize, 0))
	}
	if s.pos.Y < 0-shipSize {
		s.moveShipBy(pixel.V(0, screenHeight+shipSize))
	}
	if s.pos.X < 0-shipSize {
		s.moveShipBy(pixel.V(screenWidth+shipSize, 0))
	}
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

func (s *Ship) Rotate(angle float64) {
	s.dir = pixel.IM.
		Rotated(pixel.ZV, angle).
		Project(s.dir)

	newPoints := s.points[:0]
	for _, point := range s.points {
		newP := pixel.IM.
			Moved(pixel.V(-s.pos.X, -s.pos.Y)).
			Rotated(pixel.ZV, angle).
			Moved(s.pos).
			Project(point)
		newPoints = append(newPoints, newP)
	}
}

func (s *Ship) Thrust() {
	s.velocity = pixel.V(
		math.Min(maxSpeed, s.dir.X*thrust+s.velocity.X),
		math.Min(maxSpeed, s.dir.Y*thrust+s.velocity.Y),
	)
}
