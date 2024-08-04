package scripts

import (
	"fmt"

	"grow.graphics/gd"
)

type Player struct {
	gd.Class[Player, gd.CharacterBody2D] `gd:"Player"`

	AnimatedSprite2D         PlayerAnimatedSprite `gd:"AnimatedSprite2D"`
	Area2D                   gd.Area2D            `gd:"Area2D"`
	AreaCollisionShape       gd.CollisionShape2D  `gd:"Area2D/AreaCollisionShape"`
	BodyCollisionShape       gd.CollisionShape2D  `gd:"BodyCollisionShape"`
	ShootingPoint            gd.Node2D            `gd:"ShootingPoint"`
	SlideDownFinishedPos     gd.Node2D            `gd:"../slide_down_finished_position"`
	LandDownMarker           gd.Marker2D          `gd:"../LandDownMarker"`
	CameraSync               gd.Camera2D          `gd:"camera_sync"`
	CastlePath               gd.PathFollow2D      `gd:"castle_path"`
	PlayerMode               PlayerMode
	IsDead                   bool
	IsOnPath                 bool
	Gravity                  gd.Float
	RunSpeedDamping          gd.Float
	Speed                    gd.Float
	JumpVelocity             gd.Float
	MinStompDegree           gd.Float
	MaxStompDegree           gd.Float
	StompYVelocity           gd.Float
	ShouldCameraSync         bool
	PointsLabelScene         gd.PackedScene
	SmallMarioCollisionShape gd.Resource
	BigMarioCollisionShape   gd.Resource
	FireballScene            gd.PackedScene
	SceneData                SceneData

	CastleEntered gd.SignalAs[func()]
	PointsScored  gd.SignalAs[func(gd.Int)]
}

const PIPE_ENTER_THRESHOLD = 10

func (pself *Player) IsOnFloor() bool {
	return pself.Super().AsCharacterBody2D().IsOnFloor()
}
func (pself *Player) Position() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *Player) SetPosition(position gd.Vector2) {
	pself.Super().AsNode2D().SetPosition(position)
}
func (pself *Pipe) GetPosition() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *Player) GetPosition() gd.Vector2 {
	return pself.Super().AsNode2D().GetPosition()
}
func (pself *Player) GetTree(ctx gd.Lifetime) gd.SceneTree {
	return pself.Super().AsNode().GetTree(pself.Temporary)
}
func (pself *Player) String(s string) gd.String {
	return pself.Temporary.String(s)
}
func (pself *Player) NodePath(s string) gd.NodePath {
	return pself.String(s).NodePath(pself.Temporary)
}
func (pself *Player) StringName(s string) gd.StringName {
	return pself.Temporary.StringName(s)
}
func (pself *Player) Variant(v any) gd.Variant {
	return pself.Temporary.Variant(v)
}
func (pself *Player) DoQueueFree() {
	pself.Super().AsNode().QueueFree()
}
func (pself *Player) TweenVector2(tween gd.Tween, property string, final_val gd.Vector2, duration float32) gd.PropertyTweener {
	return tween.TweenProperty(
		pself.Temporary, pself.AsObject(),
		pself.Temporary.String(property).NodePath(pself.Temporary),
		pself.Temporary.Variant(final_val),
		gd.Float(duration))
}
func (pself *Player) DoTweenVector2(property string, dstValue gd.Vector2, duration float32, fn any) {
	spawnTween := pself.GetTree(pself.Temporary).CreateTween(pself.Temporary)
	pself.TweenVector2(spawnTween, property, dstValue, duration)
	spawnTween.TweenCallback(pself.Temporary, pself.Temporary.Callable(fn))
}
func (pself *Player) DoTweenChainVector2(property string, dstValue gd.Vector2, duration float32, chainVal gd.Vector2, cahainDuration float32, fn any) {
	spawnTween := pself.GetTree(pself.Temporary).CreateTween(pself.Temporary)
	pself.TweenVector2(spawnTween, property, dstValue, duration)
	pself.TweenVector2(spawnTween.Chain(pself.Temporary), property, chainVal, 4)
	spawnTween.TweenCallback(pself.Temporary, pself.Temporary.Callable(fn))
}
func (pself *Player) SetPositionX(val float64) {
	pself.SetPosition(gd.NewVector2(val, gd.Float(pself.GetPosition()[gd.Y])))
}
func (pself *Player) SetPositionY(val float64) {
	pself.SetPosition(gd.NewVector2(gd.Float(pself.GetPosition()[gd.X]), val))
}
func (pself *Player) SetVelocity(position gd.Vector2) {
	pself.Super().AsCharacterBody2D().SetVelocity(position)
}
func (pself *Player) GetVelocity() gd.Vector2 {
	return pself.Super().AsCharacterBody2D().GetVelocity()
}

func (pself *Player) SetVelocityX(val float64) {
	pself.SetVelocity(gd.NewVector2(val, gd.Float(pself.GetVelocity()[gd.Y])))
}
func (pself *Player) SetVelocityY(val float64) {
	pself.SetVelocity(gd.NewVector2(gd.Float(pself.GetVelocity()[gd.X]), val))
}
func (pself *Player) GetVelocityX() float64 {
	return gd.Float(pself.GetVelocity()[gd.X])
}
func (pself *Player) GetVelocityY() float64 {
	return gd.Float(pself.GetVelocity()[gd.Y])
}
func (pself *Player) Ready() {
	if pself.SceneData.ReturnPoint != gd.NewVector2(0, 0) {
		pself.Super().AsNode2D().SetGlobalPosition(pself.SceneData.ReturnPoint)
	}
}
func (pself *Player) GetGlobalPosition() gd.Vector2 {
	return pself.CameraSync.Super().AsNode2D().GetGlobalPosition()
}
func (pself *Player) PhysicsProcess(delta gd.Float) {
	globalPos := pself.CameraSync.Super().AsNode2D().GetGlobalPosition()
	cameraLeftBound := globalPos.X() - pself.CameraSync.Super().AsCanvasItem().GetViewportRect().Size.X()/2/pself.CameraSync.GetZoom().X()

	if !pself.IsOnFloor() {
		pself.SetVelocityY(pself.GetVelocityY() + pself.Gravity*float64(delta))
	}

	if globalPos.X() < cameraLeftBound+8 && pself.GetVelocity().X() < 0 {
		pself.SetVelocity(gd.NewVector2(0, 0))
		return
	}

	input := gd.Input(pself.Temporary)
	if input.IsActionJustPressed(pself.StringName("jump"), false) && pself.IsOnFloor() {
		pself.SetVelocityY(pself.JumpVelocity)
	}

	if input.IsActionJustReleased(pself.StringName("jump"), false) && pself.GetVelocity().Y() < 0 {
		pself.SetVelocityY(pself.GetVelocity().Y() * 0.5)
	}

	direction := input.GetAxis(pself.StringName("left"), pself.StringName("right"))
	if direction != 0 {
		pself.SetVelocityX(lerp(pself.GetVelocity().X(), pself.Speed*direction, pself.RunSpeedDamping*delta))
	} else {
		pself.SetVelocityX(gd.MoveToward(pself.GetVelocity().X(), 0, pself.Speed*delta))
	}

	if input.IsActionJustPressed(pself.StringName("shoot"), false) && pself.PlayerMode == SHOOTING {
		pself.Shoot()
	} else {
		pself.AnimatedSprite2D.TriggerAnimation(pself.GetVelocity(), direction, pself.PlayerMode)
	}
	collision := pself.Super().GetLastSlideCollision(pself.Temporary)
	// TODO(tanjp) check is not null
	if collision.AsPointer().Pointer()[0] != 0 {
		pself.HandleMovementCollision(collision)
	}

	pself.Super().MoveAndSlide()
}

func (pself *Player) Process(delta gd.Float) {
	camPos := pself.CameraSync.AsNode2D().GetGlobalPosition()
	if pself.GetGlobalPosition().X() > camPos.X() && pself.ShouldCameraSync {
		pself.CameraSync.AsNode2D().SetGlobalPosition(gd.NewVector2(pself.GetGlobalPosition().X(), camPos.Y()))
	}

	if pself.IsOnPath {
		pself.CastlePath.SetProgress(pself.CastlePath.GetProgress() + delta*pself.Speed/2)
		if pself.CastlePath.GetProgressRatio() > 0.97 {
			pself.IsOnPath = false
			pself.LandDown()
		}
	}
}

func (pself *Player) OnArea2DAreaEntered(area gd.Area2D) {
	if enemy, ok := gd.As[*Enemy](pself.Temporary, area); ok {
		pself.HandleEnemyCollision(enemy)
	}
	if shroom, ok := gd.As[*Shroom](pself.Temporary, area); ok {
		pself.HandleShroomCollision(shroom.Super().AsNode2D())
		area.Super().AsNode().QueueFree()
	}
	if _, ok := gd.As[*ShootingFlower](pself.Temporary, area); ok {
		pself.HandleFlowerCollision()
		area.Super().AsNode().QueueFree()
	}
}

func (pself *Player) HandleEnemyCollision(enemy *Enemy) {
	if enemy == nil && pself.IsDead {
		return
	}
	levelManager := pself.GetLevelManager()
	if koopa, ok := gd.As[*Koopa](pself.Temporary, enemy); ok && koopa.InAShell {

		koopa.OnStomp(pself.GetGlobalPosition())
		pself.SpawnPointsLabel(enemy)
		levelManager.OnPointsScored(100)
	} else {
		rad := pself.GetPosition().AngleToPoint(enemy.GetPosition())
		angleOfCollision := pself.Temporary.RadToDeg(gd.Float(rad))

		if angleOfCollision > pself.MinStompDegree && pself.MaxStompDegree > angleOfCollision {
			enemy.Die()
			pself.OnEnemyStomped()
			pself.SpawnPointsLabel(enemy)
			levelManager.OnPointsScored(100)
		} else {
			pself.Die()
		}
	}
}

func (pself *Player) SetPhysicsProcess(value bool) {
	pself.Super().AsNode().SetPhysicsProcess(value)
}
func (pself *Player) HandleShroomCollision(area gd.Node2D) {
	if pself.PlayerMode == SMALL {
		pself.SetPhysicsProcess(false)
		pself.AnimatedSprite2D.Play("small_to_big")
		pself.SetCollisionShapes(false)
	}
}

func (pself *Player) HandleFlowerCollision() {
	pself.SetPhysicsProcess(false)
	animationName := "small_to_shooting"
	if pself.PlayerMode != SMALL {
		animationName = "big_to_shooting"
	}
	pself.AnimatedSprite2D.Play(animationName)
	pself.SetCollisionShapes(false)
}

func (pself *Player) SpawnPointsLabel(enemy *Enemy) {
	pointsLabel := pself.PointsLabelScene.Instantiate(pself.Temporary, gd.PackedSceneGenEditState(0))
	if label, ok := gd.As[*gd.Label](pself.Temporary, pointsLabel); ok {
		label.Super().AsControl().SetPosition(enemy.GetPosition().Add(gd.NewVector2(-20, -20)), true)
	}

	pself.GetTree(pself.Temporary).GetRoot(pself.Temporary).AsNode().AddChild(pointsLabel, true, gd.NodeInternalMode(0))
	pself.PointsScored.Emit(100)
}

func (pself *Player) OnEnemyStomped() {
	pself.SetVelocityY(pself.StompYVelocity)
}

func (pself *Player) Die() {
	if pself.PlayerMode == SMALL {
		pself.IsDead = true
		pself.AnimatedSprite2D.Play("death")
		pself.Area2D.AsCollisionObject2D().SetCollisionMaskValue(3, false)
		pself.Super().AsCollisionObject2D().SetCollisionLayerValue(1, false)
		pself.SetPhysicsProcess(false)
		pself.DoTweenChainVector2("position",
			pself.GetPosition().Add(gd.NewVector2(0, -48)), 0.5,
			pself.GetPosition().Add(gd.NewVector2(0, 256)), 1,
			func() { pself.GetTree(pself.Temporary).ReloadCurrentScene() })
	} else {
		pself.BigToSmall()
	}
}

func (pself *Player) HandleMovementCollision(collision gd.KinematicCollision2D) {
	collisionAngle := pself.Temporary.RadToDeg(collision.GetAngle(gd.NewVector2(0, 1)))
	collider := collision.GetCollider(pself.Temporary)
	if block, ok := gd.As[*Block](pself.Temporary, collider); ok {
		if pself.Temporary.Roundf(collisionAngle) == 180 {
			block.Bump(pself.PlayerMode)
		}
	}

	input := gd.Input(pself.Temporary)
	if pipe, ok := gd.As[*Pipe](pself.Temporary, collider); ok {
		if pself.Temporary.Roundf(collisionAngle) == 0 &&
			input.IsActionJustPressed(pself.StringName("down"), false) &&
			pself.Temporary.Absf(pipe.GetPosition().X()-pself.GetPosition().X()) < PIPE_ENTER_THRESHOLD &&
			pipe.IsTraversable {

			fmt.Println("GO DOWN")
			pself.HandlePipeCollision()
		}
	}
}

func (pself *Player) SetCollisionShapes(isSmall bool) {
	collisionShape := pself.SmallMarioCollisionShape
	if !isSmall {
		collisionShape = pself.BigMarioCollisionShape
	}
	pself.AreaCollisionShape.AsObject().SetDeferred(pself.StringName("shape"), pself.Variant(collisionShape))
	pself.BodyCollisionShape.AsObject().SetDeferred(pself.StringName("shape"), pself.Variant(collisionShape))

}

func (pself *Player) BigToSmall() {
	pself.Super().AsCollisionObject2D().SetCollisionLayerValue(1, false)
	pself.SetPhysicsProcess(false)
	animationName := "small_to_big"
	if pself.PlayerMode == BIG {
		animationName = "big_to_shooting"
	}
	pself.AnimatedSprite2D.Play(animationName)
	pself.SetCollisionShapes(true)
}

func (pself *Player) Shoot() {
	pself.AnimatedSprite2D.Play("shoot")
	pself.SetPhysicsProcess(false)

	// TODO (tanjp): Implement Fireball
	//fireball := pself.FireballScene.Instantiate()
	//fireball.(*Fireball).SetDirection(sign(pself.AnimatedSprite2D.GetScale().X))
	//fireball.SetGlobalPosition(pself.ShootingPoint.GetGlobalPosition())
	//pself.GetTree(pself.Temporary).GetRoot().AddChild(fireball)
}

func (pself *Player) HandlePipeCollision() {
	pself.SetPhysicsProcess(false)
	pself.Super().AsCanvasItem().SetZIndex(-3)
	pself.DoTweenVector2("position", pself.GetPosition().Add(gd.NewVector2(0, 32)), 1, pself.SwitchToUnderground)
}

func (pself *Player) SwitchToUnderground() {
	pself.SceneData.PlayerMode = pself.PlayerMode
	pself.SceneData.Coins = pself.GetLevelManager().Coins
	pself.SceneData.Points = pself.GetLevelManager().Points
	pself.GetTree(pself.Temporary).ChangeSceneToFile(pself.String("res://Scenes/underground.tscn"))
}

func (pself *Player) HandlePipeConnectorEntranceCollision() {
	pself.SetPhysicsProcess(false)
	pself.DoTweenVector2("position", pself.GetPosition().Add(gd.NewVector2(32, 0)), 1, pself.SwitchToMain)
}
func (pself *Player) GetLevelManager() *LevelManager {
	levelManagerNode := pself.GetTree(pself.Temporary).GetFirstNodeInGroup(pself.Temporary, pself.Temporary.StringName("level_manager"))
	if levelManager, ok := gd.As[*LevelManager](pself.Temporary, levelManagerNode); ok {
		return levelManager
	}
	return nil
}
func (pself *Player) SwitchToMain() {
	pself.SceneData.PlayerMode = pself.PlayerMode
	pself.SceneData.Coins = pself.GetLevelManager().Coins
	pself.SceneData.Points = pself.GetLevelManager().Points
	pself.GetTree(pself.Temporary).ChangeSceneToFile(pself.String("res://Scenes/main.tscn"))
}

func (pself *Player) OnPoleHit() {
	pself.SetPhysicsProcess(false)
	pself.SetVelocity(gd.NewVector2(0, 0))
	if pself.IsOnPath {
		return
	}

	pself.AnimatedSprite2D.OnPole(pself.PlayerMode)
	pself.DoTweenVector2("position", pself.SlideDownFinishedPos.GetPosition(), 2, pself.SlideDownFinished)
}

func (pself *Player) SlideDownFinished() {
	animationPrefix := pself.PlayerMode.String()
	pself.IsOnPath = true
	pself.AnimatedSprite2D.Play(fmt.Sprintf("%s_jump", animationPrefix))
	pself.Super().AsNode().Reparent(pself.CastlePath.AsNode(), false)
}

func (pself *Player) GetRootNode(name string) gd.Node {
	return pself.GetTree(pself.Temporary).GetRoot(pself.Temporary).AsNode().GetNode(pself.Temporary, pself.NodePath((name)))
}
func (pself *Player) LandDown() {
	dstParent := pself.GetRootNode("main")
	pself.Super().AsNode().Reparent(dstParent, false)
	distanceToMarker := pself.LandDownMarker.AsNode2D().GetPosition().Y() - pself.GetPosition().Y()
	pself.DoTweenVector2("position", pself.GetPosition().Add(gd.NewVector2(0, distanceToMarker-pself.GetHalfSpriteSize())), 0.5, pself.GoToCastle)
}

func (pself *Player) GoToCastle() {
	animationPrefix := pself.PlayerMode.String()
	pself.AnimatedSprite2D.Play(fmt.Sprintf("%s_run", animationPrefix))
	pself.DoTweenVector2("position", pself.GetPosition().Add(gd.NewVector2(75, 0)), 0.5, pself.Finish)
}

func (pself *Player) Finish() {
	pself.DoQueueFree()
	pself.CastleEntered.Emit()
}

func (pself *Player) GetHalfSpriteSize() float64 {
	if pself.PlayerMode == SMALL {
		return 8
	}
	return 16
}
