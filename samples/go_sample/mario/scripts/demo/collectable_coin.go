package demo

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

type CollectableCoin struct {
	Area2DImpl
}

func (pself *CollectableCoin) V_on_body_entered(body Node2D) {
	pself.QueueFree()
	node := pself.GetTree().GetFirstNodeInGroup_StrExt("level_manager")
	if node != nil {
		println("V_on_body_enter find level_manager succ")
	} else {
		println("V_on_body_enter can not find level_manager ")
	}
	println("V_on_body_entered")
}
func (pself *CollectableCoin) V_ready() {
	println("CollectableCoin V_ready")
}
