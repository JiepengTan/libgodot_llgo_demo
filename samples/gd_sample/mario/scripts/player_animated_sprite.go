package scripts

import (
	"grow.graphics/gd"
	"strings"
)

type PlayerAnimatedSprite struct {
	gd.Class[PlayerAnimatedSprite, gd.AnimatedSprite2D] `gd:"PlayerAnimatedSprite"`

	FrameCount gd.Int
}

func (pself *PlayerAnimatedSprite) TriggerAnimation(velocity gd.Vector2, direction gd.Int, playerMode Player.PlayerMode) {
	animationPrefix := strings.ToLower(playerMode.String())

	if !pself.GetParent().AsNode().IsOnFloor() {
		pself.Play(animationPrefix + "_jump")
	} else if gd.Sign(velocity.X()) != gd.Sign(direction) && velocity.X() != 0 && direction != 0 {
		pself.Play(animationPrefix + "_slide")
		pself.SetScale(gd.NewVector2(gd.Float64(direction), pself.Scale().Y()))
	} else {
		if pself.Scale().X() == 1 && gd.Sign(velocity.X()) == -1 {
			pself.SetScale(gd.NewVector2(-1, pself.Scale().Y()))
		} else if pself.Scale().X() == -1 && gd.Sign(velocity.X()) == 1 {
			pself.SetScale(gd.NewVector2(1, pself.Scale().Y()))
		}

		if velocity.X() != 0 {
			pself.Play(animationPrefix + "_run")
		} else {
			pself.Play(animationPrefix + "_idle")
		}
	}
}

func (pself *PlayerAnimatedSprite) OnAnimationFinished() {
	switch pself.Animation() {
	case "small_to_big":
		pself.ResetPlayerProperties()
		switch pself.GetParent().(*Player).PlayerMode {
		case Player.PlayerMode_BIG:
			pself.GetParent().(*Player).PlayerMode = Player.PlayerMode_SMALL
		case Player.PlayerMode_SMALL:
			pself.GetParent().(*Player).PlayerMode = Player.PlayerMode_BIG
		}
	case "small_to_shooting", "big_to_shooting":
		pself.ResetPlayerProperties()
		pself.GetParent().(*Player).PlayerMode = Player.PlayerMode_SHOOTING
	case "shoot":
		pself.GetParent().(*Player).SetPhysicsProcess(true)
	}
}

func (pself *PlayerAnimatedSprite) ResetPlayerProperties() {
	pself.SetOffset(gd.NewVector2(0, 0))
	pself.GetParent().(*Player).SetPhysicsProcess(true)
	pself.GetParent().(*Player).SetCollisionLayerValue(1, true)
	pself.FrameCount = 0
}

func (pself *PlayerAnimatedSprite) OnFrameChanged() {
	if pself.Animation() == "small_to_big" || pself.Animation() == "small_to_shooting" {
		playerMode := pself.GetParent().(*Player).PlayerMode
		pself.FrameCount++

		if pself.FrameCount%2 == 1 {
			if playerMode == Player.PlayerMode_BIG {
				pself.SetOffset(gd.NewVector2(0, 0))
			} else {
				pself.SetOffset(gd.NewVector2(0, -8))
			}
		} else {
			if playerMode == Player.PlayerMode_BIG {
				pself.SetOffset(gd.NewVector2(0, 8))
			} else {
				pself.SetOffset(gd.NewVector2(0, 0))
			}
		}
	}
}

func (pself *PlayerAnimatedSprite) OnPole(playerMode Player.PlayerMode) {
	animationPrefix := strings.ToLower(playerMode.String())
	pself.Play(animationPrefix + "_pole")
}
