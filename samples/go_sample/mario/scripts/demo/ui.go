package demo

import (
    "fmt"
    . "github.com/godot-go/godot-go/pkg/builtin"
    . "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

type UI struct {
    CanvasLayerImpl
    CenterContainer CenterContainer  `godot:"MarginContainer/CenterContainer"`
    ScoreLabel      Label `godot:"MarginContainer/HBoxContainer/ScoreLabel"`
    CoinsLabel      Label `godot:"MarginContainer/HBoxContainer/CoinsLabel"`
}

func (pself *UI) SetScore(points int) {	
	println("UI SetScore")
    pself.ScoreLabel.SetText_StrExt(fmt.Sprintf("SCORE: %d", points))
}

func (pself *UI) SetCoins(coins int) {
	println("UI SetCoins")
    pself.CoinsLabel.SetText_StrExt(fmt.Sprintf("COINS: %d", coins))
}

func (pself *UI) OnFinish() {
	println("UI OnFinish")
    pself.CenterContainer.SetVisible(true)
}
