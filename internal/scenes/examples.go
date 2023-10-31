package scenes

import (
	"github.com/fiurgeist/ascii-ray-tracer/internal/camera"
	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/light"
	"github.com/fiurgeist/ascii-ray-tracer/internal/material"
	"github.com/fiurgeist/ascii-ray-tracer/internal/object"
	"github.com/fiurgeist/ascii-ray-tracer/internal/scene"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

func ColoredSpheres() scene.Scene {
	camera := camera.NewCamera(vector.Vector{X: 0, Y: 2, Z: -8}, vector.Z)
	background := color.Grey
	objects := []object.Object{
		object.NewSphere(vector.Vector{X: -4, Y: 0, Z: 4}, 1, material.NewMaterial(color.Yellow)),
		object.NewSphere(vector.Vector{X: -2, Y: 0, Z: 2}, 1, material.NewMaterial(color.Red)),
		object.NewSphere(vector.Vector{X: 0, Y: 0, Z: 0}, 1, material.NewMaterial(color.Cyan)),
		object.NewSphere(vector.Vector{X: 2, Y: 0, Z: 2}, 1, material.NewMaterial(color.Green)),
		object.NewSphere(vector.Vector{X: 4, Y: 0, Z: 4}, 1, material.NewMaterial(color.Blue)),
	}
	lights := []light.Light{
		light.NewLight(vector.Vector{X: 5, Y: 10, Z: -5}, color.White),
	}
	return scene.Scene{Camera: camera, Background: background, Objects: objects, Lights: lights}
}

func AssortedObjects() scene.Scene {
	camera := camera.NewCamera(vector.Vector{X: -10, Y: 10, Z: -20}, vector.Vector{X: 0, Y: 4, Z: 0})
	background := color.Grey
	objects := []object.Object{
		object.NewPlane(vector.Y, 0, material.NewMaterial(color.White)),
		object.NewBox(vector.Vector{X: -2, Y: 0, Z: -2}, vector.Vector{X: 2, Y: 4, Z: 2}, material.NewMaterial(color.Red)),
		object.NewSphere(vector.Vector{X: 7, Y: 0, Z: 2}, 2, material.NewMaterial(color.Magenta)),
		object.NewSphere(vector.Vector{X: 6, Y: 1, Z: -4}, 1, material.NewMaterial(color.Yellow)),
		object.NewSphere(vector.Vector{X: -2, Y: 2, Z: 4}, 2, material.NewMaterial(color.Green)),
		object.NewSphere(vector.Vector{X: -4, Y: 4, Z: 10}, 4, material.NewMaterial(color.Blue)),
		object.NewSphere(vector.Vector{X: -3.2, Y: 1, Z: -1}, 1, material.NewMaterial(color.Cyan)),
	}
	lights := []light.Light{
		light.NewLight(vector.Vector{X: -30, Y: 25, Z: -12}, color.White),
	}
	return scene.Scene{Camera: camera, Background: background, Objects: objects, Lights: lights}
}
