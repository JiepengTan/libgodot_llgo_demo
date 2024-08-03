package demo

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

type LevelManager struct {
	NodeImpl
	Coins  int
	Points int
}

func (pself *LevelManager) getUI() *UI {
	return ObjectCastToGeneric[*UI](pself.GetTree(pself.Temporary).GetRoot().GetNode_StrExt("UI"))
}

func (pself *LevelManager) V_ready() {
	ui := pself.getUI()
	if SceneData.Points != 0 {
		ui.SetScore(SceneData.Points)
		pself.Points = SceneData.Points
	}
	if SceneData.Coins != 0 {
		pself.Coins = SceneData.Coins
		ui.SetCoin(SceneData.Coins)
	}
}

func (pself *LevelManager) OnCastleEntered() {
	ui := pself.getUI()
	ui.OnFinish()
}

func (pself *LevelManager) OnPointsScored(point int) {
	pself.Points += point
	ui := pself.getUI()
	ui.SetScore(pself.Points)
}

func (pself *LevelManager) OnCoinCollected() {
	pself.Coins += 1
	ui := pself.getUI()
	ui.SetCoin(pself.Coins)
}
