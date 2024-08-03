package scripts

import (
	"grow.graphics/gd"
)

type Brick struct {
	Block

	GPUParticles2D gd.GPUParticles2D `gd:"GPUParticles2D"`
	Sprite2D       gd.Sprite2D       `gd:"Sprite2D"`
}

func (pself *Brick) Bump(playerMode PlayerMode) {
	if playerMode == SMALL {
		pself.Block.Bump(playerMode)
	} else if !pself.GPUParticles2D.IsEmitting() {
		pself.Super().AsCollisionObject2D().SetCollisionLayerValue(5, false)
		pself.GPUParticles2D.SetEmitting(true)
		pself.Sprite2D.AsCanvasItem().SetVisible(false)
		pself.Block.CheckForEnemyCollision()
	}
}

func (pself *Brick) OnGPUParticles2DFinished() {
	pself.Super().AsNode().QueueFree()
}
