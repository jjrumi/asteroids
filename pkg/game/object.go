package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Object struct {
	*imdraw.IMDraw
	points       []pixel.Vec
	heading      float64
	position     pixel.Vec
	velocity     pixel.Vec
	acceleration pixel.Vec
}

func (o *Object) moveBy(v pixel.Vec) {
	o.position = o.position.Add(v)

	newPoints := o.points[:0]
	for _, point := range o.points {
		np := point.Add(v)
		newPoints = append(newPoints, np)
	}
}

func (o *Object) updatePosition(screenWidth float64, screenHeight float64) {
	o.moveBy(o.velocity)

	// keep object on the screen - go over the edge
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

func (o *Object) rotate(angle float64) {
	o.heading += angle

	newPoints := o.points[:0]
	for _, point := range o.points {
		newP := pixel.IM.
			Moved(pixel.V(-o.position.X, -o.position.Y)).
			Rotated(pixel.ZV, angle).
			Moved(o.position).
			Project(point)
		newPoints = append(newPoints, newP)
	}
}
