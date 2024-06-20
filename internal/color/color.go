package color

var (
	White   = Color{r: 255, g: 255, b: 255}
	Black   = Color{r: 0, g: 0, b: 0}
	Grey    = Color{r: 127, g: 127, b: 127}
	Red     = Color{r: 255, g: 0, b: 0}
	Green   = Color{r: 0, g: 255, b: 0}
	Blue    = Color{r: 0, g: 0, b: 255}
	Yellow  = Color{r: 255, g: 255, b: 0}
	Magenta = Color{r: 255, g: 0, b: 255}
	Cyan    = Color{r: 0, g: 255, b: 255}
)

const max = uint8(255)

type Color struct {
	r uint8
	g uint8
	b uint8
}

func NewColor(r, g, b uint8) Color {
	return Color{r: r, b: b, g: g}
}

func clampedAdd(a uint8, b uint8) uint8 {
	sum := uint16(a) + uint16(b)
	if sum > 255 {
		return max
	}
	return uint8(sum)
}

func (c Color) Add(other Color) Color {
	return Color{
		r: clampedAdd(c.r, other.r),
		g: clampedAdd(c.g, other.g),
		b: clampedAdd(c.b, other.b),
	}
}

func clampedScale(val uint8, factor float64) uint8 {
	scaled := float64(val) * factor
	if scaled > 255 {
		return max
	}
	return uint8(scaled)
}

func (c Color) Scale(factor float64) Color {
	return Color{
		r: clampedScale(c.r, factor),
		g: clampedScale(c.g, factor),
		b: clampedScale(c.b, factor),
	}
}

func (c Color) Multiply(other Color) Color {
	return Color{
		r: uint8(float64(c.r) * float64(other.r) / 255.0),
		g: uint8(float64(c.g) * float64(other.g) / 255.0),
		b: uint8(float64(c.b) * float64(other.b) / 255.0),
	}
}

func (c Color) R() uint8 { return c.r }
func (c Color) G() uint8 { return c.g }
func (c Color) B() uint8 { return c.b }
