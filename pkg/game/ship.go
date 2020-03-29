package game

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const shipSize = 20
const thrust = 0.2
const spaceFriction = 0.99

type Ship struct {
	*imdraw.IMDraw
	points       []pixel.Vec
	heading      float64
	position     pixel.Vec
	velocity     pixel.Vec
	acceleration pixel.Vec
}

func NewShip(pos pixel.Vec) *Ship {
	points := []pixel.Vec{
		pixel.V(-shipSize, -shipSize),
		pixel.V(shipSize, -shipSize),
		pixel.V(0, shipSize),
	}

	ship := &Ship{
		IMDraw:       imdraw.New(nil),
		points:       points,
		heading:      math.Pi / 2,
		position:     pixel.ZV,
		velocity:     pixel.ZV,
		acceleration: pixel.ZV,
	}

	ship.moveShipBy(pos)

	return ship
}

func (s *Ship) moveShipBy(v pixel.Vec) {
	s.position = s.position.Add(v)

	newPoints := s.points[:0]
	for _, point := range s.points {
		np := point.Add(v)
		newPoints = append(newPoints, np)
	}
}

func (s *Ship) Update(winWidth float64, winHeight float64) {
	s.updatePosition(winWidth, winHeight)
	s.redrawShip()
}

func (s *Ship) updatePosition(screenWidth float64, screenHeight float64) {
	s.moveShipBy(s.velocity)
	s.velocity = s.velocity.Add(s.acceleration)
	s.velocity = s.velocity.Scaled(spaceFriction)
	s.acceleration = pixel.ZV

	// keep ship on the screen - go over the edge
	if s.position.Y > screenHeight+shipSize {
		s.moveShipBy(pixel.V(0, -screenHeight-shipSize))
	}
	if s.position.X > screenWidth+shipSize {
		s.moveShipBy(pixel.V(-screenWidth-shipSize, 0))
	}
	if s.position.Y < 0-shipSize {
		s.moveShipBy(pixel.V(0, screenHeight+shipSize))
	}
	if s.position.X < 0-shipSize {
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
	s.heading += angle

	newPoints := s.points[:0]
	for _, point := range s.points {
		newP := pixel.IM.
			Moved(pixel.V(-s.position.X, -s.position.Y)).
			Rotated(pixel.ZV, angle).
			Moved(s.position).
			Project(point)
		newPoints = append(newPoints, newP)
	}
}

func (s *Ship) Thrust() {
	s.acceleration.X += thrust * math.Cos(s.heading)
	s.acceleration.Y += thrust * math.Sin(s.heading)
}
