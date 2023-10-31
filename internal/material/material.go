package material

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

type Material struct {
	color  color.Color
	finish Finish
}

func NewMaterial(color color.Color, finish Finish) Material {
	return Material{color: color, finish: finish}
}

func (m Material) AmbientColor() color.Color {
	return m.color.Scale(m.finish.ambient)
}

func (m Material) DiffuseColor() color.Color {
	return m.color.Scale(m.finish.diffuse)
}

func (m Material) HighlightFor(reflection, light vector.Vector, lightColor color.Color) color.Color {
	return m.finish.highlightFor(reflection, light, lightColor)
}

func (m Material) Reflection() float64 {
	return m.finish.reflection
}
