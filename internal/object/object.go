package object

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/ray"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

type Object interface {
	ClosestDistanceAlongRay(ray ray.Ray) float64
	Color() color.Color
	NormalAt(point vector.Vector) vector.Vector
}
