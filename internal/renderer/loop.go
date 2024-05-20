package renderer

import (
	"fmt"
	"math"
	"time"

	"github.com/fiurgeist/ascii-ray-tracer/internal/scene"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

const speed = 0.5 // speed in radians per second

var _ Renderer = (*GameLoopRenderer)(nil)

type GameLoopRenderer struct {
	Renderer ConsoleRenderer

	angle  float64
	radius float64
}

func (r *GameLoopRenderer) Render(scene scene.Scene) {
	rVec := scene.Camera.Location.Substract(scene.Camera.LookAt)
	r.radius = math.Sqrt(rVec.X*rVec.X + rVec.Z*rVec.Z)
	r.angle = math.Asin(rVec.Z / r.radius)

	fmt.Printf("%s2J", esc)

	delta := 0.0

	for {
		start := time.Now()

		r.update(delta, scene)
		r.Renderer.render(scene)

		delta = time.Since(start).Seconds()
		fmt.Printf("%s%d;%dH", esc, r.Renderer.Height/2+1, 1)
		fmt.Printf("FPS: %.1f\n", 1/delta)
	}
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
