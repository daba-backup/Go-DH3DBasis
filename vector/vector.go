package go_dh3dbasis_vector

import (
	"fmt"
	"math"
)

type Vector struct {
	x float32
	y float32
	z float32
}

func (v *Vector) String() string {
	str := fmt.Sprintf("(%g,%g,%g)", v.x, v.y, v.z)

	return str
}

func VGet(x float32, y float32, z float32) Vector {
	var v Vector

	v.x = x
	v.y = y
	v.z = z

	return v
}
func VAdd(lhs Vector, rhs Vector) Vector {
	var ret Vector

	ret.x = lhs.x + rhs.x
	ret.y = lhs.y + rhs.y
	ret.z = lhs.z + rhs.z

	return ret
}
func VSub(lhs Vector, rhs Vector) Vector {
	var ret Vector

	ret.x = lhs.x - rhs.x
	ret.y = lhs.y - rhs.y
	ret.z = lhs.z - rhs.z

	return ret
}
func VSize(v Vector) float32 {
	size := math.Sqrt(float64(v.x*v.x + v.y*v.y + v.z*v.z))
	return float32(size)
}
func VSquareSize(v Vector) float32 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}
func VNorm(v Vector) Vector {
	size := VSize(v)

	var ret Vector
	ret.x = v.x / size
	ret.y = v.y / size
	ret.z = v.z / size

	return ret
}
func VScale(v Vector, scale float32) Vector {
	var ret Vector

	ret.x = v.x * scale
	ret.y = v.y * scale
	ret.z = v.z * scale

	return ret
}
func VDot(lhs Vector, rhs Vector) float32 {
	return lhs.x*rhs.x + lhs.y*rhs.y + lhs.z*rhs.z
}
func VCross(lhs Vector, rhs Vector) Vector {
	var ret Vector

	ret.x = lhs.y*rhs.z - lhs.z*rhs.y
	ret.y = lhs.z*rhs.x - lhs.x*rhs.z
	ret.z = lhs.x*rhs.y - lhs.y*rhs.x

	return ret
}
func VAverage(v []Vector) Vector {
	var ret Vector
	ret.x = 0.0
	ret.y = 0.0
	ret.z = 0.0

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

	sin_th := v.y / d
	th = float32(math.Asin(float64(sin_th)))

	return th
}
func VAngleH(v Vector) float32 {
	xz_vec := VGet(v.x, 0.0, v.z)
	xz_length := VSize(xz_vec)
	if xz_length < 1.0E-8 {
		return 0.0
	}

	var th float32

	cos_th := v.x / xz_length
	th = float32(math.Acos(float64(cos_th)))

	if v.z >= 0.0 {
		th *= (-1.0)
	}

	return th
}
func VGetFromAngles(v_rotate float32, h_rotate float32) Vector {
	var ret Vector

	ret.x = float32(math.Cos(float64(h_rotate)))
	ret.y = float32(math.Sin(float64(v_rotate)))
	ret.z = float32(-math.Sin(float64(h_rotate)))

	ret = VNorm(ret)

	return ret
}
