package main

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/renderer"
	"github.com/fiurgeist/ascii-ray-tracer/internal/scenes"
)

func main() {
	scene := scenes.AssortedObjects()

	renderer := renderer.ConsoleRenderer{Width: 160, Height: 90}
	renderer.Render(scene)
}
