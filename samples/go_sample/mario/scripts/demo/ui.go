package demo

import (
	"fmt"
	. "github.com/godot-go/godot-go/pkg/builtin"
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

type UI struct {
	CanvasLayerImpl
	CenterContainer CenterContainer `gdbind:"MarginContainer/CenterContainer"`
	ScoreLabel      Label           `gdbind:"MarginContainer/HBoxContainer/ScoreLabel"`
	CoinsLabel      Label           `gdbind:"MarginContainer/HBoxContainer/CoinsLabel"`
}

func (pself *UI) SetScore(points int) {
	println("UI SetScore")
	pself.ScoreLabel.SetText_StrExt(fmt.Sprintf("SCORE: %d", points))
}

func (pself *UI) SetCoin(coins int) {
	println("UI SetCoin")
	pself.CoinsLabel.SetText_StrExt(fmt.Sprintf("COINS: %d", coins))
}

func (pself *UI) OnFinish() {
	println("UI OnFinish")
	pself.CenterContainer.SetVisible(true)
}
