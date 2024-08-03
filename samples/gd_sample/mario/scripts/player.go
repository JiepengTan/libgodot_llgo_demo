package scripts

import (
	"fmt"

	"grow.graphics/gd"
)

type Player struct {
	gd.Class[Player, gd.CharacterBody2D] `gd:"Player"`

	AnimatedSprite2D         *gd.PlayerAnimatedSprite `gd:"AnimatedSprite2D"`
	Area2D                   gd.Area2D                `gd:"Area2D"`
	AreaCollisionShape       gd.CollisionShape2D      `gd:"Area2D/AreaCollisionShape"`
	BodyCollisionShape       gd.CollisionShape2D      `gd:"BodyCollisionShape"`
	ShootingPoint            gd.Node2D                `gd:"ShootingPoint"`
	SlideDownFinishedPos     gd.Node2D                `gd:"../slide_down_finished_position"`
	LandDownMarker           gd.Marker2D              `gd:"../LandDownMarker"`
	CameraSync               *gd.Camera2D             `gd:"camera_sync"`
	CastlePath               *gd.PathFollow2D         `gd:"castle_path"`
	PlayerMode               PlayerMode
	IsDead                   bool
	IsOnPath                 bool
	Gravity                  float64
	RunSpeedDamping          float64
	Speed                    float64
	JumpVelocity             float64
	MinStompDegree           float64
	MaxStompDegree           float64
	StompYVelocity           float64
	ShouldCameraSync         bool
	PointsLabelScene         *gd.PackedScene
	SmallMarioCollisionShape *gd.Resource
	BigMarioCollisionShape   *gd.Resource
	FireballScene            *gd.PackedScene
}

type PlayerMode int

const (
	SMALL PlayerMode = iota
	BIG
	SHOOTING
)

const PIPE_ENTER_THRESHOLD = 10

func (p *Player) _Ready() {
	if gd.SceneData.ReturnPoint != nil && gd.SceneData.ReturnPoint != gd.NewVector2(0, 0) {
		p.SetGlobalPosition(gd.SceneData.ReturnPoint)
	}
}

func (p *Player) _PhysicsProcess(delta gd.Float) {
	cameraLeftBound := p.CameraSync.GetGlobalPosition().X - p.CameraSync.GetViewportRect().Size.X/2/p.CameraSync.GetZoom().X

	if !p.IsOnFloor() {
		p.SetVelocityY(p.GetVelocityY() + p.Gravity*float64(delta))
	}

	if p.GetGlobalPosition().X < cameraLeftBound+8 && p.GetVelocity().X < 0 {
		p.SetVelocity(gd.NewVector2(0, 0))
		return
	}

	if gd.Input.IsActionJustPressed("jump") && p.IsOnFloor() {
		p.SetVelocityY(p.JumpVelocity)
	}

	if gd.Input.IsActionJustReleased("jump") && p.GetVelocity().Y < 0 {
		p.SetVelocityY(p.GetVelocity().Y * 0.5)
	}

	direction := gd.Input.GetAxis("left", "right")
	if direction != 0 {
		p.SetVelocityX(gd.Lerp(p.GetVelocity().X, p.Speed*direction, p.RunSpeedDamping*delta))
	} else {
		p.SetVelocityX(gd.MoveToward(p.GetVelocity().X, 0, p.Speed*delta))
	}

	if gd.Input.IsActionJustPressed("shoot") && p.PlayerMode == SHOOTING {
		p.Shoot()
	} else {
		p.AnimatedSprite2D.TriggerAnimation(p.GetVelocity(), direction, p.PlayerMode)
	}

	if collision := p.GetLastSlideCollision(); collision != nil {
		p.HandleMovementCollision(collision)
	}

	p.MoveAndSlide()
}

func (p *Player) _Process(delta gd.Float) {
	if p.GetGlobalPosition().X > p.CameraSync.GetGlobalPosition().X && p.ShouldCameraSync {
		p.CameraSync.SetGlobalPositionX(p.GetGlobalPosition().X)
	}

	if p.IsOnPath {
		p.CastlePath.SetProgress(p.CastlePath.GetProgress() + delta*p.Speed/2)
		if p.CastlePath.GetProgressRatio() > 0.97 {
			p.IsOnPath = false
			p.LandDown()
		}
	}
}

func (p *Player) _OnArea2DAreaEntered(area gd.Area2D) {
	switch area := area.(type) {
	case *Enemy:
		p.HandleEnemyCollision(area)
	case *Shroom:
		p.HandleShroomCollision(area)
		area.QueueFree()
	case *ShootingFlower:
		p.HandleFlowerCollision()
		area.QueueFree()
	}
}

func (p *Player) HandleEnemyCollision(enemy *Enemy) {
	if enemy == nil && p.IsDead {
		return
	}
	levelManager := gd.GetTree().GetFirstNodeInGroup("level_manager")

	if koopa, ok := enemy.(*Koopa); ok && koopa.InAShell {
		koopa.OnStomp(p.GetGlobalPosition())
		p.SpawnPointsLabel(enemy)
		levelManager.OnPointsScored(100)
	} else {
		angleOfCollision := gd.RadToDeg(p.GetPosition().AngleToPoint(enemy.GetPosition()))

		if angleOfCollision > p.MinStompDegree && p.MaxStompDegree > angleOfCollision {
			enemy.Die()
			p.OnEnemyStomped()
			p.SpawnPointsLabel(enemy)
			levelManager.OnPointsScored(100)
		} else {
			p.Die()
		}
	}
}

func (p *Player) HandleShroomCollision(area gd.Node2D) {
	if p.PlayerMode == SMALL {
		p.SetPhysicsProcess(false)
		p.AnimatedSprite2D.Play("small_to_big")
		p.SetCollisionShapes(false)
	}
}

func (p *Player) HandleFlowerCollision() {
	p.SetPhysicsProcess(false)
	animationName := "small_to_shooting"
	if p.PlayerMode != SMALL {
		animationName = "big_to_shooting"
	}
	p.AnimatedSprite2D.Play(animationName)
	p.SetCollisionShapes(false)
}

func (p *Player) SpawnPointsLabel(enemy *Enemy) {
	pointsLabel := p.PointsLabelScene.Instantiate()
	pointsLabel.SetPosition(enemy.GetPosition().Add(gd.NewVector2(-20, -20)))
	gd.GetTree().GetRoot().AddChild(pointsLabel)
	p.EmitSignal("points_scored", 100)
}

func (p *Player) OnEnemyStomped() {
	p.SetVelocityY(p.StompYVelocity)
}

func (p *Player) Die() {
	if p.PlayerMode == SMALL {
		p.IsDead = true
		p.AnimatedSprite2D.Play("death")
		p.Area2D.SetCollisionMaskValue(3, false)
		p.SetCollisionLayerValue(1, false)
		p.SetPhysicsProcess(false)
		deathTween := gd.GetTree().CreateTween()
		deathTween.TweenProperty(p, "position", p.GetPosition().Add(gd.NewVector2(0, -48)), 0.5)
		deathTween.Chain().TweenProperty(p, "position", p.GetPosition().Add(gd.NewVector2(0, 256)), 1)
		deathTween.TweenCallback(func() { gd.GetTree().ReloadCurrentScene() })
	} else {
		p.BigToSmall()
	}
}

func (p *Player) HandleMovementCollision(collision gd.KinematicCollision2D) {
	if block, ok := collision.GetCollider().(*Block); ok {
		collisionAngle := gd.RadToDeg(collision.GetAngle())
		if roundf(collisionAngle) == 180 {
			block.Bump(p.PlayerMode)
		}
	}

	if pipe, ok := collision.GetCollider().(*Pipe); ok {
		collisionAngle := gd.RadToDeg(collision.GetAngle())
		if roundf(collisionAngle) == 0 && gd.Input.IsActionJustPressed("down") && absf(pipe.GetPosition().X-p.GetPosition().X) < PIPE_ENTER_THRESHOLD && pipe.IsTraversable() {
			fmt.Println("GO DOWN")
			p.HandlePipeCollision()
		}
	}
}

func (p *Player) SetCollisionShapes(isSmall bool) {
	collisionShape := p.SmallMarioCollisionShape
	if !isSmall {
		collisionShape = p.BigMarioCollisionShape
	}
	p.AreaCollisionShape.SetDeferred("shape", collisionShape)
	p.BodyCollisionShape.SetDeferred("shape", collisionShape)
}

func (p *Player) BigToSmall() {
	p.SetCollisionLayerValue(1, false)
	p.SetPhysicsProcess(false)
	animationName := "small_to_big"
	if p.PlayerMode == BIG {
		animationName = "big_to_shooting"
	}
	p.AnimatedSprite2D.Play(animationName, 1.0, true)
	p.SetCollisionShapes(true)
}

func (p *Player) Shoot() {
	p.AnimatedSprite2D.Play("shoot")
	p.SetPhysicsProcess(false)

	fireball := p.FireballScene.Instantiate()
	fireball.(*Fireball).SetDirection(sign(p.AnimatedSprite2D.GetScale().X))
	fireball.SetGlobalPosition(p.ShootingPoint.GetGlobalPosition())
	gd.GetTree().GetRoot().AddChild(fireball)
}

func (p *Player) HandlePipeCollision() {
	p.SetPhysicsProcess(false)
	p.SetZIndex(-3)
	pipeTween := gd.GetTree().CreateTween()
	pipeTween.TweenProperty(p, "position", p.GetPosition().Add(gd.NewVector2(0, 32)), 1)
	pipeTween.TweenCallback(p.SwitchToUnderground)
}

func (p *Player) SwitchToUnderground() {
	levelManager := gd.GetTree().GetFirstNodeInGroup("level_manager")
	gd.SceneData.PlayerMode = p.PlayerMode
	gd.SceneData.Coins = levelManager.(*LevelManager).Coins
	gd.SceneData.Points = levelManager.(*LevelManager).Points
	gd.GetTree().ChangeSceneToFile("res://Scenes/underground.tscn")
}

func (p *Player) HandlePipeConnectorEntranceCollision() {
	p.SetPhysicsProcess(false)
	pipeTween := gd.GetTree().CreateTween()
	pipeTween.TweenProperty(p, "position", p.GetPosition().Add(gd.NewVector2(32, 0)), 1)
	pipeTween.TweenCallback(p.SwitchToMain)
}

func (p *Player) SwitchToMain() {
	levelManager := gd.GetTree().GetFirstNodeInGroup("level_manager")
	gd.SceneData.PlayerMode = p.PlayerMode
	gd.SceneData.Coins = levelManager.(*LevelManager).Coins
	gd.SceneData.Points = levelManager.(*LevelManager).Points
	gd.GetTree().ChangeSceneToFile("res://Scenes/main.tscn")
}

func (p *Player) OnPoleHit() {
	p.SetPhysicsProcess(false)
	p.SetVelocity(gd.NewVector2(0, 0))
	if p.IsOnPath {
		return
	}

	p.AnimatedSprite2D.OnPole(p.PlayerMode)

	slideDownTween := gd.GetTree().CreateTween()
	slideDownPos := p.SlideDownFinishedPos.GetPosition()
	slideDownTween.TweenProperty(p, "position", slideDownPos, 2)
	slideDownTween.TweenCallback(p.SlideDownFinished)
}

func (p *Player) SlideDownFinished() {
	animationPrefix := PlayerModeToString(p.PlayerMode)
	p.IsOnPath = true
	p.AnimatedSprite2D.Play(fmt.Sprintf("%s_jump", animationPrefix))
	p.Reparent(p.CastlePath)
}

func (p *Player) LandDown() {
	p.Reparent(gd.GetTree().GetRoot().GetNode("main"))
	distanceToMarker := p.LandDownMarker.GetPosition().Y - p.GetPosition().Y
	landTween := gd.GetTree().CreateTween()
	landTween.TweenProperty(p, "position", p.GetPosition().Add(gd.NewVector2(0, distanceToMarker-p.GetHalfSpriteSize())), 0.5)
	landTween.TweenCallback(p.GoToCastle)
}

func (p *Player) GoToCastle() {
	animationPrefix := PlayerModeToString(p.PlayerMode)
	p.AnimatedSprite2D.Play(fmt.Sprintf("%s_run", animationPrefix))

	runToCastleTween := gd.GetTree().CreateTween()
	runToCastleTween.TweenProperty(p, "position", p.GetPosition().Add(gd.NewVector2(75, 0)), 0.5)
	runToCastleTween.TweenCallback(p.Finish)
}

func (p *Player) Finish() {
	p.QueueFree()
	p.EmitSignal("castle_entered")
}

func (p *Player) GetHalfSpriteSize() float64 {
	if p.PlayerMode == SMALL {
		return 8
	}
	return 16
}

func PlayerModeToString(mode PlayerMode) string {
	switch mode {
	case SMALL:
		return "small"
	case BIG:
		return "big"
	case SHOOTING:
		return "shooting"
	default:
		return ""
	}
}
