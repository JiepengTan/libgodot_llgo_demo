package demo

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
	. "github.com/godot-go/godot-go/pkg/core"
	. "github.com/godot-go/godot-go/pkg/ffi"
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

func RegisterClassHUD() {
	ClassDBRegisterClass[*HUD](&HUD{}, []GDExtensionPropertyInfo{}, nil, func(t GDClass) {
		ClassDBAutoRegister[*HUD](t)
		// signals
		ClassDBAddSignal(t, "start_game")
	})
}

type HUD struct {
	CanvasLayerImpl
	StartButton Button `godot:"StartButton"`
}

func (pself *HUD) getScoreLabel() Label {
	return GetNode[Label](pself, "ScoreLabel")
}

func (pself *HUD) getMessageLabel() Label {
	return GetNode[Label](pself, "MessageLabel")
}

func (pself *HUD) getMessageTimer() Timer {
	return GetNode[Timer](pself, "MessageTimer")
}

func (pself *HUD) getStartButton() Button {
	return GetNode[Button](pself, "StartButton")
}

func (pself *HUD) showMessage_StrExt(text string) {
	gameOverMessage := NewVariantGoString(text)
	defer gameOverMessage.Destroy()
	pself.ShowMessage(gameOverMessage)
}

func (pself *HUD) ShowMessage(text Variant) {
	// $MessageLabel.text = text
	pself.getMessageLabel().SetText_StrExt(text.ToGoString())
	// $MessageLabel.show()
	pself.getMessageLabel().Show()
	// $MessageTimer.start()
	pself.getMessageTimer().Start(-1)
}

func (pself *HUD) ShowGameOver() {
	// show_message("Game Over")
	pself.showMessage_StrExt("Game Over")
	// await $MessageTimer.timeout
	DelayCallTimer(pself, "show_game_over_await_message_timer_timeout", pself.getMessageTimer())
}

func (pself *HUD) ShowGameOverAwaitMessageTimerTimeout() {
	// $MessageLabel.text = "Dodge the\nCreeps"
	messageLabel := pself.getMessageLabel()
	messageLabel.SetText_StrExt("Dodge the\nCreeps")

	// $MessageLabel.show()
	messageLabel.Show()
	DelayCall(pself, "show_game_over_await_scene_tree_timer_timeout", 1)
}

func (pself *HUD) ShowGameOverAwaitSceneTreeTimerTimeout() {
	// $StartButton.show()
	pself.getStartButton().Show()
}

func (pself *HUD) UpdateScore(score Variant) {
	// $ScoreLabel.text = str(score)
	pself.getScoreLabel().SetText_StrExt(score.ToGoString())
}

func (pself *HUD) V_OnPressed_StartButton() {
	pself.getStartButton().Hide()
	pself.EmitSignal_StrExt("start_game")
}

func (pself *HUD) V_OnTimeout_MessageTimer() {
	// $MessageLabel.hide()
	pself.getMessageLabel().Hide()
}
