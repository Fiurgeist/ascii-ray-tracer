package main

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/camera"
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/renderer"
	"github.com/fiurgeist/ascii-ray-tracer/internal/scene"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

func main() {
	camera := camera.NewCamera(vector.Vector{X: 0, Y: 2, Z: -8}, vector.Z)
	scene := scene.Scene{Background: color.Red, Camera: camera}

	renderer := renderer.ConsoleRenderer{Width: 32, Height: 18}
	renderer.Render(scene)
}
