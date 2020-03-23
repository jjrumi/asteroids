package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Ship struct {
	*imdraw.IMDraw
	pos pixel.Vec
	r   float64
}

func NewShip(pos pixel.Vec) *Ship {
	ship := &Ship{
		imdraw.New(nil),
		pos,
		20,
	}

	ship.DrawShip()

	return ship
}

func (s *Ship) DrawShip() {
	s.Reset()

	s.Color = pixel.RGB(255, 255, 255)
	s.Push(pixel.V(s.pos.X-s.r, s.pos.Y-s.r))
	s.Push(pixel.V(s.pos.X+s.r, s.pos.Y-s.r))
	s.Push(pixel.V(s.pos.X, s.pos.Y+s.r+s.r/2))
	s.Polygon(1)
}

func (s *Ship) Move() {
	pixel.IM.Moved(pixel.V(10, 10))
}
