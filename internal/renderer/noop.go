package renderer

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/fiurgeist/ascii-ray-tracer/internal/scene"
)

var _ Renderer = (*NoopRenderer)(nil)

type NoopRenderer struct {
	Width    int
	Height   int
	Parallel int
}

func (r NoopRenderer) Render(scene scene.Scene) {
	start := time.Now()

	inc := int(math.Ceil(float64(r.Width) / float64(r.Parallel)))
	startX := 0
	endX := inc

	var wg sync.WaitGroup
	for range r.Parallel {
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

	fmt.Printf("Ray-tracing took %fs\n", time.Since(start).Seconds())
}

func (r NoopRenderer) renderSection(scene scene.Scene, startX, endX int) {
	width := float64(r.Width)
	height := float64(r.Height)

	for py := 0; py < r.Height; py++ {
		for px := startX; px < endX; px++ {
			x := (float64(px) / width) - 0.5
			y := (float64(py) / height) - 0.5

			scene.Trace(x, y)
		}
	}
}
