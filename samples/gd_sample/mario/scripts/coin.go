package scripts

import (
	"grow.graphics/gd"
)

type Coin struct {
	gd.Class[Coin, gd.AnimatedSprite2D]
}

func (pself *Coin) Ready() {
	coinTween := pself.Super().AsNode().GetTree(pself.Temporary).CreateTween(pself.Temporary)

	dstPos := pself.Super().AsNode2D().GetPosition().Add(gd.NewVector2(0, -40))
	coinTween.TweenProperty(
		pself.Temporary, pself.AsObject(),
		pself.Temporary.String("position").NodePath(pself.Temporary),
		pself.Temporary.Variant(dstPos),
		gd.Float(0.3))

	srcPos := pself.Super().AsNode2D().GetPosition()
	coinTween.Chain(pself.Temporary).TweenProperty(
		pself.Temporary, pself.AsObject(),
		pself.Temporary.String("position").NodePath(pself.Temporary),
		pself.Temporary.Variant(srcPos),
		gd.Float(0.3))
	coinTween.TweenCallback(pself.Temporary, pself.Temporary.Callable(pself.DoQueueFree))
}

func (pself *Coin) DoQueueFree() {
	pself.Super().AsNode().QueueFree()
}
