package scripts

import (
	"grow.graphics/gd"
)

type LevelManager struct {
	gd.Class[LevelManager, gd.Node] `gd:"LevelManager"`
	UI                              UI        `gd:"../UI"`
	Player                          Player    `gd:"../Player"`
	SceneData                       SceneData `gd:"../SceneData"`

	Points gd.Int
	Coins  gd.Int
}

func (pself *LevelManager) Ready() {
	pself.Points = 0
	pself.Coins = 0
	pself.Player.CastleEntered.Connect(pself.Temporary.Callable(pself.UI.OnButtonPressed), 0)
	pself.Player.PointsScored.Connect(pself.Temporary.Callable(pself.OnPointsScored), 0)
	if pself.SceneData.Points != 0 {
		pself.UI.SetPoints(pself.SceneData.Points)
		pself.Points = pself.SceneData.Points
	}
	if pself.SceneData.Coins != 0 {
		pself.UI.SetCoins(pself.SceneData.Coins)
		pself.Coins = pself.SceneData.Coins
	}
}

func (pself *LevelManager) OnPointsScored(point gd.Int) {
	pself.Points += point
	pself.UI.SetPoints(pself.Points)
}

func (pself *LevelManager) OnCoinCollected() {
	pself.Coins += 1
	pself.UI.SetCoins(pself.Coins)
}
