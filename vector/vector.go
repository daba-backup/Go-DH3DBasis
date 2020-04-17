package go_dh3dbasis_vector

import (
	"fmt"
	"math"
)

type Vector struct {
	X float32
	Y float32
	Z float32
}

func (v *Vector) String() string {
	str := fmt.Sprintf("(%g,%g,%g)", v.X, v.Y, v.Z)

	return str
}

func VGet(x float32, y float32, z float32) Vector {
	var v Vector

	v.X = x
	v.Y = y
	v.Z = z

	return v
}
func VAdd(lhs Vector, rhs Vector) Vector {
	var ret Vector

	ret.X = lhs.X + rhs.X
	ret.Y = lhs.Y + rhs.Y
	ret.Z = lhs.Z + rhs.Z

	return ret
}
func VSub(lhs Vector, rhs Vector) Vector {
	var ret Vector

	ret.X = lhs.X - rhs.X
	ret.Y = lhs.Y - rhs.Y
	ret.Z = lhs.Z - rhs.Z

	return ret
}
func VSize(v Vector) float32 {
	size := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
	return float32(size)
}
func VSquareSize(v Vector) float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
func VNorm(v Vector) Vector {
	size := VSize(v)

	var ret Vector
	ret.X = v.X / size
	ret.Y = v.Y / size
	ret.Z = v.Z / size

	return ret
}
func VScale(v Vector, scale float32) Vector {
	var ret Vector

	ret.X = v.X * scale
	ret.Y = v.Y * scale
	ret.Z = v.Z * scale

	return ret
}
func VDot(lhs Vector, rhs Vector) float32 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}
func VCross(lhs Vector, rhs Vector) Vector {
	var ret Vector

	ret.X = lhs.Y*rhs.Z - lhs.Z*rhs.Y
	ret.Y = lhs.Z*rhs.X - lhs.X*rhs.Z
	ret.Z = lhs.X*rhs.Y - lhs.Y*rhs.X

	return ret
}
func VAverage(v []Vector) Vector {
	var ret Vector
	ret.X = 0.0
	ret.Y = 0.0
	ret.Z = 0.0

	v_num := len(v)
	for i := 0; i < v_num; i++ {
		ret = VAdd(ret, v[i])
	}
	ret = VScale(ret, 1.0/float32(v_num))

	return ret
}
func VAngleV(v Vector) float32 {
	d := VSize(v)
	if d < 1.0E-8 {
		return 0.0
	}

	var th float32

	sin_th := v.Y / d
	th = float32(math.Asin(float64(sin_th)))

	return th
}
func VAngleH(v Vector) float32 {
	xz_vec := VGet(v.X, 0.0, v.Z)
	xz_length := VSize(xz_vec)
	if xz_length < 1.0E-8 {
		return 0.0
	}

	var th float32

	cos_th := v.X / xz_length
	th = float32(math.Acos(float64(cos_th)))

	if v.Z >= 0.0 {
		th *= (-1.0)
	}

	return th
}
func VGetFromAngles(v_rotate float32, h_rotate float32) Vector {
	var ret Vector

	ret.X = float32(math.Cos(float64(h_rotate)))
	ret.Y = float32(math.Sin(float64(v_rotate)))
	ret.Z = float32(-math.Sin(float64(h_rotate)))

	ret = VNorm(ret)

	return ret
}
