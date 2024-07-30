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

func (c *HUD) getScoreLabel() Label {
	return GetNode[Label](c, "ScoreLabel")
}

func (c *HUD) getMessageLabel() Label {
	return GetNode[Label](c, "MessageLabel")
}

func (c *HUD) getMessageTimer() Timer {
	return GetNode[Timer](c, "MessageTimer")
}

func (c *HUD) getStartButton() Button {
	return GetNode[Button](c, "StartButton")
}

func (c *HUD) showMessage_StrExt(text string) {
	gameOverMessage := NewVariantGoString(text)
	defer gameOverMessage.Destroy()
	c.ShowMessage(gameOverMessage)
}

func (c *HUD) ShowMessage(text Variant) {
	// $MessageLabel.text = text
	c.getMessageLabel().SetText_StrExt(text.ToGoString())
	// $MessageLabel.show()
	c.getMessageLabel().Show()
	// $MessageTimer.start()
	c.getMessageTimer().Start(-1)
}

func (c *HUD) ShowGameOver() {
	// show_message("Game Over")
	c.showMessage_StrExt("Game Over")
	// await $MessageTimer.timeout
	DelayCallTimer(c, "show_game_over_await_message_timer_timeout", c.getMessageTimer())
}

func (c *HUD) ShowGameOverAwaitMessageTimerTimeout() {
	// $MessageLabel.text = "Dodge the\nCreeps"
	messageLabel := c.getMessageLabel()
	messageLabel.SetText_StrExt("Dodge the\nCreeps")

	// $MessageLabel.show()
	messageLabel.Show()
	DelayCall(c, "show_game_over_await_scene_tree_timer_timeout", 1)
}

func (c *HUD) ShowGameOverAwaitSceneTreeTimerTimeout() {
	// $StartButton.show()
	c.getStartButton().Show()
}

func (c *HUD) UpdateScore(score Variant) {
	// $ScoreLabel.text = str(score)
	c.getScoreLabel().SetText_StrExt(score.ToGoString())
}

func (c *HUD) V_OnPressed_StartButton() {
	c.getStartButton().Hide()
	c.EmitSignal_StrExt("start_game")
}

func (c *HUD) V_OnTimeout_MessageTimer() {
	// $MessageLabel.hide()
	c.getMessageLabel().Hide()
}
