package scripts

import (
	"grow.graphics/gd"
)

type PipeConnector struct {
	gd.Class[PipeConnector, gd.StaticBody2D] `gd:"PipeConnector"`
	SceneData                                SceneData `gd:"../SceneData"`
	ReturnPoint                              gd.Vector2
}

func (pself *PipeConnector) OnEntranceBodyEntered(body gd.Node2D) {
	pself.SceneData.ReturnPoint = pself.ReturnPoint
	player, ok := body.(*Player)
	if ok {
		player.HandlePipeConnectorEntranceCollision()
	}
}
