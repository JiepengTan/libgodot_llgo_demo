package scripts

import (
	"fmt"

	"grow.graphics/gd"
)

type UI struct {
	gd.Class[UI, gd.CanvasLayer] `gd:"GdUI"`

	CenterContainer gd.CenterContainer `gd:"MarginContainer/CenterContainer"`
	ScoreLabel      gd.Label           `gd:"MarginContainer/HBoxContainer/ScoreLabel"`
	CoinsLabel      gd.Label           `gd:"MarginContainer/HBoxContainer/CoinsLabel"`
}

func (pself *UI) SetScore(points int) {
	pself.ScoreLabel.SetText(pself.Temporary.String(fmt.Sprintf("SCORE: %d", points)))
}

func (pself *UI) SetCoins(coins int) {
	pself.CoinsLabel.SetText(pself.Temporary.String(fmt.Sprintf("COINS: %d", coins)))
}

func (pself *UI) OnFinish() {
	pself.CenterContainer.AsCanvasItem().SetVisible(true)
}
