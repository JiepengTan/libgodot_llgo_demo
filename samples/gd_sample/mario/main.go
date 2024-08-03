package main

import (
	"gd-demos/scripts"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		panic("could not link to godot")
	}
	gd.Register[scripts.UI](godot)
	gd.Register[scripts.SceneData](godot)
	gd.Register[scripts.LevelManager](godot)
}
