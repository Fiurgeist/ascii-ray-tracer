package vector

import "math"

var (
	O  = Vector{X: 0, Y: 0, Z: 0}
	U  = Vector{X: 1, Y: 1, Z: 1}
	X  = Vector{X: 1, Y: 0, Z: 0}
	Y  = Vector{X: 0, Y: 1, Z: 0}
	Z  = Vector{X: 0, Y: 0, Z: 1}
	XI = Vector{X: -1, Y: 0, Z: 0}
	YI = Vector{X: 0, Y: -1, Z: 0}
	ZI = Vector{X: 0, Y: 0, Z: -1}
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v Vector) Squid() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.Squid())
}

func (v Vector) Invert() Vector {
	return Vector{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Vector) Scale(factor float64) Vector {
	return Vector{
		X: v.X * factor,
		Y: v.Y * factor,
		Z: v.Z * factor,
	}
}

func (v Vector) Normalize() Vector {
	length := v.Length()
	return Vector{
		X: v.X / length,
		Y: v.Y / length,
		Z: v.Z / length,
	}
}

func (v Vector) Dot(other Vector) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vector) Cross(other Vector) Vector {
	return Vector{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vector) Substract(other Vector) Vector {
	return Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}
