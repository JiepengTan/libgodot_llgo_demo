package scripts

import (
	"grow.graphics/gd"
)

type FallDownArea struct {
	gd.Class[FallDownArea, gd.Area2D] `gd:"FallDownArea"`
}

func (pself *FallDownArea) OnBodyEntered(body gd.Node) {
	if player, ok := gd.As[*Player](pself.Temporary, body); ok {
		player.Die()
	}
}
