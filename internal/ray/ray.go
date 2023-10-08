package ray

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

type Ray struct {
	Start     vector.Vector
	Direction vector.Vector
}

func (r Ray) Trace(background color.Color) color.Color {
	return background
}
