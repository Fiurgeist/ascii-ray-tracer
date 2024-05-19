package main

import (
	"flag"
	"log"

	"github.com/fiurgeist/ascii-ray-tracer/internal/renderer"
	"github.com/fiurgeist/ascii-ray-tracer/internal/scenes"
)

func main() {
	width := flag.Int("width", 160, "rendered width")
	height := flag.Int("height", 90, "rendered height")
	mode := flag.String("mode", "console", "render mode")
	parallel := flag.Int("parallel", 1, "number of parallel render chunks")
	flag.Parse()

	scene := scenes.AssortedObjects()

	var renderMode renderer.Renderer

	switch *mode {
	case "console":
		renderMode = renderer.ConsoleRenderer{Width: *width, Height: *height}
	case "noop":
		renderMode = renderer.NoopRenderer{Width: *width, Height: *height}
	default:
		log.Fatalf("unsuppoted render mode %s", *mode)
	}

	renderMode.Render(scene, *parallel)
}
