package camera

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/ray"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

const (
	width  float64 = 4
	height float64 = 9 / 4
)

type Camera struct {
	location  vector.Vector
	lookAt    vector.Vector
	direction vector.Vector
	right     vector.Vector
	down      vector.Vector
}

func NewCamera(location vector.Vector, lookAt vector.Vector) Camera {
	direction := lookAt.Substract(location).Normalize()
	right := vector.Y.Cross(direction).Normalize().Scale(width / 2)
	down := right.Cross(direction).Normalize().Scale(height / 2)
	return Camera{
		location:  location,
		lookAt:    lookAt,
		direction: direction,
		right:     right,
		down:      down,
	}
}

func (c Camera) RayFor(x float64, y float64) ray.Ray {
	xRay := c.right.Scale(x)
	yRay := c.down.Scale(y)
	rayDirection := c.direction.Add(xRay).Add(yRay)
	return ray.NewRay(c.location, rayDirection)
}
