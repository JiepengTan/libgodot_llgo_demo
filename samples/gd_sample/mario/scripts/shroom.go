package scripts

import (
	"grow.graphics/gd"
)

type Shroom struct {
	gd.Class[Shroom, gd.Area2D] `gd:"Shroom"`

	HorizontalSpeed         float64         `gd:"horizontal_speed"`
	MaxVerticalSpeed        float64         `gd:"max_vertical_speed"`
	VerticalVelocityGain    float64         `gd:"vertical_velocity_gain"`
	ShapeCast2D             *gd.ShapeCast2D `gd:"ShapeCast2D"`
	AllowHorizontalMovement bool
	VerticalSpeed           float64
}

func (s *Shroom) _Ready() {
	spawnTween := gd.GetTree().CreateTween()
	spawnTween.TweenProperty(s, "position", s.GetPosition().Add(gd.NewVector2(0, -16)), 0.4)
	spawnTween.TweenCallback(func() { s.AllowHorizontalMovement = true })
}

func (s *Shroom) _Process(delta gd.Float) {
	if s.AllowHorizontalMovement {
		s.SetPositionX(s.GetPosition().X + float64(delta)*s.HorizontalSpeed)
	}

	if !s.ShapeCast2D.IsColliding() && s.AllowHorizontalMovement {
		s.VerticalSpeed = gd.Lerp(s.VerticalSpeed, s.MaxVerticalSpeed, s.VerticalVelocityGain)
		s.SetPositionY(s.GetPosition().Y + float64(delta)*s.VerticalSpeed)
	} else {
		s.VerticalSpeed = 0
	}
}
