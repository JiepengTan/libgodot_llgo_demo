// auto generated file. DO NOT EDIT
package autogen

import "C"
import (
	"godot-go-demo-projects/2d/dodgethecreep/scripts/demo"

	"github.com/godot-go/godot-go/pkg/core"
)

func RegisterSceneInitializer() {
	registerClassHUD()
}
func registerClassHUD() {
	core.AutoRegisterClassDB[*demo.HUD]()
}
