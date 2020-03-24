package game

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const shipSize = 20

type Ship struct {
	*imdraw.IMDraw
	pos    pixel.Vec
	points []pixel.Vec
	radius float64
}

func NewShip(pos pixel.Vec) *Ship {
	points := []pixel.Vec{
		pixel.V(pos.X-shipSize, pos.Y-shipSize),
		pixel.V(pos.X+shipSize, pos.Y-shipSize),
		pixel.V(pos.X, pos.Y+shipSize+shipSize/2),
	}

	return &Ship{
		imdraw.New(nil),
		pos,
		points,
		shipSize,
	}
}

func (s *Ship) ReDraw() {
	s.Clear()
	s.Reset()

	s.Color = pixel.RGB(255, 255, 255)
	s.Push(s.points...)
	s.Polygon(1)
}

func (s *Ship) RotateLeft() {
	log.Print("rotating left...")
}

func (s *Ship) RotateRight() {
	log.Print("rotating right...")
}

func (s *Ship) Accelerate() {
	log.Print("accelerating...")
	s.pos = pixel.IM.Moved(pixel.V(0, 5)).Project(s.pos)
	newPoints := s.points[:0]
	for _, point := range s.points {
		newPoints = append(newPoints, pixel.IM.Moved(pixel.V(0,5)).Project(point))
	}
	s.points = newPoints
}
