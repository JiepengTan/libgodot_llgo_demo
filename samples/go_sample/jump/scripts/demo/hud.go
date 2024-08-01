package demo

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

// @autobind signal "start_game"
type HUD struct {
	CanvasLayerImpl
	ScoreLabel   Label  `gdbind:"ScoreLabel"`
	MessageLabel Label  `gdbind:"MessageLabel"`
	MessageTimer Timer  `gdbind:"MessageTimer"`
	StartButton  Button `gdbind:"StartButton"`
}

func (pself *HUD) GetSignals() []string {
	return []string{"start_game"}
}

func (pself *HUD) showMessage_StrExt(text string) {
	gameOverMessage := NewVariantGoString(text)
	defer gameOverMessage.Destroy()
	pself.ShowMessage(gameOverMessage)
}

func (pself *HUD) ShowMessage(text Variant) {
	// $MessageLabel.text = text
	pself.MessageLabel.SetText_StrExt(text.ToGoString())
	// $MessageLabel.show()
	pself.MessageLabel.Show()
	// $MessageTimer.start()
	pself.MessageTimer.Start(-1)
}

func (pself *HUD) ShowGameOver() {
	// show_message("Game Over")
	pself.showMessage_StrExt("Game Over")
	// await $MessageTimer.timeout
	DelayCallTimer(pself, "show_game_over_await_message_timer_timeout", pself.MessageTimer)
}

func (pself *HUD) ShowGameOverAwaitMessageTimerTimeout() {
	// $MessageLabel.text = "Dodge the\nCreeps"
	pself.MessageLabel.SetText_StrExt("Dodge the\nCreeps")
	// $MessageLabel.show()
	pself.MessageLabel.Show()
	DelayCall(pself, "show_game_over_await_scene_tree_timer_timeout", 1)
}

func (pself *HUD) ShowGameOverAwaitSceneTreeTimerTimeout() {
	// $StartButton.show()
	pself.StartButton.Show()
}

func (pself *HUD) UpdateScore(score Variant) {
	// $ScoreLabel.text = str(score)
	pself.ScoreLabel.SetText_StrExt(score.ToGoString())
}

func (pself *HUD) V_on_StartButton_pressed() {
	pself.StartButton.Hide()
	pself.EmitSignal_StrExt("start_game")
}

func (pself *HUD) V_on_MessageTimer_timeout() {
	// $MessageLabel.hide()
	pself.MessageLabel.Hide()
}
