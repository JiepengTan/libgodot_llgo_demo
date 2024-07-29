package main

import (
	libgodot "libgodotdemo/pkg"
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

func main() {
	println("Hello, libgodot llgo111!")
	var instance c.Pointer
	strs := make([]*c.Char, 7)
	strs[0] = c.Str("aaaa")
	strs[1] = c.Str("--path")
	strs[2] = c.Str("../../go_sample/project")
	strs[3] = c.Str("--rendering-method")
	strs[4] = c.Str("gl_compatibility")
	strs[5] = c.Str("--rendering-driver")
	strs[6] = c.Str("opengl3")
	/*
		// pass the path of the project to the godot instance
		strs[0] = c.Str(os.Args[0])
		projectPath := "../../project/"
		if len(os.Args) > 1 {
			projectPath = os.Args[1]
		}
		strs[2] = c.Str(projectPath)
	*/
	instance = libgodot.CreateGodotInstance(7, &strs[0], nil)
	if !libgodot.StartGodotInstance(instance) {
		println("Failed to start Godot instance")
	}
	for !libgodot.IterateGodotInstance(instance) {
	}
	libgodot.DestroyGodotInstance(instance)
}
