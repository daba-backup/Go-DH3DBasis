package coloru8

import (
	"fmt"
)

type ColorU8 struct {
	R float32
	G float32
	B float32
	A float32
}

func (c *ColorU8) String() string {
	str := fmt.Sprintf("(%g,%g,%g,%g)", c.R, c.G, c.B, c.A)

	return str
}

func GetColorU8FromFloat32Components(r float32, g float32, b float32, a float32) ColorU8 {
	var color ColorU8

	color.R = r
	color.G = g
	color.B = b
	color.A = a

	return color
}
func GetColorU8FromIntComponents(r int, g int, b int, a int) ColorU8 {
	var color ColorU8

	color.R = float32(r) / 255.0
	color.G = float32(g) / 255.0
	color.B = float32(b) / 255.0
	color.A = float32(a) / 255.0

	return color
}
