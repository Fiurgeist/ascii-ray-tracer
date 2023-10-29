package scene

import (
	"math"

	"github.com/fiurgeist/ascii-ray-tracer/internal/camera"
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/light"
	"github.com/fiurgeist/ascii-ray-tracer/internal/object"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

type Scene struct {
	Background color.Color
	Camera     camera.Camera
	Objects    []object.Object
	Lights     []light.Light
}

func (s Scene) Trace(x float64, y float64) color.Color {
	var nearestObject object.Object
	shortestDistance := math.Inf(1)
	ray := s.Camera.RayFor(x, y)

	for _, obj := range s.Objects {
		if distance := obj.ClosestDistanceAlongRay(ray); distance < shortestDistance {
			shortestDistance = distance
			nearestObject = obj
		}
	}

	if nearestObject == nil {
		return s.Background
	}
	point := ray.PointAtDistance(shortestDistance)
	return s.colorAt(point, nearestObject)
}

func (s Scene) colorAt(point vector.Vector, object object.Object) color.Color {
	normal := object.NormalAt(point)

	color := color.Black
	for _, light := range s.Lights {
		lightVector := light.Position().Substract(point)
		brightness := normal.Dot(lightVector.Normalize())
		if brightness <= 0 {
			continue
		}
		illumination := object.Color().Multiply(light.Color()).Scale(brightness)
		color = color.Add(illumination)
	}
	return color
}
