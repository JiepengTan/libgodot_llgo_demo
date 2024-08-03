package scripts

import (
	"grow.graphics/gd"
)

type Label struct {
	gd.Class[Label, gd.Label] `gd:"Label"`
}

func (l *Label) _Ready() {
	labelTween := gd.GetTree().CreateTween()
	labelTween.TweenProperty(l, "position", l.GetPosition().Add(gd.NewVector2(0, -10)), 0.4)
	labelTween.TweenCallback(l.QueueFree)
}
