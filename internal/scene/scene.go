package scene

import (
	"math"

	"github.com/fiurgeist/ascii-ray-tracer/internal/camera"
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/light"
	"github.com/fiurgeist/ascii-ray-tracer/internal/object"
	"github.com/fiurgeist/ascii-ray-tracer/internal/ray"
	"github.com/fiurgeist/ascii-ray-tracer/internal/settings"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

type Scene struct {
	Background color.Color
	Camera     *camera.Camera
	Objects    []object.Object
	Lights     []light.Light
}

func (s Scene) Trace(x float64, y float64) color.Color {
	ray := s.Camera.RayFor(x, y)
	return s.traceRay(ray, 0)
}

func (s Scene) traceRay(ray ray.Ray, depth int) color.Color {
	if depth > settings.MAX_DEPTH {
		return color.Black
	}
	var nearestObject object.Object
	shortestDistance := math.Inf(1)
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
	return s.colorAt(point, nearestObject, ray, depth+1)
}

func (s Scene) colorAt(point vector.Vector, object object.Object, ray ray.Ray, depth int) color.Color {
	normal := object.NormalAt(point)
	color := object.Material().AmbientColor()
	reflectionVector := ray.Reflect(normal)
	reflection := s.reflectionFor(object, point, reflectionVector, depth)
	color = color.Add(reflection)
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
	length := lightVector.Length()

	for _, object := range s.Objects {
		if object.ClosestDistanceAlongRay(ray.NewRay(point, lightVector)) <= length {
			return true
		}
	}
	return false
}

func (s Scene) reflectionFor(object object.Object, point vector.Vector, reflectionVector vector.Vector, depth int) color.Color {
	if object.Material().Reflection() == 0 {
		return color.Black
	}
	reflectedRay := ray.NewRay(point, reflectionVector)
	reflectedColor := s.traceRay(reflectedRay, depth)
	return reflectedColor.Scale(object.Material().Reflection())
}
