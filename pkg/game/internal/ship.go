package internal

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const shipSize = 20
const thrust = 0.2
const rotationAngle = math.Pi / 30
const shipFriction = 0.99

type Ship interface {
	Object

	Thrust()
	RotateLeft()
	RotateRight()
	Fire()

	DetectCollision(a Asteroid) bool
}

type ship struct {
	*object
}

func NewShip(pos pixel.Vec) Ship {
	points := []pixel.Vec{
		pixel.V(-shipSize, -shipSize),
		pixel.V(0, -shipSize/4),
		pixel.V(shipSize, -shipSize),
		pixel.V(0, shipSize),
	}

	ship := &ship{
		object: &object{
			IMDraw:         imdraw.New(nil),
			points:         points,
			heading:        math.Pi / 2,
			position:       pixel.ZV,
			velocity:       pixel.ZV,
			acceleration:   pixel.ZV,
			boundingRadius: shipSize,
		},
	}

	ship.moveBy(pos)

	return ship
}

func (s *ship) Update(winWidth float64, winHeight float64) {
	s.velocity = s.velocity.Add(s.acceleration)
	s.velocity = s.velocity.Scaled(shipFriction)
	s.acceleration = pixel.ZV

	s.object.Update(winWidth, winHeight)
}

func (s *ship) Thrust() {
	s.acceleration.X += thrust * math.Cos(s.heading)
	s.acceleration.Y += thrust * math.Sin(s.heading)
}
func (s *ship) RotateLeft() {
	s.rotate(rotationAngle)
}
func (s *ship) RotateRight() {
	s.rotate(-rotationAngle)
}

func (s *ship) rotate(angle float64) {
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

func (s *ship) DetectCollision(a Asteroid) bool {
	return s.object.collides(a.(*asteroid).object)
}

func (s *ship) Fire() {
	panic("todo")
}
