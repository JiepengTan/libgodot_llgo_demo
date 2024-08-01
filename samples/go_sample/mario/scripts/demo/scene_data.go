package demo

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
)

var SceneData = struct {
	ReturnPoint Vector2
	PlayerMode  PlayerMode
	Points      int
	Coins       int
}{
	ReturnPoint: Vector2{},         // Initialize with zero vector
	PlayerMode:  PLAYER_MODE_SMALL, // Initialize with PlayerMode1
	Points:      0,
	Coins:       0,
}
