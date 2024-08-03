package scripts

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdnative"
)

type Koopa struct {
	Enemy

	InAShell                         bool
	SlideSpeed                       float64             `gd:"slide_speed"`
	CollisionShape2D                 gd.CollisionShape2D `gd:"CollisionShape2D"`
	KoopaFullCollisionShape          gd.Shape2D          `gd:"res://Resources/CollisionShapes/koopa_full_collision_shape.tres"`
	KoopaShellCollisionShape         gd.Shape2D          `gd:"res://Resources/CollisionShapes/koopa_shell_collision_shape.tres"`
	KoopaShellCollisionShapePosition gdnative.Vector2    `gd:"KoopaShellCollisionShapePosition"`
}

func (pself *Koopa) Ready() {
	pself.CollisionShape2D.SetShape(pself.KoopaFullCollisionShape)
}

func (pself *Koopa) Die() {
	if !pself.InAShell {
		pself.Enemy.Die()
	}

	pself.CollisionShape2D.SetDeferred("shape", pself.KoopaShellCollisionShape)
	pself.CollisionShape2D.SetDeferred("position", pself.KoopaShellCollisionShapePosition)
	pself.InAShell = true
}

func (pself *Koopa) OnStomp(playerPosition gdnative.Vector2) {
	pself.SetCollisionMaskValue(1, false)
	pself.SetCollisionLayerValue(3, false)
	pself.SetCollisionLayerValue(4, true)

	var movementDirection float64 = 1
	if playerPosition.X() <= pself.GlobalPosition().X() {
		movementDirection = -1
	}
	pself.HorizontalSpeed = -movementDirection * pself.SlideSpeed
}
