package object

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/ray"
)

type Object interface {
	ClosestDistanceAlongRay(ray ray.Ray) float64
	Color() color.Color
}
