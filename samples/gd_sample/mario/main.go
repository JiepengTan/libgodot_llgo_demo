package main

import (
	"grow.graphics/gd/gdextension"
)

func main() {
	_, ok := gdextension.Link()
	if !ok {
		panic("could not link to godot")
	}
}
