package ray

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

type Ray struct {
	Start     vector.Vector
	Direction vector.Vector
}

func NewRay(start vector.Vector, direction vector.Vector) Ray {
	return Ray{Start: start, Direction: direction.Normalize()}
}
