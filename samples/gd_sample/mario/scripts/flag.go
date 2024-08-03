package scripts

import (
	"grow.graphics/gd"
)

type PoleArea struct {
	gd.Class[PoleArea, gd.Area2D]
}

func (pself *PoleArea) OnBodyEntered(body gd.Node) {
	if player, ok := gd.As[*Player](pself.Temporary, body); ok {
		player.OnPoleHit()
	}
}
