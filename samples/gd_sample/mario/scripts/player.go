package scripts

import (
	"grow.graphics/gd"
)

// TODO(tanjp): Add fields and methods.
type Player struct {
	gd.Class[Player, gd.Area2D] `gd:"Player"`
}

func (pself *Player) Die() {
}

func (pself *Player) OnPoleHit() {
}

func (pself *Player) HandlePipeConnectorEntranceCollision() {
}
