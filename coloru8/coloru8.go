package go_dh3dbasis_coloru8

type ColorU8 struct {
	r float32
	g float32
	b float32
	a float32
}

func GetColorU8FromFloat32Components(r float32, g float32, b float32, a float32) ColorU8 {
	var color ColorU8

	color.r = r
	color.g = g
	color.b = b
	color.a = a

	return color
}
func GetColorU8FromIntComponents(r int, g int, b int, a int) ColorU8 {
	var color ColorU8

	color.r = float32(r) / 255.0
	color.g = float32(g) / 255.0
	color.b = float32(b) / 255.0
	color.a = float32(a) / 255.0

	return color
}
