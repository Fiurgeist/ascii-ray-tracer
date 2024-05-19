package renderer

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/scene"
)

const (
	// pixel = "██" // two of those in Liberation Mono are almost quadratic
	pixel = "▀" // "double pixel" (coloring foreground and background separately)
	// ANSI codes: https://en.wikipedia.org/wiki/ANSI_escape_code
	esc        = "\u001B["
	colorReset = esc + "0m"
)

var _ Renderer = (*ConsoleRenderer)(nil)

type ConsoleRenderer struct {
	Width  int
	Height int
}

func (r ConsoleRenderer) Render(scene scene.Scene, parallel int) {
	start := time.Now()

	fmt.Printf("%s2J", esc)
	fmt.Printf("%s%d;%dH", esc, 1, 1)

	inc := int(math.Ceil(float64(r.Width) / float64(parallel)))
	startX := 0
	endX := inc

	var wg sync.WaitGroup
	for range parallel {
		wg.Add(1)

		go func(startX, endX int) {
			defer wg.Done()
			r.renderSection(scene, startX, endX)
		}(startX, endX)

		startX = endX
		endX += inc
		if endX > r.Width {
			endX = r.Width
		}
	}

	wg.Wait()

	fmt.Printf("%s%d;%dH", esc, r.Height/2+1, 1)
	fmt.Printf("Rendering took %fs\n", time.Since(start).Seconds())
}

func (r ConsoleRenderer) renderSection(scene scene.Scene, startX, endX int) {
	width := float64(r.Width)
	height := float64(r.Height)

	for py := 0; py < r.Height; py += 2 {
		var sb strings.Builder

		for px := startX; px < endX; px++ {
			x := (float64(px) / width) - 0.5
			y1 := (float64(py) / height) - 0.5
			y2 := (float64(py+1) / height) - 0.5

			foregroundColor := scene.Trace(x, y1)
			backgroundColor := scene.Trace(x, y2)

			sb.WriteString(fmt.Sprintf(
				"%s%s%s%s",
				ansiColor(foregroundColor), ansiBackgroundColor(backgroundColor), pixel, colorReset,
			))
		}

		fmt.Printf("%s%d;%dH%s", esc, py/2+1, startX, sb.String())
	}
}

func ansiColor(color color.Color) string {
	return fmt.Sprintf("%s38;2;%d;%d;%dm", esc, color.R(), color.G(), color.B())
}

func ansiBackgroundColor(color color.Color) string {
	return fmt.Sprintf("%s48;2;%d;%d;%dm", esc, color.R(), color.G(), color.B())
}
