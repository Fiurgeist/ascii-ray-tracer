package renderer

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/fiurgeist/ascii-ray-tracer/internal/scene"
	"github.com/fiurgeist/ascii-ray-tracer/internal/shader"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

const speed = 0.5 // speed in radians per second

var _ Renderer = (*GameLoopRenderer)(nil)

type GameLoopRenderer struct {
	Renderer ConsoleRenderer

	wg      sync.WaitGroup
	running bool
	angle   float64
	radius  float64
}

func (r *GameLoopRenderer) Render(scene scene.Scene, processor string) {
	r.wg.Add(1)
	defer r.wg.Done()

	if processor == "gpu" {
		r.Renderer.shader = &shader.Shader{Width: int32(r.Renderer.Width), Height: int32(r.Renderer.Height)}

		r.Renderer.shader.Init(scene)
		defer r.Renderer.shader.Delete()
	}

	rVec := scene.Camera.Location.Substract(scene.Camera.LookAt)
	r.radius = math.Sqrt(rVec.X*rVec.X + rVec.Z*rVec.Z)
	r.angle = math.Asin(rVec.Z / r.radius)
	r.running = true

	fmt.Printf("%s2J", esc)

	delta := 0.0

	for r.running {
		start := time.Now()

		r.update(delta, scene)

		if processor == "gpu" {
			r.Renderer.gpuRender(scene)
		} else {
			r.Renderer.cpuRender(scene)
		}

		delta = time.Since(start).Seconds()
		fmt.Printf("%s%d;%dH", esc, r.Renderer.Height/2+1, 1)
		fmt.Printf("FPS: %.1f\n", 1/delta)
	}
}

func (r *GameLoopRenderer) Stop() {
	r.running = false
	r.wg.Wait()

	ResetConsole()
}

func (r *GameLoopRenderer) update(delta float64, scene scene.Scene) {
	// there's probably a better way updating the position using proper vector math... but I'm not seeing it right now
	r.angle += speed * delta

	scene.Camera.UpdateLocation(vector.Vector{
		X: math.Cos(r.angle) * r.radius,
		Z: math.Sin(r.angle) * r.radius,
		Y: scene.Camera.Location.Y,
	})
}
