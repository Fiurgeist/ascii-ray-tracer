package object

import (
	"log"
	"math"

	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/material"
	"github.com/fiurgeist/ascii-ray-tracer/internal/ray"
	"github.com/fiurgeist/ascii-ray-tracer/internal/settings"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

var _ Object = (*Box)(nil)

type Box struct {
	lowerCorner vector.Vector
	upperCorner vector.Vector
	material    material.Material
}

func NewBox(corner1 vector.Vector, corner2 vector.Vector, material material.Material) Box {
	lowerCorner := vector.Vector{X: math.Min(corner1.X, corner2.X), Y: math.Min(corner1.Y, corner2.Y), Z: math.Min(corner1.Z, corner2.Z)}
	upperCorner := vector.Vector{X: math.Max(corner1.X, corner2.X), Y: math.Max(corner1.Y, corner2.Y), Z: math.Max(corner1.Z, corner2.Z)}

	return Box{
		lowerCorner: lowerCorner,
		upperCorner: upperCorner,
		material:    material,
	}
}

func (b Box) Color() color.Color { return b.material.Color() }

func (b Box) ClosestDistanceAlongRay(ray ray.Ray) float64 {
	intersections := b.intersectionsOnAxis(ray, getX, getY, getZ)
	intersections = append(intersections, b.intersectionsOnAxis(ray, getY, getX, getZ)...)
	intersections = append(intersections, b.intersectionsOnAxis(ray, getZ, getX, getY)...)

	shortest := math.Inf(1)
	for _, distance := range intersections {
		if distance > settings.THRESHOLD && distance < shortest {
			shortest = distance
		}
	}

	return shortest
}

func (b Box) NormalAt(point vector.Vector) vector.Vector {
	if math.Abs(b.lowerCorner.X-point.X) < settings.THRESHOLD {
		return vector.XI
	}
	if math.Abs(b.upperCorner.X-point.X) < settings.THRESHOLD {
		return vector.X
	}
	if math.Abs(b.lowerCorner.Y-point.Y) < settings.THRESHOLD {
		return vector.YI
	}
	if math.Abs(b.upperCorner.Y-point.Y) < settings.THRESHOLD {
		return vector.Y
	}
	if math.Abs(b.lowerCorner.Z-point.Z) < settings.THRESHOLD {
		return vector.ZI
	}
	if math.Abs(b.upperCorner.Z-point.Z) < settings.THRESHOLD {
		return vector.Z
	}

	log.Panicf("THRESHOLD to small - Box: %v; Point: %v\n", b, point)
	return vector.O
}

func (b Box) intersectionsOnAxis(
	ray ray.Ray,
	axis func(vector vector.Vector) float64,
	otherAxis1 func(vector vector.Vector) float64,
	otherAxis2 func(vector vector.Vector) float64,
) []float64 {
	if axis(ray.Direction) == 0 {
		return []float64{}
	}

	var intersections []float64
	if intersection, contained := b.intersectionForVertex(b.lowerCorner, ray, axis, otherAxis1, otherAxis2); contained {
		intersections = append(intersections, intersection)
	}
	if intersection, contained := b.intersectionForVertex(b.upperCorner, ray, axis, otherAxis1, otherAxis2); contained {
		intersections = append(intersections, intersection)
	}

	return intersections
}

func (b Box) intersectionForVertex(
	vertex vector.Vector,
	ray ray.Ray,
	axis func(vector vector.Vector) float64,
	otherAxis1 func(vector vector.Vector) float64,
	otherAxis2 func(vector vector.Vector) float64,
) (float64, bool) {
	intersection := (axis(vertex) - axis(ray.Start)) / axis(ray.Direction)
	point := ray.Start.Add(ray.Direction.Scale(intersection))

	return intersection, b.contains(point, otherAxis1) && b.contains(point, otherAxis2)
}

func (b Box) contains(point vector.Vector, axis func(vector vector.Vector) float64) bool {
	return axis(b.lowerCorner) < axis(point) && axis(point) < axis(b.upperCorner)
}

func getX(vector vector.Vector) float64 { return vector.X }
func getY(vector vector.Vector) float64 { return vector.Y }
func getZ(vector vector.Vector) float64 { return vector.Z }
