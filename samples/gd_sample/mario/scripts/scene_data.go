package scripts

import (
	"grow.graphics/gd"
)

type PlayerMode int

const (
	SMALL PlayerMode = iota
	BIG
	SHOOTING
)

type SceneData struct {
	gd.Class[SceneData, gd.Node] `gd:"SceneData"`

	ReturnPoint gd.Vector2
	//PlayerMode  Player_PlayerMode
	Points gd.Int
	Coins  gd.Int
}
