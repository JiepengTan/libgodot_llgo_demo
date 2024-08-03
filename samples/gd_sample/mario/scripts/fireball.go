package scripts

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdnative"
)

type Fireball struct {
	gd.Class[Fireball, gd.Area2D] `gd:"Fireball"`

	RayCast2D                 gd.RayCast2D  `gd:"RayCast2D"`
	HorizontalSpeed           float64       `gd:"horizontal_speed"`
	VerticalSpeed             float64       `gd:"vertical_speed"`
	Amplitude                 float64       `gd:"amplitude"`
	IsMovingUp                bool
	Direction                 float64
	VerticalMovementStartPos  gdnative.Vector2
}

func (pself *Fireball) Process(delta float64) {
	pself.SetPosition(pself.Position().Add(gdnative.NewVector2(delta*pself.HorizontalSpeed*pself.Direction, 0)))

	if pself.RayCast2D.IsColliding() {
		pself.IsMovingUp = true
		pself.VerticalMovementStartPos = pself.Position()
	}

	if pself.IsMovingUp {
		pself.SetPosition(pself.Position().Add(gdnative.NewVector2(0, -pself.VerticalSpeed*delta)))
		if pself.VerticalMovementStartPos.Y() - pself.Amplitude >= pself.Position().Y() {
			pself.IsMovingUp = false
		}
	} else {
		pself.SetPosition(pself.Position().Add(gdnative.NewVector2(0, pself.VerticalSpeed*delta)))
	}
}

func (pself *Fireball) OnAreaEntered(area gd.Area2D) {
	if enemy, ok := area.(*Enemy); ok {
		enemy.DieFromHit()
	}
	pself.QueueFree()
}
