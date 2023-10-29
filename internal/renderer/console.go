package renderer

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/scene"
)

const (
	// pixel = "██" // two of those in Liberation Mono are pretty much quadratic
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

func (r ConsoleRenderer) Render(scene scene.Scene) {
	var sb strings.Builder
	start := time.Now()
	width := float64(r.Width)
	height := float64(r.Height)

	for py := 0; py < r.Height; py += 2 {
		for px := 0; px < r.Width; px++ {
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
		sb.WriteString("\n")
	}
	io.WriteString(os.Stdout, sb.String())
	io.WriteString(os.Stdout, fmt.Sprintf("Rendering took %fs\n", time.Since(start).Seconds()))
}

func ansiColor(color color.Color) string {
	return fmt.Sprintf("%s38;2;%d;%d;%dm", esc, color.R(), color.G(), color.B())
}

func ansiBackgroundColor(color color.Color) string {
	return fmt.Sprintf("%s48;2;%d;%d;%dm", esc, color.R(), color.G(), color.B())
}
