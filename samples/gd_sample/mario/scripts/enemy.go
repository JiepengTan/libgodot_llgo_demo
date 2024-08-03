package scripts

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdnative"
)

type Enemy struct {
	gd.Class[Enemy, gd.Area2D] `gd:"Enemy"`

	HorizontalSpeed   float64
	VerticalSpeed     float64
	RayCast2D         gd.RayCast2D         `gd:"RayCast2D"`
	AnimatedSprite2D  gd.AnimatedSprite2D  `gd:"AnimatedSprite2D"`
	PointsLabelScene  gd.PackedScene       `gd:"res://Scenes/points_label.tscn"`
}

func (pself *Enemy) Ready() {
	pself.SetProcess(false)
}

func (pself *Enemy) Process(delta float64) {
	pself.SetPosition(pself.Position().Add(gdnative.NewVector2(-delta*pself.HorizontalSpeed, 0)))

	if !pself.RayCast2D.IsColliding() {
		pself.SetPosition(pself.Position().Add(gdnative.NewVector2(0, delta*pself.VerticalSpeed)))
	}
}

func (pself *Enemy) Die() {
	pself.HorizontalSpeed = 0
	pself.VerticalSpeed = 0
	pself.AnimatedSprite2D.Play("dead")
}

func (pself *Enemy) DieFromHit() {
	pself.SetCollisionLayerValue(3, false)
	pself.SetCollisionMaskValue(3, false)
	pself.SetRotationDegrees(180)
	pself.HorizontalSpeed = 0
	pself.VerticalSpeed = 0

	dieTween := pself.GetTree().CreateTween()
	dieTween.TweenProperty(pself, "position", pself.Position().Add(gdnative.NewVector2(0, -25)), 0.2)
	dieTween.Chain().TweenProperty(pself, "position", pself.Position().Add(gdnative.NewVector2(0, 500)), 4)

	pointsLabel := pself.PointsLabelScene.Instantiate().(gd.Node2D)
	pointsLabel.SetPosition(pself.Position().Add(gdnative.NewVector2(-20, -20)))
	pself.GetTree().GetRoot().AddChild(pointsLabel)

	levelManager := pself.GetTree().GetFirstNodeInGroup("level_manager").(*LevelManager)
	levelManager.OnPointsScored(100)
}

func (pself *Enemy) OnAreaEntered(area gd.Area2D) {
	if koopa, ok := area.(*Koopa); ok && koopa.InAShell() && koopa.HorizontalSpeed != 0 {
		pself.DieFromHit()
	}
}

func (pself *Enemy) OnVisibleOnScreenNotifier2DScreenExited() {
	pself.QueueFree()
}

func (pself *Enemy) OnBodyEntered(body gd.Node) {
	if _, ok := body.(*Pipe); ok {
		pself.HorizontalSpeed *= -1
	}
}

func (pself *Enemy) OnVisibleOnScreenNotifier2DScreenEntered() {
	pself.SetProcess(true)
}
