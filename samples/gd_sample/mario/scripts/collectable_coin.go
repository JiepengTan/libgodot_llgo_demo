package scripts

import (
	"grow.graphics/gd"
)

type CollectableCoin struct {
	gd.Class[CollectableCoin, gd.Area2D] `gd:"CollectableCoin"`
}

func (pself *CollectableCoin) OnBodyEntered(body gd.Node) {
	if _, ok := gd.As[*Player](pself.Temporary, body); ok {
		pself.Super().AsNode().QueueFree()
		tree := pself.Super().AsNode().GetTree(pself.Temporary)
		levelManagerNode := tree.GetFirstNodeInGroup(pself.Temporary, pself.Temporary.StringName("level_manager"))
		if levelManager, ok := gd.As[*LevelManager](pself.Temporary, levelManagerNode); ok {
			levelManager.OnCoinCollected()
		}
	}
}
