package scripts

import (
	"grow.graphics/gd"
)

type Brick struct {
	Block

	GPUParticles2D gd.GPUParticles2D `gd:"GPUParticles2D"`
	Sprite2D       gd.Sprite2D       `gd:"Sprite2D"`
}

func (pself *Brick) Bump(playerMode Player.PlayerMode) {
	if playerMode == PlayerModeSmall {
		pself.Block.Bump(playerMode)
	} else if !pself.GPUParticles2D.IsEmitting() {
		pself.SetCollisionLayerValue(5, false)
		pself.GPUParticles2D.SetEmitting(true)
		pself.Sprite2D.SetVisible(false)
		pself.Block.CheckForEnemyCollision()
	}
}

func (pself *Brick) OnGPUParticles2DFinished() {
	pself.QueueFree()
}
