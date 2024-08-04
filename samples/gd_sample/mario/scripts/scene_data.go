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

func (pself PlayerMode) String() string {
	switch pself {
	case SMALL:
		return "small"
	case BIG:
		return "big"
	case SHOOTING:
		return "shooting"
	default:
		return ""
	}
}

type SceneData struct {
	gd.Class[SceneData, gd.Node] `gd:"SceneData"`

	ReturnPoint gd.Vector2
	PlayerMode  PlayerMode
	Points      gd.Int
	Coins       gd.Int
}
