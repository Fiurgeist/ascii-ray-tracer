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

func (r Ray) PointAtDistance(distance float64) vector.Vector {
	return r.Start.Add(r.Direction.Scale(distance))
}

func (r Ray) Reflect(normal vector.Vector) vector.Vector {
	normalSquid := normal.Squid()
	if normalSquid == 0 {
		return r.Direction
	}

	return r.Direction.Substract(normal.Scale(2 * r.Direction.Dot(normal) / normalSquid))
}
