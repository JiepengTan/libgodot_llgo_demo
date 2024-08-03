package scripts

import (
	"grow.graphics/gd"
)

type CustomEnemy struct {
	Enemy
}

func (pself *CustomEnemy) Die() {
	pself.Enemy.Die()
	pself.SetCollisionLayerValue(3, false)
	pself.SetCollisionMaskValue(1, false)
	timer := pself.GetTree().CreateTimer(1.5)
	timer.Connect("timeout", pself.QueueFree)
}
