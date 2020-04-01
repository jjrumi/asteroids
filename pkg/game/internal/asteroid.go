package internal

import (
	"log"
	"math"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const radiusOffset = 5
const minRadius = 20
const maxRadius = 70
const asteroidPoints = 10
const blastLifeTime = 30

type AsteroidPool interface {
	Object
	Create(pos pixel.Vec)
	DetectShipCollision(s Ship) bool
	HandleBoltCollision(boltPool LaserBoltPool) []Asteroid
}

func NewAsteroidPool() AsteroidPool {
	return &asteroidPool{
		pool: make([]Asteroid, 0),
	}
}

type asteroidPool struct {
	pool []Asteroid
}

func (p *asteroidPool) Create(pos pixel.Vec) {
	p.pool = append(p.pool, NewAsteroid(pos))
}

func (p *asteroidPool) Update(winWidth float64, winHeight float64) {
	for _, asteroid := range p.pool {
		asteroid.Update(winWidth, winHeight)
	}

	// Get rid of dead asteroids:
	var newPool []Asteroid
	for _, asteroid := range p.pool {
		if asteroid.isAlive() {
			newPool = append(newPool, asteroid)
		}
	}
	p.pool = newPool
}

func (p *asteroidPool) Render(target pixel.Target) {
	for _, asteroid := range p.pool {
		asteroid.Render(target)
	}
}

func (p *asteroidPool) DetectShipCollision(s Ship) bool {
	for _, a := range p.pool {
		if a.(*asteroid).polygon.collides(s.(*ship).polygon) {
			return true
		}
	}
	return false
}

func (p *asteroidPool) HandleBoltCollision(boltPool LaserBoltPool) []Asteroid {
	destroyed := make([]Asteroid, 0)
	bp := boltPool.(*laserBoltPool)
	for _, a := range p.pool {
		for _, b := range bp.pool {
			if b.isAlive() && a.(*asteroid).polygon.collides(b.(*laserBolt).polygon) {
				a.destroy()
				b.destroy()
				destroyed = append(destroyed, a)
				continue
			}
		}
	}

	return destroyed
}

type Asteroid interface {
	Object
	destroy()
	isAlive() bool
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
		polygon: &polygon{
			IMDraw:         imdraw.New(nil),
			points:         points,
			heading:        math.Pi / 2,
			position:       pixel.ZV,
			velocity:       velocity,
			acceleration:   pixel.ZV,
			boundingRadius: boundingRadius,
		},
		alive: true,
	}

	asteroid.moveBy(pos)

	return asteroid
}

type asteroid struct {
	*polygon
	alive bool
}

func (a *asteroid) destroy() {
	a.alive = false
}

func (a *asteroid) isAlive() bool {
	return a.alive
}

type AsteroidBlastPool interface {
	Object
	Create(destroyed []Asteroid)
}

func NewAsteroidBlastPool() AsteroidBlastPool {
	return &asteroidBlastPool{}
}

type asteroidBlastPool struct {
	pool []AsteroidBlast
}

func (p *asteroidBlastPool) Create(list []Asteroid) {
	for _, a := range list {
		p.pool = append(p.pool, NewAsteroidBlast(a))
	}
}

func (p *asteroidBlastPool) Update(winWidth float64, winHeight float64) {
	for _, blast := range p.pool {
		blast.Update(winWidth, winHeight)
	}

	// Get rid of dead bolts:
	var newPool []AsteroidBlast
	for _, blast := range p.pool {
		if blast.isAlive() {
			newPool = append(newPool, blast)
		}
	}
	p.pool = newPool
}

func (p *asteroidBlastPool) Render(target pixel.Target) {
	for _, blast := range p.pool {
		blast.Render(target)
	}
}

type AsteroidBlast interface {
	Object
	isAlive() bool
}

func NewAsteroidBlast(a Asteroid) AsteroidBlast {
	log.Printf("Blast at: %v", a.(*asteroid).polygon.position)
	return &asteroidBlast{
		polygon: &polygon{
			IMDraw: imdraw.New(nil),
			points: nil, // TODO: points?!
			position: a.(*asteroid).polygon.position,

		},
		lifeLeft: blastLifeTime,
	}
}

type asteroidBlast struct {
	*polygon
	lifeLeft float64
}

func (b *asteroidBlast) Update(winWidth float64, winHeight float64) {
	if !b.isAlive() {
		return
	}

	b.lifeLeft--

	b.polygon.Update(winWidth, winHeight)
}

func (b *asteroidBlast) Render(target pixel.Target) {
	if !b.isAlive() {
		return
	}

	b.polygon.Render(target)
}

func (b *asteroidBlast) isAlive() bool {
	return b.lifeLeft > 0
}
