package object

import (
	"math"

	"github.com/fiurgeist/ascii-ray-tracer/internal/material"
	"github.com/fiurgeist/ascii-ray-tracer/internal/ray"
	"github.com/fiurgeist/ascii-ray-tracer/internal/settings"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

var _ Object = (*Sphere)(nil)

type Sphere struct {
	center   vector.Vector
	radius   float64
	material material.Material
}

func NewSphere(center vector.Vector, radius float64, material material.Material) Sphere {
	return Sphere{
		center:   center,
		radius:   radius,
		material: material,
	}
}

func (s Sphere) Material() material.Material { return s.material }

func (s Sphere) ClosestDistanceAlongRay(ray ray.Ray) float64 {
	os := ray.Start.Substract(s.center)
	b := 2 * os.Dot(ray.Direction)
	c := os.Squid() - s.radius*s.radius

	discriminant := b*b - 4*c
	if discriminant < 0 {
		return math.Inf(1)
	}
	if discriminant == 0 {
		return -b / 2
	}

	distance1 := (-b - math.Sqrt(discriminant)) / 2
	distance2 := (-b + math.Sqrt(discriminant)) / 2
	if distance1 > settings.THRESHOLD && distance1 < distance2 {
		return distance1
	}
	if distance2 > settings.THRESHOLD {
		return distance2
	}

	return math.Inf(1)
}

func (s Sphere) NormalAt(point vector.Vector) vector.Vector {
	return point.Add(s.center.Invert()).Normalize()
}
