package game

import (
	"math"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const minRadius = 10
const maxRadius = 20

type Asteroid struct {
	*imdraw.IMDraw
	points   []pixel.Vec
	pos      pixel.Vec
	dir      pixel.Vec
	velocity pixel.Vec
}

func NewAsteroid(pos pixel.Vec) *Asteroid {
	radius := (rand.Float64() * 10) + 10
	angle := float64(0)
	points := make([]pixel.Vec, 10)
	for i := float64(0); i < 10; i++ {
		angle = (2 * math.Pi / 10) * i
		relRadius := radius*rand.Float64() + maxRadius
		points[int(i)] = pixel.V(relRadius*math.Cos(angle), relRadius*math.Sin(angle))
	}

	vX := -1 + rand.Float64()*(1+1)
	vY := -1 + rand.Float64()*(1+1)
	velocity := pixel.V(vX, vY)

	asteroid := &Asteroid{
		imdraw.New(nil),
		points,
		pixel.ZV,
		pixel.V(0, 1),
		velocity,
	}

	asteroid.moveAsteroidBy(pos)

	return asteroid
}

func (s *Asteroid) Update(winWidth float64, winHeight float64) {
	s.updatePosition(winWidth, winHeight)
	s.redrawAsteroid()
}

func (s *Asteroid) moveAsteroidBy(v pixel.Vec) {
	tm := pixel.IM.Moved(v)
	s.pos = tm.Project(s.pos)

	newPoints := s.points[:0]
	for _, point := range s.points {
		newPoints = append(newPoints, tm.Project(point))
	}
}

func (s *Asteroid) updatePosition(screenWidth float64, screenHeight float64) {
	s.moveAsteroidBy(pixel.V(s.velocity.X, s.velocity.Y))

	// go over the edge
	if s.pos.Y > screenHeight+shipSize {
		s.moveAsteroidBy(pixel.V(0, -screenHeight-shipSize))
	}
	if s.pos.X > screenWidth+shipSize {
		s.moveAsteroidBy(pixel.V(-screenWidth-shipSize, 0))
	}
	if s.pos.Y < 0-shipSize {
		s.moveAsteroidBy(pixel.V(0, screenHeight+shipSize))
	}
	if s.pos.X < 0-shipSize {
		s.moveAsteroidBy(pixel.V(screenWidth+shipSize, 0))
	}
}

func (s *Asteroid) redrawAsteroid() {
	s.Clear()
	s.Reset()

	s.Color = pixel.RGB(1, 1, 1)
	s.Push(s.points...)
	s.Polygon(1)
}

func (s *Asteroid) Rotate(angle float64) {
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
