package internal

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const laserBoltSpeed float64 = 9
const laserBoltLifeTime float64 = 60

type LaserBoltPool interface {
	Create(bolt LaserBolt)
	UpdateElements(winWidth float64, winHeight float64)
	RenderElements(target pixel.Target)
}

func NewLaserBoltPool() LaserBoltPool {
	return &laserBoltPool{
		pool: newPool(),
	}
}

type laserBoltPool struct {
	pool pool
}

func (p *laserBoltPool) Create(bolt LaserBolt) {
	p.pool.create(bolt.(*laserBolt))
}

func (p *laserBoltPool) UpdateElements(winWidth float64, winHeight float64) {
	for _, e := range p.pool.list() {
		e.(*laserBolt).Update(winWidth, winHeight)
	}

	p.pool.purge()
}

func (p *laserBoltPool) RenderElements(target pixel.Target) {
	for _, e := range p.pool.list() {
		e.(*laserBolt).Render(target)
	}
}

type LaserBolt gameComponent

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

	l.polygon.update(winWidth, winHeight)
}

func (l *laserBolt) Render(target pixel.Target) {
	if !l.isAlive() {
		return
	}

	l.polygon.render(target)
}

func (l *laserBolt) isAlive() bool {
	return l.lifeLeft > 0 && l.alive
}

func (l *laserBolt) destroy() {
	l.alive = false
}
