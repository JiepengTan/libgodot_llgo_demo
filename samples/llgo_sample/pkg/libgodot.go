package libgodot

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

const (
	LLGoPackage = "link: $(pkg-config --libs godot); -lgodot"
)

// llgo:link CreateGodotInstance C.libgodot_create_godot_instance
func CreateGodotInstance(p_argc c.Int, p_argv **c.Char, p_platform_data c.Pointer) c.Pointer {
	return nil
}

// llgo:link DestroyGodotInstance C.libgodot_destroy_godot_instance
func DestroyGodotInstance(p_godot_instance c.Pointer) {}

// llgo:link StartGodotInstance C.libgodot_start_godot_instance
func StartGodotInstance(p_godot_instance c.Pointer) bool { return false }

// llgo:link IterateGodotInstance C.libgodot_iteration_godot_instance
func IterateGodotInstance(p_godot_instance c.Pointer) bool { return false }
