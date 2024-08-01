package demo

import (
	"fmt"
	. "github.com.godot-go/godot-go/pkg/gdclassimpl"
	. "github.com/godot-go/godot-go/pkg/builtin"
)

func (pself *Player) InitFieldsInternal() {
	pself.run_speed_damping = 0.5
	pself.speed = 200.0
	pself.jump_velocity = -350
	pself.min_stomp_degree = 35
	pself.max_stomp_degree = 145
	pself.stomp_y_velocity = -150
	pself.camera_sync = nil
	pself.should_camera_sync = true
	pself.castle_path = nil
	pself.player_mode = PLAYER_MODE_SMALL
	pself.is_dead = false
	pself.is_on_path = false
}

func (pself *Player) RegisterClassDB() {
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_FLOAT, "run_speed_damping", "set_run_speed_damping", "get_run_speed_damping")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_FLOAT, "speed", "set_speed", "get_speed")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_FLOAT, "jump_velocity", "set_jump_velocity", "get_jump_velocity")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_FLOAT, "min_stomp_degree", "set_min_stomp_degree", "get_min_stomp_degree")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_FLOAT, "max_stomp_degree", "set_max_stomp_degree", "get_max_stomp_degree")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_FLOAT, "stomp_y_velocity", "set_stomp_y_velocity", "get_stomp_y_velocity")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_OBJECT, "camera_sync", "set_camera_sync", "get_camera_sync")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_BOOL, "should_camera_sync", "set_should_camera_sync", "get_should_camera_sync")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_OBJECT, "castle_path", "set_castle_path", "get_castle_path")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_BOOL, "is_dead", "set_is_dead", "get_is_dead")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_BOOL, "is_on_path", "set_is_on_path", "get_is_on_path")
	ClassDBAddProperty(t, GDEXTENSION_VARIANT_TYPE_VECTOR2, "custom_position", "set_custom_position", "get_custom_position")
}

// Getters and Setters
func (pself *Player) GetRunSpeedDamping() float32 {
	return pself.run_speed_damping
}

func (pself *Player) SetRunSpeedDamping(value float32) {
	pself.run_speed_damping = value
}

func (pself *Player) GetSpeed() float32 {
	return pself.speed
}

func (pself *Player) SetSpeed(value float32) {
	pself.speed = value
}

func (pself *Player) GetJumpVelocity() float32 {
	return pself.jump_velocity
}

func (pself *Player) SetJumpVelocity(value float32) {
	pself.jump_velocity = value
}

func (pself *Player) GetMinStompDegree() float32 {
	return pself.min_stomp_degree
}

func (pself *Player) SetMinStompDegree(value float32) {
	pself.min_stomp_degree = value
}

func (pself *Player) GetMaxStompDegree() float32 {
	return pself.max_stomp_degree
}

func (pself *Player) SetMaxStompDegree(value float32) {
	pself.max_stomp_degree = value
}

func (pself *Player) GetStompYVelocity() float32 {
	return pself.stomp_y_velocity
}

func (pself *Player) SetStompYVelocity(value float32) {
	pself.stomp_y_velocity = value
}

func (pself *Player) GetCameraSync() Camera2D {
	return pself.camera_sync
}

func (pself *Player) SetCameraSync(value Camera2D) {
	pself.camera_sync = value
}

func (pself *Player) GetShouldCameraSync() bool {
	return pself.should_camera_sync
}

func (pself *Player) SetShouldCameraSync(value bool) {
	pself.should_camera_sync = value
}

func (pself *Player) GetCastlePath() PathFollow2D {
	return pself.castle_path
}

func (pself *Player) SetCastlePath(value PathFollow2D) {
	pself.castle_path = value
}

func (pself *Player) GetIsDead() bool {
	return pself.is_dead
}

func (pself *Player) SetIsDead(value bool) {
	pself.is_dead = value
}

func (pself *Player) GetIsOnPath() bool {
	return pself.is_on_path
}

func (pself *Player) SetIsOnPath(value bool) {
	pself.is_on_path = value
}
