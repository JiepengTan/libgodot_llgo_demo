package scripts

import (
	"grow.graphics/gd"
)

type CollectableCoin struct {
	gd.Class[CollectableCoin, gd.Area2D] `gd:"CollectableCoin"`
}

func (pself *CollectableCoin) OnBodyEntered(body gd.Node) {
	if _, ok := body.(*Player); ok {
		pself.QueueFree()
		levelManager := pself.GetTree().GetFirstNodeInGroup("level_manager").(*LevelManager)
		levelManager.OnCoinCollected()
	}
}
