package scripts

type CustomEnemy struct {
	Enemy
}

func (pself *CustomEnemy) Die() {
	pself.Enemy.Die()
	pself.Super().AsCollisionObject2D().SetCollisionLayerValue(3, false)
	pself.Super().AsCollisionObject2D().SetCollisionMaskValue(1, false)
	timer := pself.GetTree(pself.Temporary).CreateTimer(pself.Temporary, 1.5, true, false, false)
	timer.AsObject().Connect(pself.Temporary.StringName("timeout"), pself.Temporary.Callable(pself.DoQueueFree), 0)
}
