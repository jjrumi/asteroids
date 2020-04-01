package internal

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Object interface {
	Update(winWidth float64, winHeight float64)
	Render(target pixel.Target)
}

type polygon struct {
	*imdraw.IMDraw
	points         []pixel.Vec
	heading        float64
	position       pixel.Vec
	velocity       pixel.Vec
	acceleration   pixel.Vec
	boundingRadius float64
}

func (o *polygon) Render(target pixel.Target) {
	o.Clear()
	o.Reset()

	o.Color = pixel.RGB(1, 1, 1)
	o.Push(o.points...)
	o.Polygon(2)

	o.Draw(target)
}

func (o *polygon) Update(screenWidth float64, screenHeight float64) {
	o.moveBy(o.velocity)

	// keep polygon on the screen - go over the edge
	if o.position.Y > screenHeight+shipSize {
		o.moveBy(pixel.V(0, -screenHeight-shipSize))
	}
	if o.position.X > screenWidth+shipSize {
		o.moveBy(pixel.V(-screenWidth-shipSize, 0))
	}
	if o.position.Y < 0-shipSize {
		o.moveBy(pixel.V(0, screenHeight+shipSize))
	}
	if o.position.X < 0-shipSize {
		o.moveBy(pixel.V(screenWidth+shipSize, 0))
	}
}

func (o *polygon) moveBy(v pixel.Vec) {
	o.position = o.position.Add(v)

	newPoints := o.points[:0]
	for _, point := range o.points {
		np := point.Add(v)
		newPoints = append(newPoints, np)
	}
}

func (o *polygon) collides(o2 *polygon) bool {
	return o.detectRadiusOverlappingCollision(o2)
}

func (o *polygon) detectRadiusOverlappingCollision(o2 *polygon) bool {
	distanceX := o.position.X - o2.position.X
	distanceY := o.position.Y - o2.position.Y
	distance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)

	return distance < (o.boundingRadius + o2.boundingRadius)
}
