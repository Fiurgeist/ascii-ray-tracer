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
