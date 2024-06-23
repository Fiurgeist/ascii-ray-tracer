package renderer

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/scene"
	"github.com/fiurgeist/ascii-ray-tracer/internal/shader"
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
	Width    int
	Height   int
	Parallel int
	shader   *shader.Shader
}

type PixelGetter func(x, y int) color.Color

func (r ConsoleRenderer) Render(scene scene.Scene, processor string) {
	if processor == "gpu" {
		r.shader = &shader.Shader{Width: int32(r.Width), Height: int32(r.Height)}

		r.shader.Init(scene)
		defer r.shader.Delete()
	}

	start := time.Now()

	fmt.Printf("%s2J", esc)

	if processor == "gpu" {
		r.gpuRender(scene)
	} else {
		r.cpuRender(scene)
	}

	fmt.Printf("%s%d;%dH", esc, r.Height/2+1, 1)
	fmt.Printf("Rendering took: %fs\n", time.Since(start).Seconds())
}

func (r ConsoleRenderer) cpuRender(scene scene.Scene) {
	fmt.Printf("%s%d;%dH", esc, 1, 1)

	inc := int(math.Ceil(float64(r.Width) / float64(r.Parallel)))
	startX := 0
	endX := inc
	width := float64(r.Width)
	height := float64(r.Height)

	getPixel := func(px, py int) color.Color {
		x := (float64(px) / width) - 0.5
		y := (float64(py) / height) - 0.5

		return scene.Trace(x, y)
	}

	var wg sync.WaitGroup
	for range r.Parallel {
		wg.Add(1)

		go func(startX, endX int) {
			defer wg.Done()
			r.renderSection(startX, endX, getPixel)
		}(startX, endX)

		startX = endX
		endX += inc
		if endX > r.Width {
			endX = r.Width
		}
	}

	wg.Wait()
}

func (r ConsoleRenderer) gpuRender(scene scene.Scene) {
	fmt.Printf("%s%d;%dH", esc, 1, 1)

	pixels := r.shader.Compute(scene)
	getPixel := func(x, y int) color.Color {
		pos := (x + y*r.Width) * 3

		return color.NewColor(uint8(pixels[pos]), uint8(pixels[pos+1]), uint8(pixels[pos+2]))
	}

	r.renderSection(0, r.Width, getPixel)
}

func (r ConsoleRenderer) renderSection(startX, endX int, getPixel PixelGetter) {
	for y := 0; y < r.Height; y += 2 {
		var sb strings.Builder

		for x := startX; x < endX; x++ {
			foregroundColor := getPixel(x, y)
			backgroundColor := getPixel(x, y+1)

			sb.WriteString(fmt.Sprintf(
				"%s%s%s%s",
				ansiColor(foregroundColor), ansiBackgroundColor(backgroundColor), pixel, colorReset,
			))
		}

		fmt.Printf("%s%d;%dH%s", esc, y/2+1, startX, sb.String())
	}
}

func ansiColor(color color.Color) string {
	return fmt.Sprintf("%s38;2;%d;%d;%dm", esc, color.R(), color.G(), color.B())
}

func ansiBackgroundColor(color color.Color) string {
	return fmt.Sprintf("%s48;2;%d;%d;%dm", esc, color.R(), color.G(), color.B())
}

func ResetConsole() {
	fmt.Printf(colorReset)
	fmt.Printf("%s2J", esc)
	fmt.Printf("%s%d;%dH", esc, 1, 1)
}
