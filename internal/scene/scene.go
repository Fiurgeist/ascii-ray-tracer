package scene

import (
	"math"

	"github.com/fiurgeist/ascii-ray-tracer/internal/camera"
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/light"
	"github.com/fiurgeist/ascii-ray-tracer/internal/object"
	"github.com/fiurgeist/ascii-ray-tracer/internal/ray"
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
	return s.colorAt(point, nearestObject, ray)
}

func (s Scene) colorAt(point vector.Vector, object object.Object, ray ray.Ray) color.Color {
	normal := object.NormalAt(point)
	color := object.Material().AmbientColor()
	reflectionVector := ray.Reflect(normal)
	for _, light := range s.Lights {
		lightVector := light.Position().Substract(point)
		if s.inShadow(point, lightVector) {
			continue
		}

		brightness := normal.Dot(lightVector.Normalize())
		if brightness <= 0 {
			continue
		}

		illumination := object.Material().DiffuseColor().Multiply(light.Color()).Scale(brightness)
		color = color.Add(illumination)

		highlight := object.Material().HighlightFor(reflectionVector, lightVector, light.Color())
		color = color.Add(highlight)
	}
	return color
}

func (s Scene) inShadow(point vector.Vector, lightVector vector.Vector) bool {
	for _, object := range s.Objects {
		if object.ClosestDistanceAlongRay(ray.NewRay(point, lightVector)) <= lightVector.Length() {
			return true
		}
	}
	return false
}
