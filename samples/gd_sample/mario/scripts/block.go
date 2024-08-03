package scripts

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdnative"
)

type Block struct {
	gd.Class[Block, gd.StaticBody2D] `gd:"Block"`

	RayCast2D gd.RayCast2D `gd:"RayCast2D"`
}

func (pself *Block) Bump(playerMode gd.Ref) {
	bumpTween := pself.GetTree().CreateTween()
	bumpTween.TweenProperty(pself, "position", pself.Position().Add(gdnative.NewVector2(0, -5)), 0.12)
	bumpTween.Chain().TweenProperty(pself, "position", pself.Position(), 0.12)
	pself.CheckForEnemyCollision()
}

func (pself *Block) CheckForEnemyCollision() {
	if pself.RayCast2D.IsColliding() {
		collider := pself.RayCast2D.GetCollider()
		if enemy, ok := collider.(*Enemy); ok {
			enemy.DieFromHit()
		}
	}
}
