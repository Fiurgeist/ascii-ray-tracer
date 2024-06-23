package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/fiurgeist/ascii-ray-tracer/internal/renderer"
	"github.com/fiurgeist/ascii-ray-tracer/internal/scenes"
)

func main() {
	width := flag.Int("width", 160, "rendered width")
	height := flag.Int("height", 90, "rendered height")
	display := flag.String("display", "console", "display mode; console or noop")
	processor := flag.String("processor", "cpu", "processor type; cpu or gpu")
	output := flag.String("output", "still", "output type; still or loop")
	parallel := flag.Int("parallel", 1, "if cpu: number of parallel render chunks")
	flag.Parse()

	scene := scenes.AssortedObjects()

	var renderMode renderer.Renderer

	if *display == "noop" {
		renderMode = &renderer.NoopRenderer{Width: *width, Height: *height, Parallel: *parallel}
	} else {
		consoleMode := &renderer.ConsoleRenderer{Width: *width, Height: *height, Parallel: *parallel}

		renderMode = consoleMode

		if *output == "loop" {
			renderMode = &renderer.GameLoopRenderer{
				Renderer: *consoleMode,
			}
		}
	}

	if *output == "loop" {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		go func() {
			renderMode.Render(scene, *processor)
		}()

		<-stop

		renderMode.(*renderer.GameLoopRenderer).Stop()
		renderer.ResetConsole()
	} else {
		renderMode.Render(scene, *processor)
	}
}
