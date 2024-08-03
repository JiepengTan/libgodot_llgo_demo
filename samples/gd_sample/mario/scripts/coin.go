package scripts

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdnative"
)

type Coin struct {
	gd.Class[Coin, gd.AnimatedSprite2D]
}

func (pself *Coin) Ready() {
	coinTween := pself.GetTree().CreateTween()
	coinTween.TweenProperty(pself, "position", pself.Position().Add(gdnative.NewVector2(0, -40)), 0.3)
	coinTween.Chain().TweenProperty(pself, "position", pself.Position(), 0.3)
	coinTween.TweenCallback(pself.QueueFree)
}
