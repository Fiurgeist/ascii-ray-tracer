package object

import (
	"math"

	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/material"
	"github.com/fiurgeist/ascii-ray-tracer/internal/ray"
	"github.com/fiurgeist/ascii-ray-tracer/internal/settings"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

var _ Object = (*Plane)(nil)

type Plane struct {
	normal   vector.Vector
	distance float64
	material material.Material
}

func NewPlane(normal vector.Vector, distance float64, material material.Material) Plane {
	return Plane{
		normal:   normal,
		distance: distance,
		material: material,
	}
}

func (p Plane) Color() color.Color { return p.material.Color() }

func (p Plane) ClosestDistanceAlongRay(ray ray.Ray) float64 {
	a := ray.Direction.Dot(p.normal)
	if a == 0 {
		return math.Inf(1)
	}

	b := p.normal.Dot(ray.Start.Add(p.normal.Scale(p.distance).Invert()))
	distance := -b / a
	if distance > settings.THRESHOLD {
		return distance
	}

	return math.Inf(1)
}

func (p Plane) NormalAt(point vector.Vector) vector.Vector {
	return p.normal
}
