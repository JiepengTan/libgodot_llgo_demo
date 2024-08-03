package scripts

import (
	"grow.graphics/gd"
)

type Block struct {
	gd.Class[Block, gd.StaticBody2D] `gd:"Block"`

	RayCast2D gd.RayCast2D `gd:"RayCast2D"`
}

func (pself *Block) Bump(playerMode PlayerMode) {
	bumpTween := pself.Super().AsNode().GetTree(pself.Temporary).CreateTween(pself.Temporary)

	dstPos := pself.Super().AsNode2D().GetPosition().Add(gd.NewVector2(0, -5))
	bumpTween.TweenProperty(
		pself.Temporary, pself.AsObject(),
		pself.Temporary.String("position").NodePath(pself.Temporary),
		pself.Temporary.Variant(dstPos),
		gd.Float(0.12))

	srcPos := pself.Super().AsNode2D().GetPosition()
	bumpTween.Chain(pself.Temporary).TweenProperty(
		pself.Temporary, pself.AsObject(),
		pself.Temporary.String("position").NodePath(pself.Temporary),
		pself.Temporary.Variant(srcPos),
		gd.Float(0.12))
	pself.CheckForEnemyCollision()
}

func (pself *Block) CheckForEnemyCollision() {
	/* // TODO(tanjp)
	if pself.RayCast2D.IsColliding() {
		collider := pself.RayCast2D.GetCollider()
		if enemy, ok := gd.As[*Enemy](pself.Temporary, collider); ok {
			enemy.DieFromHit()
		}
	}
	*/
}
