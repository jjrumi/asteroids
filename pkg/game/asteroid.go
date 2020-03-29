package game

import (
	"math"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const minRadius = 10
const maxRadius = 20
const asteroidPoints = 10

type Asteroid struct {
	*Object
}

func NewAsteroid(pos pixel.Vec) *Asteroid {
	radius := (rand.Float64() * minRadius) + minRadius
	angle := float64(0)
	points := make([]pixel.Vec, asteroidPoints)
	for i := float64(0); i < asteroidPoints; i++ {
		angle = (2 * math.Pi / asteroidPoints) * i
		relRadius := radius*rand.Float64() + maxRadius
		points[int(i)] = pixel.V(relRadius*math.Cos(angle), relRadius*math.Sin(angle))
	}

	vX := -1 + rand.Float64()*(1+1)
	vY := -1 + rand.Float64()*(1+1)
	velocity := pixel.V(vX, vY)

	asteroid := &Asteroid{
		Object: &Object{
			IMDraw:       imdraw.New(nil),
			points:       points,
			heading:      math.Pi / 2,
			position:     pixel.ZV,
			velocity:     velocity,
			acceleration: pixel.ZV,
		},
	}

	asteroid.moveBy(pos)

	return asteroid
}

func (s *Asteroid) Update(winWidth float64, winHeight float64) {
	s.updatePosition(winWidth, winHeight)
	s.redrawAsteroid()
}

func (s *Asteroid) redrawAsteroid() {
	s.Clear()
	s.Reset()

	s.Color = pixel.RGB(1, 1, 1)
	s.Push(s.points...)
	s.Polygon(1)
}
