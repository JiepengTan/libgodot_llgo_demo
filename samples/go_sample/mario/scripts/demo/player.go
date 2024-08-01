package demo

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

// @gdclass Player
type Player struct {
	CharacterBody2DImpl
	animated_sprite_2d           AnimatedSprite2D `gdbind:"AnimatedSprite2D"`
	area_2d                      Area2D           `gdbind:"Area2D"`
	area_collision_shape         CollisionShape2D `gdbind:"Area2D/AreaCollisionShape"`
	body_collision_shape         CollisionShape2D `gdbind:"BodyCollisionShape"`
	shooting_point               Marker2D         `gdbind:"ShootingPoint"`
	slide_down_finished_position Marker2D         `gdbind:"../slide_down_finished_position"` // TODO
	land_down_marker             Marker2D         `gdbind:"../LandDownMarker"`
	//@gdgroup Locomotion
	run_speed_damping float32 `gdexport:"0.5"`
	speed             float32 `gdexport:"200.0"`
	jump_velocity     float32 `gdexport:"-350"`
	//@gdgroup Stomping enemies
	min_stomp_degree float32 `gdexport:"35"`
	max_stomp_degree float32 `gdexport:"145"`
	stomp_y_velocity float32 `gdexport:"-150"`
	//@gdgroup Camera sync
	camera_sync        Camera2D `gdexport:"nil"`
	should_camera_sync bool     `gdexport:"true"`

	//@gdgroup
	castle_path    PathFollow2D `gdexport:"nil"`
	player_mode    PlayerMode
	is_dead        bool
	is_on_path     bool
	input          Input
	projectSetting ProjectSettings
}

const PIPE_ENTER_THRESHOLD = 10

type PlayerMode int

const (
	PLAYER_MODE_SMALL PlayerMode = iota
	PLAYER_MODE_BIG
	PLAYER_MODE_SHOOTING
)

func (p PlayerMode) String() string {
	switch p {
	case PLAYER_MODE_SMALL:
		return "SMALL"
	case PLAYER_MODE_BIG:
		return "BIG"
	case PLAYER_MODE_SHOOTING:
		return "SHOOTING"
	default:
		return "UNKNOWN"
	}
}

func (pself *Player) GetGravity() float32 {
	val := pself.projectSetting.GetSettingWithOverride_StrExt("physics/2d/default_gravity")
	return val.ToFloat32()
}

func (pself *Player) PointsLabelScene() PackedScene {
	return pself.LoadResource("res://Scenes/points_label.tscn")
}

func (pself *Player) SmallMarioCollisionShape() PackedScene {
	return pself.LoadResource("res://Resources/CollisionShapes/small_mario_collision_shape.tres")
}

func (pself *Player) BigMarioCollisionShape() PackedScene {
	return pself.LoadResource("res://Resources/CollisionShapes/big_mario_collision_shape.tres")
}

func (pself *Player) FireballScene() PackedScene {
	return pself.LoadResource("res://Scenes/fireball.tscn")
}

func (pself *Player) LoadResource(path string) PackedScene {
	// TODO(tanjp) implement GD static methods
	return nil
}

func (pself *Player) V_ready() {
	pself.input = GetInputSingleton()
	pself.projectSetting = GetProjectSettingSingleton()

	if SceneData.ReturnPoint.Equal_Vector2(ZeroVector2()) {
		pself.SetGlobalPosition(SceneData.ReturnPoint)
	}
}

func (pself *Player) V_physics_process(delta float32) {
	rect := pself.camera_sync.GetViewportRect()
	camera_left_bound := pself.camera_sync.GetGlobalPosition().GetX() - rect.MemberGetsize().GetX()/2/pself.camera_sync.GetZoom().GetX()
	velocity := pself.GetVelocity()
	// Apply gravity
	if !pself.IsOnFloor() {
		velocity.SetY(velocity.GetY() + pself.GetGravity()*delta)
	}

	if pself.GetGlobalPosition().GetX() < camera_left_bound+8 && velocity.GetX() < 0 {
		pself.SetVelocity(ZeroVector2())
		return
	}

	// Handle jumps
	if pself.input.IsActionJustPressed_StrExt("jump", false) && pself.IsOnFloor() {
		velocity.SetY(pself.jump_velocity)
	}

	if pself.input.IsActionJustReleased_StrExt("jump", false) && velocity.GetY() < 0 {
		velocity.SetY(velocity.GetY() * 0.5)
	}
	pself.SetVelocity(velocity)
}

func (pself *Player) V_process(delta float32) {
	if pself.GetGlobalPosition().GetX() > pself.camera_sync.GetGlobalPosition().GetX() && pself.should_camera_sync {
		pos := pself.camera_sync.GetGlobalPosition()
		pos.SetX(pself.GetGlobalPosition().GetX())
		pself.camera_sync.SetGlobalPosition(pos)
	}

}
