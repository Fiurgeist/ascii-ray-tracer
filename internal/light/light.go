package light

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

type Light struct {
	position vector.Vector
	color    color.Color
}

func NewLight(position vector.Vector, color color.Color) Light {
	return Light{position: position, color: color}
}

func (l Light) Position() vector.Vector {
	return l.position
}

func (l Light) Color() color.Color {
	return l.color
}
