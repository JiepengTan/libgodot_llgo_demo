package scripts

import (
	"grow.graphics/gd"
)

type Pipe struct {
	gd.Class[Pipe, gd.Area2D] `gd:"Enemy"`
}

type Enemy struct {
	gd.Class[Enemy, gd.Area2D] `gd:"Enemy"`

	HorizontalSpeed  float64
	VerticalSpeed    float64
	RayCast2D        gd.RayCast2D        `gd:"RayCast2D"`
	AnimatedSprite2D gd.AnimatedSprite2D `gd:"AnimatedSprite2D"`
	PointsLabelScene gd.PackedScene      `gd:"res://Scenes/points_label.tscn"`
}

func (pself *Enemy) Ready() {
	pself.Super().AsNode().SetProcess(false)
}
func (pself *Enemy) SetPosition(position gd.Vector2) {
	pself.Super().AsNode2D().SetPosition(position)
}
func (pself *Enemy) Position() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *Enemy) GetPosition() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *Enemy) GetTree(ctx gd.Lifetime) gd.SceneTree {
	return pself.Super().AsNode().GetTree(pself.Temporary)
}
func (pself *Enemy) String(s string) gd.String {
	return pself.Temporary.String(s)
}
func (pself *Enemy) NodePath(s string) gd.NodePath {
	return pself.String(s).NodePath(pself.Temporary)
}
func (pself *Enemy) StringName(s string) gd.StringName {
	return pself.Temporary.StringName(s)
}
func (pself *Enemy) DoQueueFree() {
	pself.Super().AsNode().QueueFree()
}
func (pself *Enemy) Variant(v any) gd.Variant {
	return pself.Temporary.Variant(v)
}
func (pself *Enemy) TweenVector2(tween gd.Tween, property string, final_val gd.Vector2, duration float32) gd.PropertyTweener {
	return tween.TweenProperty(
		pself.Temporary, pself.AsObject(),
		pself.Temporary.String(property).NodePath(pself.Temporary),
		pself.Temporary.Variant(final_val),
		gd.Float(duration))
}

func (pself *Enemy) Process(delta float64) {
	pself.SetPosition(pself.GetPosition().Add(gd.NewVector2(-delta*pself.HorizontalSpeed, 0)))

	if !pself.RayCast2D.IsColliding() {
		pself.SetPosition(pself.GetPosition().Add(gd.NewVector2(0, delta*pself.VerticalSpeed)))
	}
}

func (pself *Enemy) Die() {
	pself.HorizontalSpeed = 0
	pself.VerticalSpeed = 0
	pself.AnimatedSprite2D.Play(pself.Temporary.StringName("dead"), 1, false)
}

func (pself *Enemy) DieFromHit() {
	pself.Super().AsCollisionObject2D().SetCollisionLayerValue(3, false)
	pself.Super().AsCollisionObject2D().SetCollisionMaskValue(3, false)
	pself.Super().AsNode2D().SetRotationDegrees(180)
	pself.HorizontalSpeed = 0
	pself.VerticalSpeed = 0

	dieTween := pself.GetTree(pself.Temporary).CreateTween(pself.Temporary)
	pself.TweenVector2(dieTween, "position", pself.Position().Add(gd.NewVector2(0, -25)), 0.2)
	pself.TweenVector2(dieTween.Chain(pself.Temporary), "position", pself.Position().Add(gd.NewVector2(0, 500)), 4)

	pointsLabel, _ := gd.As[*gd.Node2D](pself.Temporary, pself.PointsLabelScene.Instantiate(pself.Temporary, gd.PackedSceneGenEditState(0)))
	pointsLabel.SetPosition(pself.Position().Add(gd.NewVector2(-20, -20)))
	pself.GetTree(pself.Temporary).GetRoot(pself.Temporary).AsNode().AddChild(pointsLabel.AsNode(), true, gd.NodeInternalMode(0))

	levelManagerNode := pself.GetTree(pself.Temporary).GetFirstNodeInGroup(pself.Temporary, pself.Temporary.StringName("level_manager"))
	if levelManager, ok := gd.As[*LevelManager](pself.Temporary, levelManagerNode); ok {
		levelManager.OnPointsScored(100)
	}
}

func (pself *Enemy) OnAreaEntered(area gd.Area2D) {
	if koopa, ok := gd.As[*Koopa](pself.Temporary, area); ok {
		if koopa.InAShell && koopa.HorizontalSpeed != 0 {
			pself.DieFromHit()
		}
	}
}

func (pself *Enemy) OnVisibleOnScreenNotifier2DScreenExited() {
	pself.DoQueueFree()
}

func (pself *Enemy) OnBodyEntered(body gd.Node) {
	if _, ok := gd.As[*Pipe](pself.Temporary, body); ok {
		pself.HorizontalSpeed *= -1
	}
}

func (pself *Enemy) OnVisibleOnScreenNotifier2DScreenEntered() {
	pself.Super().AsNode().SetProcess(true)
}
