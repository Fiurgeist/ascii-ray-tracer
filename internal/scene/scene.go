package scene

import (
	"math"

	"github.com/fiurgeist/ascii-ray-tracer/internal/camera"
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/object"
)

type Scene struct {
	Background color.Color
	Camera     camera.Camera
	Objects    []object.Object
}

func (s Scene) Trace(x float64, y float64) color.Color {
	var nearestObject object.Object
	shortestDistance := math.Inf(1)
	ray := s.Camera.RayFor(x, y)

	for _, obj := range s.Objects {
		if distance := obj.ClosestDistanceAlongRay(ray); distance < shortestDistance {
			shortestDistance = distance
			nearestObject = obj
			// log.Printf("%f - %v\n", shortestDistance, obj)
		}
	}

	if nearestObject == nil {
		return s.Background
	}
	// log.Printf("%f - %v\n", shortestDistance, nearestObject)
	return nearestObject.Color()
}
