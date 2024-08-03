package scripts

import (
	"grow.graphics/gd"
)

type Koopa struct {
	Enemy

	InAShell                         bool
	SlideSpeed                       float64             `gd:"slide_speed"`
	CollisionShape2D                 gd.CollisionShape2D `gd:"CollisionShape2D"`
	KoopaFullCollisionShape          gd.Shape2D          `gd:"res://Resources/CollisionShapes/koopa_full_collision_shape.tres"`
	KoopaShellCollisionShape         gd.Shape2D          `gd:"res://Resources/CollisionShapes/koopa_shell_collision_shape.tres"`
	KoopaShellCollisionShapePosition gd.Vector2          `gd:"KoopaShellCollisionShapePosition"`
}

func (pself *Koopa) Ready() {
	pself.CollisionShape2D.SetShape(pself.KoopaFullCollisionShape)
}

func (pself *Koopa) Die() {
	if !pself.InAShell {
		pself.Enemy.Die()
	}
	pself.CollisionShape2D.AsObject().SetDeferred(pself.StringName("shape"), pself.Variant(pself.KoopaShellCollisionShape))
	pself.CollisionShape2D.AsObject().SetDeferred(pself.StringName("position"), pself.Variant(pself.KoopaShellCollisionShapePosition))
	pself.InAShell = true
}

func (pself *Koopa) OnStomp(playerPosition gd.Vector2) {
	pself.Super().AsCollisionObject2D().SetCollisionMaskValue(1, false)
	pself.Super().AsCollisionObject2D().SetCollisionLayerValue(3, false)
	pself.Super().AsCollisionObject2D().SetCollisionLayerValue(4, true)

	var movementDirection float64 = 1
	if playerPosition.X() <= pself.Super().AsNode2D().GetGlobalPosition().X() {
		movementDirection = -1
	}
	pself.HorizontalSpeed = -movementDirection * pself.SlideSpeed
}
