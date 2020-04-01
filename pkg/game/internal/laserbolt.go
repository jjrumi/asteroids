package internal

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const laserBoltSpeed float64 = 9
const laserBoltLifeTime float64 = 60

type LaserBoltPool interface {
	Object
	Create(bolt LaserBolt)
}

func NewLaserBoltPool() LaserBoltPool {
	return &laserBoltPool{
		pool: make([]LaserBolt, 0),
	}
}

type laserBoltPool struct {
	pool []LaserBolt
}

func (p *laserBoltPool) Create(bolt LaserBolt) {
	p.pool = append(p.pool, bolt)
}

func (p *laserBoltPool) Update(winWidth float64, winHeight float64) {
	for _, bolt := range p.pool {
		bolt.Update(winWidth, winHeight)
	}

	// Get rid of dead bolts:
	var newPool []LaserBolt
	for _, bolt := range p.pool {
		if bolt.isAlive() {
			newPool = append(newPool, bolt)
		}
	}
	p.pool = newPool
}

func (p *laserBoltPool) Render(target pixel.Target) {
	for _, bolt := range p.pool {
		bolt.Render(target)
	}
}

type LaserBolt interface {
	Object
	isAlive() bool
	destroy()
}

type laserBolt struct {
	*polygon
	lifeLeft float64
	alive    bool
}

func NewLaserBolt(pos pixel.Vec, angle float64) LaserBolt {
	velocity := pixel.Unit(angle).Scaled(laserBoltSpeed)

	points := []pixel.Vec{
		pos,
		pixel.V(pos.X+velocity.X, pos.Y+velocity.Y),
	}

	return &laserBolt{
		polygon: &polygon{
			IMDraw:         imdraw.New(nil),
			points:         points,
			position:       pos,
			velocity:       velocity,
			acceleration:   pixel.ZV,
			boundingRadius: 1,
		},
		lifeLeft: laserBoltLifeTime,
		alive:    true,
	}
}

func (l *laserBolt) Update(winWidth float64, winHeight float64) {
	if !l.isAlive() {
		return
	}

	l.lifeLeft--

	l.polygon.Update(winWidth, winHeight)
}

func (l *laserBolt) Render(target pixel.Target) {
	if !l.isAlive() {
		return
	}

	l.polygon.Render(target)
}

func (l *laserBolt) isAlive() bool {
	return l.lifeLeft > 0 && l.alive
}

func (l *laserBolt) destroy() {
	l.alive = false
}
