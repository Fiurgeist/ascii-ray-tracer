package material

import (
	"math"

	"github.com/fiurgeist/ascii-ray-tracer/internal/color"
	"github.com/fiurgeist/ascii-ray-tracer/internal/vector"
)

type Finish struct {
	ambient    float64
	diffuse    float64
	shiny      float64
	reflection float64
}

func NewFinish(ambient, diffuse, shiny, reflection float64) Finish {
	return Finish{ambient: ambient, diffuse: diffuse, shiny: shiny, reflection: reflection}
}

func DefaultFinish() Finish {
	return NewFinish(0.1, 0.7, 0, 0)
}

func ShinyFinish() Finish {
	return NewFinish(0.1, 0.7, 0.5, 0.5)
}

func (f Finish) highlightFor(reflection, light vector.Vector, lightColor color.Color) color.Color {
	if f.shiny == 0 {
		return color.Black
	}

	intensity := reflection.Dot(light.Normalize())
	if intensity <= 0 {
		return color.Black
	}

	exp := 32 * f.shiny * f.shiny
	intensity = math.Pow(intensity, exp)
	return lightColor.Scale(f.shiny * intensity)
}
