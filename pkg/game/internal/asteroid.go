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
	Create(a Asteroid)
	UpdateElements(winWidth float64, winHeight float64)
	RenderElements(target pixel.Target)
	DetectShipCollision(s Ship) bool
	HandleBoltCollision(boltPool LaserBoltPool) []Asteroid
}

func NewAsteroidPool() AsteroidPool {
	return &asteroidPool{
		pool: newPool(),
	}
}

type asteroidPool struct {
	pool pool
}

func (p *asteroidPool) Create(a Asteroid) {
	p.pool.create(a.(*asteroid))
}

func (p *asteroidPool) UpdateElements(winWidth float64, winHeight float64) {
	for _, e := range p.pool.list() {
		e.(*asteroid).Update(winWidth, winHeight)
	}

	p.pool.purge()
}

func (p *asteroidPool) RenderElements(target pixel.Target) {
	for _, e := range p.pool.list() {
		e.(*asteroid).Render(target)
	}
}

func (p *asteroidPool) DetectShipCollision(s Ship) bool {
	for _, e := range p.pool.list() {
		if e.(*asteroid).polygon.collides(s.(*ship).polygon) {
			return true
		}
	}
	return false
}

func (p *asteroidPool) HandleBoltCollision(boltPool LaserBoltPool) []Asteroid {
	destroyed := make([]Asteroid, 0)
	bp := boltPool.(*laserBoltPool)
	for _, a := range p.pool.list() {
		for _, b := range bp.pool.list() {
			if b.isAlive() && a.(*asteroid).polygon.collides(b.(*laserBolt).polygon) {
				a.(*asteroid).destroy()
				b.(*laserBolt).destroy()
				destroyed = append(destroyed, a.(*asteroid))
				continue
			}
		}
	}

	return destroyed
}

type Asteroid gameComponent

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

func (a *asteroid) Update(winWidth float64, winHeight float64) {
	a.polygon.update(winWidth, winHeight)
}

func (a *asteroid) Render(target pixel.Target) {
	a.polygon.render(target)
}

func (a *asteroid) destroy() {
	a.alive = false
}

func (a *asteroid) isAlive() bool {
	return a.alive
}

type AsteroidBlastPool interface {
	Create(destroyed []Asteroid)
	UpdateElements(winWidth float64, winHeight float64)
	RenderElements(target pixel.Target)
}

func NewAsteroidBlastPool() AsteroidBlastPool {
	return &asteroidBlastPool{
		pool: newPool(),
	}
}

type asteroidBlastPool struct {
	pool pool
}

func (p *asteroidBlastPool) Create(list []Asteroid) {
	for _, a := range list {
		blast := NewAsteroidBlast(a)
		p.pool.create(blast.(*asteroidBlast))
	}
}

func (p *asteroidBlastPool) UpdateElements(winWidth float64, winHeight float64) {
	for _, e := range p.pool.list() {
		e.(*asteroidBlast).Update(winWidth, winHeight)
	}

	p.pool.purge()
}

func (p *asteroidBlastPool) RenderElements(target pixel.Target) {
	for _, e := range p.pool.list() {
		e.(*asteroidBlast).Render(target)
	}
}

type AsteroidBlast gameComponent

func NewAsteroidBlast(a Asteroid) AsteroidBlast {
	log.Printf("Blast at: %v", a.(*asteroid).polygon.position)
	return &asteroidBlast{
		polygon: &polygon{
			IMDraw:   imdraw.New(nil),
			points:   nil, // TODO: points?!
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

	b.polygon.update(winWidth, winHeight)
}

func (b *asteroidBlast) Render(target pixel.Target) {
	if !b.isAlive() {
		return
	}

	b.polygon.render(target)
}

func (b *asteroidBlast) isAlive() bool {
	return b.lifeLeft > 0
}
