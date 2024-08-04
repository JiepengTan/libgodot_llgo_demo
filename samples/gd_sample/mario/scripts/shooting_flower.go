package scripts

import (
	"grow.graphics/gd"
)

type ShootingFlower struct {
	gd.Class[Shroom, gd.Area2D] `gd:"Shroom"`
}

func (pself *ShootingFlower) Position() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *ShootingFlower) TweenVector2(tween gd.Tween, property string, final_val gd.Vector2, duration float32) gd.PropertyTweener {
	return tween.TweenProperty(
		pself.Temporary, pself.AsObject(),
		pself.Temporary.String(property).NodePath(pself.Temporary),
		pself.Temporary.Variant(final_val),
		gd.Float(duration))
}
func (pself *ShootingFlower) Ready() {
	coinTween := pself.Super().AsNode().GetTree(pself.Temporary).CreateTween(pself.Temporary)
	pself.TweenVector2(coinTween, "position", pself.Position().Add(gd.NewVector2(0, -16)), 0.4)
}
