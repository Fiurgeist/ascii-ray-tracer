package renderer

import "github.com/fiurgeist/ascii-ray-tracer/internal/scene"

type Renderer interface {
	Render(scene scene.Scene, processor string)
}
