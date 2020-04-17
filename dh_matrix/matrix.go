package dh_matrix

import (
	"fmt"
	"math"

	"github.com/dabasan/go-dh3dbasis/dh_vector"
)

type Matrix struct {
	M [4][4]float32
}

func (m *Matrix) String() string {
	str := ""

	for i := 0; i < 4; i++ {
		str += fmt.Sprintf("%g %g %g %g", m.M[i][0], m.M[i][1], m.M[i][2], m.M[i][3])
		if i != 3 {
			str += "\n"
		}
	}

	return str
}

func MMult(m1 Matrix, m2 Matrix) Matrix {
	var ret Matrix

	var value float32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			value = 0.0

			for k := 0; k < 4; k++ {
				value += m1.M[i][k] * m2.M[k][j]
			}

			ret.M[i][j] = value
		}
	}

	return ret
}
func MGetIdent() Matrix {
	var ret Matrix

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret.M[i][j] = 0.0
		}
	}
	for i := 0; i < 4; i++ {
		ret.M[i][i] = 1.0
	}

	return ret
}
func MGetScale(scale dh_vector.Vector) Matrix {
	var ret Matrix

	ret.M[0][0] = scale.X
	ret.M[0][1] = 0.0
	ret.M[0][2] = 0.0
	ret.M[0][3] = 0.0
	ret.M[1][0] = 0.0
	ret.M[1][1] = scale.Y
	ret.M[1][2] = 0.0
	ret.M[1][3] = 0.0
	ret.M[2][0] = 0.0
	ret.M[2][1] = 0.0
	ret.M[2][2] = scale.Z
	ret.M[2][3] = 0.0
	ret.M[3][0] = 0.0
	ret.M[3][1] = 0.0
	ret.M[3][2] = 0.0
	ret.M[3][3] = 1.0

	return ret
}
func MGetTranslate(translate dh_vector.Vector) Matrix {
	var ret Matrix

	ret.M[0][0] = 1.0
	ret.M[0][1] = 0.0
	ret.M[0][2] = 0.0
	ret.M[0][3] = translate.X
	ret.M[1][0] = 0.0
	ret.M[1][1] = 1.0
	ret.M[1][2] = 0.0
	ret.M[1][3] = translate.Y
	ret.M[2][0] = 0.0
	ret.M[2][1] = 0.0
	ret.M[2][2] = 1.0
	ret.M[2][3] = translate.Z
	ret.M[3][0] = 0.0
	ret.M[3][1] = 0.0
	ret.M[3][2] = 0.0
	ret.M[3][3] = 1.0

	return ret
}
func MGetRotX(th float32) Matrix {
	var ret Matrix

	ret.M[0][0] = 1.0
	ret.M[0][1] = 0.0
	ret.M[0][2] = 0.0
	ret.M[0][3] = 0.0
	ret.M[1][0] = 0.0
	ret.M[1][1] = float32(math.Cos(float64(th)))
	ret.M[1][2] = float32(-math.Sin(float64(th)))
	ret.M[1][3] = 0.0
	ret.M[2][0] = 0.0
	ret.M[2][1] = float32(math.Sin(float64(th)))
	ret.M[2][2] = float32(math.Cos(float64(th)))
	ret.M[2][3] = 0.0
	ret.M[3][0] = 0.0
	ret.M[3][1] = 0.0
	ret.M[3][2] = 0.0
	ret.M[3][3] = 1.0

	return ret
}
func MGetRotY(th float32) Matrix {
	var ret Matrix

	ret.M[0][0] = float32(math.Cos(float64(th)))
	ret.M[0][1] = 0.0
	ret.M[0][2] = float32(math.Sin(float64(th)))
	ret.M[0][3] = 0.0
	ret.M[1][0] = 0.0
	ret.M[1][1] = 1.0
	ret.M[1][2] = 0.0
	ret.M[1][3] = 0.0
	ret.M[2][0] = float32(-math.Sin(float64(th)))
	ret.M[2][1] = 0.0
	ret.M[2][2] = float32(math.Cos(float64(th)))
	ret.M[2][3] = 0.0
	ret.M[3][0] = 0.0
	ret.M[3][1] = 0.0
	ret.M[3][2] = 0.0
	ret.M[3][3] = 1.0

	return ret
}
func MGetRotZ(th float32) Matrix {
	var ret Matrix

	ret.M[0][0] = float32(math.Cos(float64(th)))
	ret.M[0][1] = float32(-math.Sin(float64(th)))
	ret.M[0][2] = 0.0
	ret.M[0][3] = 0.0
	ret.M[1][0] = float32(math.Sin(float64(th)))
	ret.M[1][1] = float32(math.Cos(float64(th)))
	ret.M[1][2] = 0.0
	ret.M[1][3] = 0.0
	ret.M[2][0] = 0.0
	ret.M[2][1] = 0.0
	ret.M[2][2] = 1.0
	ret.M[2][3] = 0.0
	ret.M[3][0] = 0.0
	ret.M[3][1] = 0.0
	ret.M[3][2] = 0.0
	ret.M[3][3] = 1.0

	return ret
}
func MGetRotAxis(axis dh_vector.Vector, th float32) Matrix {
	var ret Matrix

	cos_th := float32(math.Cos(float64(th)))
	sin_th := float32(math.Sin(float64(th)))

	nx := axis.X
	ny := axis.Y
	nz := axis.Z

	ret.M[0][0] = cos_th * nx * nx * (1.0 - cos_th)
	ret.M[0][1] = nx*ny*(1.0-cos_th) - nz*sin_th
	ret.M[0][2] = nx*nz*(1.0-cos_th) + ny*sin_th
	ret.M[0][3] = 0.0
	ret.M[1][0] = ny*nx*(1.0-cos_th) + nz*sin_th
	ret.M[1][1] = cos_th + ny*ny*(1.0-cos_th)
	ret.M[1][2] = ny*nz*(1.0-cos_th) - ny*sin_th
	ret.M[1][3] = 0.0
	ret.M[2][0] = nz*nx*(1.0-cos_th) - ny*sin_th
	ret.M[2][1] = nz*ny*(1.0-cos_th) + nx*sin_th
	ret.M[2][2] = cos_th + nz*nz*(1.0-cos_th)
	ret.M[2][3] = 0.0
	ret.M[3][0] = 0.0
	ret.M[3][1] = 0.0
	ret.M[3][2] = 0.0
	ret.M[3][3] = 1.0

	return ret
}
func MTranspose(m Matrix) Matrix {
	var ret Matrix

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret.M[i][j] = m.M[j][i]
		}
	}

	return ret
}
func MGetMagic() Matrix {
	var ret Matrix

	ret.M[0][0] = 16.0
	ret.M[0][1] = 2.0
	ret.M[0][2] = 3.0
	ret.M[0][3] = 13.0
	ret.M[1][0] = 5.0
	ret.M[1][1] = 11.0
	ret.M[1][2] = 10.0
	ret.M[1][3] = 8.0
	ret.M[2][0] = 9.0
	ret.M[2][1] = 7.0
	ret.M[2][2] = 6.0
	ret.M[2][3] = 12.0
	ret.M[3][0] = 4.0
	ret.M[3][1] = 14.0
	ret.M[3][2] = 15.0
	ret.M[3][3] = 1.0

	return ret
}
func MGetDet(m Matrix) float32 {
	a00 := m.M[0][0]
	a01 := m.M[0][1]
	a02 := m.M[0][2]
	a03 := m.M[0][3]
	a10 := m.M[1][0]
	a11 := m.M[1][1]
	a12 := m.M[1][2]
	a13 := m.M[1][3]
	a20 := m.M[2][0]
	a21 := m.M[2][1]
	a22 := m.M[2][2]
	a23 := m.M[2][3]
	a30 := m.M[3][0]
	a31 := m.M[3][1]
	a32 := m.M[3][2]
	a33 := m.M[3][3]

	term0 := a00 * (a11*a22*a33 + a12*a23*a31 + a13*a21*a32 - a13*a22*a31 - a12*a21*a33 - a11*a23*a32)
	term1 := -a10 * (a01*a22*a33 + a02*a23*a31 + a03*a21*a32 - a03*a22*a31 - a02*a21*a33 - a01*a23*a32)
	term2 := a20 * (a01*a12*a33 + a02*a13*a31 + a03*a11*a32 - a03*a12*a31 - a02*a11*a33 - a01*a13*a32)
	term3 := -a30 * (a01*a12*a23 + a02*a13*a21 + a03*a11*a22 - a03*a12*a21 - a02*a11*a23 - a01*a13*a22)

	ret := term0 + term1 + term2 + term3

	return ret
}
func MInverse(m Matrix) Matrix {
	var ret Matrix

	a11 := m.M[0][0]
	a12 := m.M[0][1]
	a13 := m.M[0][2]
	a14 := m.M[0][3]
	a21 := m.M[1][0]
	a22 := m.M[1][1]
	a23 := m.M[1][2]
	a24 := m.M[1][3]
	a31 := m.M[2][0]
	a32 := m.M[2][1]
	a33 := m.M[2][2]
	a34 := m.M[2][3]
	a41 := m.M[3][0]
	a42 := m.M[3][1]
	a43 := m.M[3][2]
	a44 := m.M[3][3]

	var cofactor Matrix
	cofactor.M[0][0] = a22*a33*a44 + a23*a34*a42 + a24*a32*a43 - a24*a33*a42 - a23*a32*a44 - a22*a34*a43
	cofactor.M[0][1] = -a12*a33*a44 - a13*a34*a42 - a14*a32*a43 + a14*a33*a42 + a13*a32*a44 + a12*a34*a43
	cofactor.M[0][2] = a12*a23*a44 + a13*a24*a42 + a14*a22*a43 - a14*a23*a42 - a13*a22*a44 - a12*a24*a43
	cofactor.M[0][3] = -a12*a23*a34 - a13*a24*a32 - a14*a22*a33 + a14*a23*a32 + a13*a22*a34 + a12*a24*a33
	cofactor.M[1][0] = -a21*a33*a44 - a23*a34*a41 - a24*a31*a43 + a24*a33*a41 + a23*a31*a44 + a21*a34*a43
	cofactor.M[1][1] = a11*a33*a44 + a13*a34*a41 + a14*a31*a43 - a14*a33*a41 - a13*a31*a44 - a11*a34*a43
	cofactor.M[1][2] = -a11*a23*a44 - a13*a24*a41 - a14*a21*a43 + a14*a23*a41 + a13*a21*a44 + a11*a24*a43
	cofactor.M[1][3] = a11*a23*a34 + a13*a24*a31 + a14*a21*a33 - a14*a23*a31 - a13*a21*a34 - a11*a24*a33
	cofactor.M[2][0] = a21*a32*a44 + a22*a34*a41 + a24*a31*a42 - a24*a32*a41 - a22*a31*a44 - a21*a34*a42
	cofactor.M[2][1] = -a11*a32*a44 - a12*a34*a41 - a14*a31*a42 + a14*a32*a41 + a12*a31*a44 + a11*a34*a42
	cofactor.M[2][2] = a11*a22*a44 + a12*a24*a41 + a14*a21*a42 - a14*a22*a41 - a12*a21*a44 - a11*a24*a42
	cofactor.M[2][3] = -a11*a22*a34 - a12*a24*a31 - a14*a21*a32 + a14*a22*a31 + a12*a21*a34 + a11*a24*a32
	cofactor.M[3][0] = -a21*a32*a43 - a22*a33*a41 - a23*a31*a42 + a23*a32*a41 + a22*a31*a43 + a21*a33*a42
	cofactor.M[3][1] = a11*a32*a43 + a12*a33*a41 + a13*a31*a42 - a13*a32*a41 - a12*a31*a43 - a11*a33*a42
	cofactor.M[3][2] = -a11*a22*a43 - a12*a23*a41 - a13*a21*a42 + a13*a22*a41 + a12*a21*a43 + a11*a23*a42
	cofactor.M[3][3] = a11*a22*a33 + a12*a23*a31 + a13*a21*a32 - a13*a22*a31 - a12*a21*a33 - a11*a23*a32

	det := MGetDet(m)
	rec_det := 1.0 / det

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret.M[i][j] = cofactor.M[i][j] * rec_det
		}
	}

	return ret
}

func innerVTransform(ex_v []float32, m Matrix) dh_vector.Vector {
	var ret dh_vector.Vector

	m00 := m.M[0][0]
	m01 := m.M[0][1]
	m02 := m.M[0][2]
	m03 := m.M[0][3]
	m10 := m.M[1][0]
	m11 := m.M[1][1]
	m12 := m.M[1][2]
	m13 := m.M[1][3]
	m20 := m.M[2][0]
	m21 := m.M[2][1]
	m22 := m.M[2][2]
	m23 := m.M[2][3]

	ret.X = m00*ex_v[0] + m01*ex_v[1] + m02*ex_v[2] + m03*ex_v[3]
	ret.Y = m10*ex_v[0] + m11*ex_v[1] + m12*ex_v[2] + m13*ex_v[3]
	ret.Z = m20*ex_v[0] + m21*ex_v[1] + m22*ex_v[2] + m23*ex_v[3]

	return ret
}
func VTransform(v dh_vector.Vector, m Matrix) dh_vector.Vector {
	var ex_v [4]float32
	ex_v[0] = v.X
	ex_v[1] = v.Y
	ex_v[2] = v.Z
	ex_v[3] = 1.0

	ret := innerVTransform(ex_v[:], m)
	return ret
}
func VTransformSR(v dh_vector.Vector, m Matrix) dh_vector.Vector {
	var ex_v [4]float32
	ex_v[0] = v.X
	ex_v[1] = v.Y
	ex_v[2] = v.Z
	ex_v[3] = 0.0

	ret := innerVTransform(ex_v[:], m)
	return ret
}
