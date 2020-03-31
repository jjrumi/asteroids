package internal

import (
	"math"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const radiusOffset = 5
const minRadius = 20
const maxRadius = 70
const asteroidPoints = 10

type Asteroid interface {
	Object
}

type asteroid struct {
	*object
}

func NewAsteroid(pos pixel.Vec) Asteroid {
	radius := minRadius + rand.Float64()*(maxRadius-minRadius)
	angle := float64(0)
	points := make([]pixel.Vec, asteroidPoints)
	var boundingRadius float64
	for i := float64(0); i < asteroidPoints; i++ {
		angle = (2 * math.Pi / asteroidPoints) * i
		pointRadius := (minRadius - radiusOffset) + rand.Float64()*(radius-(minRadius-radiusOffset))
		points[int(i)] = pixel.V(pointRadius*math.Cos(angle), pointRadius*math.Sin(angle))
		boundingRadius = math.Max(pointRadius, boundingRadius)
	}

	vX := -1 + rand.Float64()*(1+1)
	vY := -1 + rand.Float64()*(1+1)
	velocity := pixel.V(vX, vY)

	asteroid := &asteroid{
		object: &object{
			IMDraw:         imdraw.New(nil),
			points:         points,
			heading:        math.Pi / 2,
			position:       pixel.ZV,
			velocity:       velocity,
			acceleration:   pixel.ZV,
			boundingRadius: boundingRadius,
		},
	}

	asteroid.moveBy(pos)

	return asteroid
}

func (a *asteroid) Update(winWidth float64, winHeight float64) {
	a.updatePosition(winWidth, winHeight)
}
