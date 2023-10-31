package main

import (
	"flag"

	"github.com/fiurgeist/ascii-ray-tracer/internal/renderer"
	"github.com/fiurgeist/ascii-ray-tracer/internal/scenes"
)

func main() {
	width := flag.Int("width", 160, "rendered width")
	height := flag.Int("height", 90, "rendered height")
	flag.Parse()

	scene := scenes.AssortedObjects()

	renderer := renderer.ConsoleRenderer{Width: *width, Height: *height}
	renderer.Render(scene)
}
