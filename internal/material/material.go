package material

import "github.com/fiurgeist/ascii-ray-tracer/internal/color"

type Material struct {
	color color.Color
}

func NewMaterial(color color.Color) Material {
	return Material{color: color}
}

func (m Material) Color() color.Color {
	return m.color
}
