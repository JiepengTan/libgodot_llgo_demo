package scripts

import (
	"grow.graphics/gd"
)

type Pipe struct {
	gd.Class[Pipe, gd.StaticBody2D] `gd:"Pipe"`

	Height        gd.Int
	IsTraversable gd.Bool

	CollisionShape2D gd.CollisionShape2D `gd:"CollisionShape2D"`
	PipeBodySprite   gd.Sprite           `gd:"PipeBodySprite"`
}

const TOP_PIPE_HEIGHT = 16

func (pself *Pipe) Ready() {
	regionRect := gd.NewRect2(pself.PipeBodySprite.RegionRect())
	regionRect.SetSize(gd.NewVector2(32, pself.Height-TOP_PIPE_HEIGHT))
	pself.PipeBodySprite.SetRegionRect(regionRect)
	pself.PipeBodySprite.SetPosition(gd.NewVector2(0, pself.Height/2))

	shape := gd.NewRectangleShape2D()
	shape.SetSize(gd.NewVector2(32, pself.Height))
	pself.CollisionShape2D.SetShape(shape)
	pself.CollisionShape2D.SetPosition(gd.NewVector2(0, pself.Height/2-TOP_PIPE_HEIGHT/2))
}
