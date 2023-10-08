package scene

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/camera"
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
)

type Scene struct {
	Background color.Color
	Camera     camera.Camera
}

func (s Scene) Trace(x float64, y float64) color.Color {
	return s.Camera.RayFor(x, y).Trace(s.Background)
}
