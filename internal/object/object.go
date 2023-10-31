package object

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/material"
	"github.com/fiurgeist/ascii-ray-tracer/internal/ray"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

type Object interface {
	ClosestDistanceAlongRay(ray ray.Ray) float64
	Material() material.Material
	NormalAt(point vector.Vector) vector.Vector
}
