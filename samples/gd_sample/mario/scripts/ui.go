package scripts

import (
	"fmt"

	"grow.graphics/gd"
)

type UI struct {
	gd.Class[UI, gd.CanvasLayer] `gd:"UI"`

	CenterContainer gd.CenterContainer `gd:"MarginContainer/CenterContainer"`
	ScoreLabel      gd.Label           `gd:"MarginContainer/HBoxContainer/ScoreLabel"`
	CoinsLabel      gd.Label           `gd:"MarginContainer/HBoxContainer/CoinsLabel"`
	Coins           gd.Int
}

func (pself *UI) SetPoints(points gd.Int) {
	pself.ScoreLabel.SetText(pself.Temporary.String(fmt.Sprintf("SCORE: " + fmt.Sprint(points))))
}

func (pself *UI) SetCoins(coins gd.Int) {
	pself.CoinsLabel.SetText(pself.Temporary.String(fmt.Sprintf("COINS: " + fmt.Sprint(coins))))
}

func (pself *UI) OnFinish() {
	pself.CenterContainer.AsCanvasItem().SetVisible(true)
}
