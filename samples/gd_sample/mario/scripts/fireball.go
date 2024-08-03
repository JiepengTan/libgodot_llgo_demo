package scripts

import (
	"grow.graphics/gd"
)

type Fireball struct {
	gd.Class[Fireball, gd.Area2D] `gd:"Fireball"`

	RayCast2D                gd.RayCast2D `gd:"RayCast2D"`
	HorizontalSpeed          float64      `gd:"horizontal_speed"`
	VerticalSpeed            float64      `gd:"vertical_speed"`
	Amplitude                float64      `gd:"amplitude"`
	IsMovingUp               bool
	Direction                float64
	VerticalMovementStartPos gd.Vector2
}

func (pself *Fireball) Position() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *Fireball) GetPosition() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *Fireball) GetTree(ctx gd.Lifetime) gd.SceneTree {
	return pself.Super().AsNode().GetTree(pself.Temporary)
}
func (pself *Fireball) SetPosition(position gd.Vector2) {
	pself.Super().AsNode2D().SetPosition(position)
}
func (pself *Fireball) SetPositionX(val float64) {
	pself.SetPosition(gd.NewVector2(val, gd.Float(pself.GetPosition()[gd.Y])))
}
func (pself *Fireball) SetPositionY(val float64) {
	pself.SetPosition(gd.NewVector2(gd.Float(pself.GetPosition()[gd.X]), val))
}
func (pself *Fireball) DoQueueFree() {
	pself.Super().AsNode().QueueFree()
}
func (pself *Fireball) Process(delta float64) {
	pself.SetPosition(pself.Position().Add(gd.NewVector2(delta*pself.HorizontalSpeed*pself.Direction, 0)))

	if pself.RayCast2D.IsColliding() {
		pself.IsMovingUp = true
		pself.VerticalMovementStartPos = pself.Position()
	}

	if pself.IsMovingUp {
		pself.SetPosition(pself.Position().Add(gd.NewVector2(0, -pself.VerticalSpeed*delta)))
		if pself.VerticalMovementStartPos.Y()-pself.Amplitude >= pself.Position().Y() {
			pself.IsMovingUp = false
		}
	} else {
		pself.SetPosition(pself.Position().Add(gd.NewVector2(0, pself.VerticalSpeed*delta)))
	}
}
func (pself *Fireball) OnAreaEntered(area gd.Area2D) {
	if enemy, ok := gd.As[*Enemy](pself.Temporary, area); ok {
		enemy.DieFromHit()
	}
	pself.DoQueueFree()
}
