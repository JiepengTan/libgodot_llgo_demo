package scripts

import (
	"grow.graphics/gd"
)

type Label struct {
	gd.Class[Label, gd.Label] `gd:"Label"`
}

func (pself *Label) Position() gd.Vector2 {
	return pself.Super().AsControl().GetPosition()
}
func (pself *Label) GetPosition() gd.Vector2 {
	return pself.Super().AsControl().GetPosition()
}
func (pself *Label) GetTree(ctx gd.Lifetime) gd.SceneTree {
	return pself.Super().AsNode().GetTree(pself.Temporary)
}

func (pself *Label) DoQueueFree() {
	pself.Super().AsNode().QueueFree()
}
func (pself *Label) TweenVector2(tween gd.Tween, property string, final_val gd.Vector2, duration float32) gd.PropertyTweener {
	return tween.TweenProperty(
		pself.Temporary, pself.AsObject(),
		pself.Temporary.String(property).NodePath(pself.Temporary),
		pself.Temporary.Variant(final_val),
		gd.Float(duration))
}
func (pself *Label) _Ready() {
	labelTween := pself.GetTree(pself.Temporary).CreateTween(pself.Temporary)
	pself.TweenVector2(labelTween, "position", pself.GetPosition().Add(gd.NewVector2(0, -10)), 0.4)

	labelTween.TweenCallback(pself.Temporary, pself.Temporary.Callable(pself.DoQueueFree))
}
