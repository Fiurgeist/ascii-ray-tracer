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
	Location  vector.Vector
	LookAt    vector.Vector
	direction vector.Vector
	right     vector.Vector
	down      vector.Vector
}

func NewCamera(location vector.Vector, lookAt vector.Vector) *Camera {
	c := &Camera{LookAt: lookAt}
	c.UpdateLocation(location)

	return c
}

func (c *Camera) UpdateLocation(location vector.Vector) {
	c.Location = location
	c.direction = c.LookAt.Substract(c.Location).Normalize()
	c.right = vector.Y.Cross(c.direction).Normalize().Scale(width / 2)
	c.down = c.right.Cross(c.direction).Normalize().Scale(height / 2)
}

func (c *Camera) RayFor(x float64, y float64) ray.Ray {
	xRay := c.right.Scale(x)
	yRay := c.down.Scale(y)
	rayDirection := c.direction.Add(xRay).Add(yRay)
	return ray.NewRay(c.Location, rayDirection)
}
