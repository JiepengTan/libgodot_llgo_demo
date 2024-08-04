package scripts

import (
	"strings"
	"unsafe"

	"grow.graphics/gd"
)

type PlayerAnimatedSprite struct {
	gd.Class[PlayerAnimatedSprite, gd.AnimatedSprite2D] `gd:"PlayerAnimatedSprite"`

	FrameCount gd.Int
}

func (pself *PlayerAnimatedSprite) String(s string) gd.String {
	return pself.Temporary.String(s)
}
func (pself *PlayerAnimatedSprite) NodePath(s string) gd.NodePath {
	return pself.String(s).NodePath(pself.Temporary)
}
func (pself *PlayerAnimatedSprite) StringName(s string) gd.StringName {
	return pself.Temporary.StringName(s)
}
func (pself *PlayerAnimatedSprite) Variant(v any) gd.Variant {
	return pself.Temporary.Variant(v)
}
func (pself *PlayerAnimatedSprite) Play(name string) {
	pself.Super().Play(pself.StringName(name), 1, false)
}
func Sign(value gd.Float) int {
	if value > 0 {
		return 1
	}
	if value < 0 {
		return -1
	}
	return 0
}
func (pself *PlayerAnimatedSprite) TriggerAnimation(velocity gd.Vector2, direction gd.Float, playerMode PlayerMode) {
	animationPrefix := strings.ToLower("small") // TODO (tanjp): Implement playerMode.String()
	parent := pself.Super().AsNode().GetParent(pself.Temporary)
	ch2d := *(*gd.CharacterBody2D)(unsafe.Pointer(&parent))
	node2d := pself.Super().AsNode2D()
	if !ch2d.IsOnFloor() {
		pself.Play(animationPrefix + "_jump")
	} else if Sign(velocity.X()) != Sign(gd.Float(direction)) && velocity.X() != 0 && direction != 0 {
		pself.Play(animationPrefix + "_slide")
		node2d.SetScale(gd.NewVector2(gd.Float(direction), node2d.GetScale().Y()))
	} else {
		if node2d.GetScale().X() == 1 && Sign(velocity.X()) == -1 {
			node2d.SetScale(gd.NewVector2(-1, node2d.GetScale().Y()))
		} else if node2d.GetScale().X() == -1 && Sign(velocity.X()) == 1 {
			node2d.SetScale(gd.NewVector2(1, node2d.GetScale().Y()))
		}
		if velocity.X() != 0 {
			pself.Play(animationPrefix + "_run")
		} else {
			pself.Play(animationPrefix + "_idle")
		}
	}
}
func (pself *PlayerAnimatedSprite) GetPlayer() *Player {
	parent := pself.Super().AsNode().GetParent(pself.Temporary)
	player, _ := gd.As[*Player](pself.Temporary, parent)
	return player
}

func (pself *PlayerAnimatedSprite) SetOffset(offset gd.Vector2) {
	pself.Super().AsAnimatedSprite2D().SetOffset(offset)
}
func (pself *PlayerAnimatedSprite) GetAnimation() string {
	return pself.Super().AsAnimatedSprite2D().GetAnimation(pself.Temporary).String()
}
func (pself *PlayerAnimatedSprite) OnAnimationFinished() {
	animName := pself.GetAnimation()
	player := pself.GetPlayer()

	switch animName {
	case "small_to_big":
		pself.ResetPlayerProperties()
		switch player.PlayerMode {
		case BIG:
			player.PlayerMode = SMALL
		case SMALL:
			player.PlayerMode = BIG
		}
	case "small_to_shooting", "big_to_shooting":
		pself.ResetPlayerProperties()
		player.PlayerMode = SHOOTING
	case "shoot":
		pself.GetPlayer().Super().AsNode().SetPhysicsProcess(true)
	}
}

func (pself *PlayerAnimatedSprite) ResetPlayerProperties() {
	pself.SetOffset(gd.NewVector2(0, 0))
	pself.GetPlayer().Super().AsNode().SetPhysicsProcess(true)
	pself.GetPlayer().Super().AsCollisionObject2D().SetCollisionLayerValue(1, true)
	pself.FrameCount = 0
}

func (pself *PlayerAnimatedSprite) OnFrameChanged() {
	animName := pself.GetAnimation()
	if animName == "small_to_big" || animName == "small_to_shooting" {
		playerMode := pself.GetPlayer().PlayerMode
		pself.FrameCount++

		if pself.FrameCount%2 == 1 {
			if playerMode == BIG {
				pself.SetOffset(gd.NewVector2(0, 0))
			} else {
				pself.SetOffset(gd.NewVector2(0, -8))
			}
		} else {
			if playerMode == BIG {
				pself.SetOffset(gd.NewVector2(0, 8))
			} else {
				pself.SetOffset(gd.NewVector2(0, 0))
			}
		}
	}
}

func (pself *PlayerAnimatedSprite) OnPole(playerMode PlayerMode) {
	animationPrefix := strings.ToLower(playerMode.String())
	pself.Play(animationPrefix + "_pole")
}
