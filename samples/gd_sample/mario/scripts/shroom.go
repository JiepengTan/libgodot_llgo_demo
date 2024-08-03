package scripts

import (
	"grow.graphics/gd"
)

type Shroom struct {
	gd.Class[Shroom, gd.Area2D] `gd:"Shroom"`

	HorizontalSpeed         gd.Float        `gd:"horizontal_speed"`
	MaxVerticalSpeed        gd.Float        `gd:"max_vertical_speed"`
	VerticalVelocityGain    gd.Float        `gd:"vertical_velocity_gain"`
	ShapeCast2D             *gd.ShapeCast2D `gd:"ShapeCast2D"`
	AllowHorizontalMovement bool
	VerticalSpeed           gd.Float
}

func (pself *Shroom) Position() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *Shroom) GetPosition() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *Shroom) GetTree(ctx gd.Lifetime) gd.SceneTree {
	return pself.Super().AsNode().GetTree(pself.Temporary)
}
func (pself *Shroom) TweenVector2(tween gd.Tween, property string, final_val gd.Vector2, duration float32) gd.PropertyTweener {
	return tween.TweenProperty(
		pself.Temporary, pself.AsObject(),
		pself.Temporary.String(property).NodePath(pself.Temporary),
		pself.Temporary.Variant(final_val),
		gd.Float(duration))
}
func (pself *Shroom) _Ready() {

	spawnTween := pself.GetTree(pself.Temporary).CreateTween(pself.Temporary)
	pself.TweenVector2(spawnTween, "position", pself.GetPosition().Add(gd.NewVector2(0, -16)), 0.4)
	spawnTween.TweenCallback(pself.Temporary, pself.Temporary.Callable(func() { pself.AllowHorizontalMovement = true }))

}
func (pself *Shroom) SetPosition(position gd.Vector2) {
	pself.Super().AsNode2D().SetPosition(position)
}
func (pself *Shroom) SetPositionX(val float64) {
	pself.SetPosition(gd.NewVector2(val, gd.Float(pself.GetPosition()[gd.Y])))
}
func (pself *Shroom) SetPositionY(val float64) {
	pself.SetPosition(gd.NewVector2(gd.Float(pself.GetPosition()[gd.X]), val))
}
func lerp(a, b float64, t float64) float64 {
	return a + (b-a)*t
}
func (pself *Shroom) _Process(delta gd.Float) {
	if pself.AllowHorizontalMovement {
		val := gd.Float(pself.GetPosition()[gd.X]) + delta*pself.HorizontalSpeed
		pself.SetPositionX(val)
	}

	if !pself.ShapeCast2D.IsColliding() && pself.AllowHorizontalMovement {
		pself.VerticalSpeed = lerp(pself.VerticalSpeed, pself.MaxVerticalSpeed, pself.VerticalVelocityGain)
		val := gd.Float(pself.GetPosition()[gd.Y]) + float64(delta)*pself.VerticalSpeed
		pself.SetPositionY(val)
	} else {
		pself.VerticalSpeed = 0
	}
}
