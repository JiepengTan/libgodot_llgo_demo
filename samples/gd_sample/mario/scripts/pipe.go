package scripts

import (
	"grow.graphics/gd"
)

type Pipe struct {
	gd.Class[Pipe, gd.StaticBody2D] `gd:"Pipe"`

	Height        gd.Float
	IsTraversable gd.Bool

	CollisionShape2D gd.CollisionShape2D `gd:"CollisionShape2D"`
	PipeBodySprite   gd.Sprite2D         `gd:"PipeBodySprite"`
}

const TOP_PIPE_HEIGHT = 16

func (pself *Pipe) Ready() {
	regionRect := pself.PipeBodySprite.GetRegionRect()
	regionRect.Size = gd.NewVector2(32, pself.Height-TOP_PIPE_HEIGHT)
	pself.PipeBodySprite.SetRegionRect(regionRect)
	pself.PipeBodySprite.AsNode2D().SetPosition(gd.NewVector2(0, pself.Height/2))
	var shape = gd.Create(pself.Temporary, new(gd.RectangleShape2D))
	shape.SetSize(gd.NewVector2(32, pself.Height))
	pself.CollisionShape2D.SetShape(shape.AsShape2D())
	pself.CollisionShape2D.AsNode2D().SetPosition(gd.NewVector2(0, pself.Height/2-TOP_PIPE_HEIGHT/2))
}
