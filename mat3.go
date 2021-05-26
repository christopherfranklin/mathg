package mathg

import "math"

/*
m11 m12 m13
m21 m22 m23
m31 m32 m33
*/
type Mat3 struct {
	M11 float64
	M21 float64
	M31 float64
	M12 float64
	M22 float64
	M32 float64
	M13 float64
	M23 float64
	M33 float64
}

func (m *Mat3) Clone() *Mat3 {
	return &Mat3{
		m.M11,
		m.M21,
		m.M31,
		m.M12,
		m.M22,
		m.M32,
		m.M13,
		m.M23,
		m.M33,
	}
}

func (m *Mat3) Zero() *Mat3 {
	return &Mat3{0., 0., 0., 0., 0., 0., 0., 0., 0.}
}

func (m *Mat3) Identity() *Mat3 {
	return &Mat3{1., 0., 0., 0., 1., 0., 0., 0., 1.}
}

func (m *Mat3) Determinant() float64 {
	return m.M11*m.M22*m.M33 +
		m.M12*m.M23*m.M31 +
		m.M13*m.M21*m.M32 -
		m.M11*m.M23*m.M32 -
		m.M12*m.M21*m.M33 -
		m.M13*m.M22*m.M31
}

func (m *Mat3) Negative() *Mat3 {
	return &Mat3{
		-m.M11,
		-m.M21,
		-m.M31,
		-m.M12,
		-m.M22,
		-m.M32,
		-m.M13,
		-m.M23,
		-m.M33,
	}
}

func (m *Mat3) Transpose() *Mat3 {
	return &Mat3{
		m.M11,
		m.M12,
		m.M13,
		m.M21,
		m.M22,
		m.M23,
		m.M31,
		m.M32,
		m.M33,
	}
}

func (m *Mat3) Cofactor() *Mat3 {
	cofactor := &Mat3{}
	minor := &Mat2{m.M22, m.M32, m.M23, m.M33}
	cofactor.M11 = minor.Determinant()
	minor = &Mat2{m.M12, m.M32, m.M13, m.M33}
	cofactor.M21 = -minor.Determinant()
	minor = &Mat2{m.M12, m.M22, m.M13, m.M23}
	cofactor.M31 = minor.Determinant()
	minor = &Mat2{m.M21, m.M31, m.M23, m.M33}
	cofactor.M12 = -minor.Determinant()
	minor = &Mat2{m.M11, m.M31, m.M13, m.M33}
	cofactor.M22 = minor.Determinant()
	minor = &Mat2{m.M11, m.M21, m.M13, m.M23}
	cofactor.M32 = -minor.Determinant()
	minor = &Mat2{m.M21, m.M31, m.M22, m.M32}
	cofactor.M13 = minor.Determinant()
	minor = &Mat2{m.M11, m.M31, m.M12, m.M32}
	cofactor.M23 = -minor.Determinant()
	minor = &Mat2{m.M11, m.M21, m.M12, m.M22}
	cofactor.M33 = minor.Determinant()
	return cofactor
}

func (m *Mat3) Multiply(m1 *Mat3) *Mat3 {
	return &Mat3{
		m.M11*m1.M11 + m.M12*m1.M21 + m.M13*m1.M31,
		m.M21*m1.M11 + m.M22*m1.M21 + m.M23*m1.M31,
		m.M31*m1.M11 + m.M32*m1.M21 + m.M33*m1.M31,
		m.M11*m1.M12 + m.M12*m1.M22 + m.M13*m1.M32,
		m.M21*m1.M12 + m.M22*m1.M22 + m.M23*m1.M32,
		m.M31*m1.M12 + m.M32*m1.M22 + m.M33*m1.M32,
		m.M11*m1.M13 + m.M12*m1.M23 + m.M13*m1.M33,
		m.M21*m1.M13 + m.M22*m1.M23 + m.M23*m1.M33,
		m.M31*m1.M13 + m.M32*m1.M23 + m.M33*m1.M33,
	}
}

func (m *Mat3) MultiplyScalar(scalar float64) *Mat3 {
	return &Mat3{
		m.M11 * scalar,
		m.M21 * scalar,
		m.M31 * scalar,
		m.M12 * scalar,
		m.M22 * scalar,
		m.M32 * scalar,
		m.M13 * scalar,
		m.M23 * scalar,
		m.M33 * scalar,
	}
}

func (m *Mat3) Inverse() *Mat3 {
	cof := m.Cofactor()
	adj := cof.Transpose()
	d := m.Determinant()
	return adj.MultiplyScalar(1. / d)
}

func (m *Mat3) Scaling(v *Vec3) *Mat3 {
	return &Mat3{
		v.X,
		m.M21,
		m.M31,
		m.M12,
		v.Y,
		m.M32,
		m.M13,
		m.M23,
		v.Z,
	}
}

func (m *Mat3) Scale(v *Vec3) *Mat3 {
	return &Mat3{
		m.M11 * v.X,
		m.M21,
		m.M31,
		m.M12,
		m.M22 * v.Y,
		m.M32,
		m.M13,
		m.M23,
		m.M33 * v.Z,
	}
}

func (m *Mat3) RotationX(angle float64) *Mat3 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	rx := m.Clone()
	rx.M22 = c
	rx.M32 = s
	rx.M23 = -s
	rx.M33 = c
	return rx
}

func (m *Mat3) RotationY(angle float64) *Mat3 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	rx := m.Clone()
	rx.M11 = c
	rx.M31 = -s
	rx.M13 = s
	rx.M33 = c
	return rx
}

func (m *Mat3) RotationZ(angle float64) *Mat3 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	rx := m.Clone()
	rx.M11 = c
	rx.M21 = s
	rx.M12 = -s
	rx.M22 = c
	return rx
}

func (m *Mat3) RotateAxis(angle float64) *Mat3 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	one_c := 1. - c
	x := m.M11
	y := m.M22
	z := m.M33
	xx := x * x
	xy := x * y
	xz := x * z
	yy := y * y
	yz := y * z
	zz := z * z
	l := xx + yy + zz
	sqrt := math.Sqrt(l)
	return &Mat3{
		(xx + (yy+zz)*c) / l,
		(xy*one_c + m.M31*sqrt*s) / l,
		(xz*one_c - m.M21*sqrt*s) / l,
		(xy + one_c - m.M31*sqrt*s) / l,
		(yy + (xx+zz)*c) / l,
		(yz + one_c + m.M11*sqrt*s) / l,
		(xz*one_c + m.M21*sqrt*s) / l,
		(yz*one_c - m.M11*sqrt*s) / l,
		(zz + (xx+yy)*c) / l,
	}
}

func (q *Quaternion) RotationMatrix() *Mat3 {
	xx := q.X * q.X
	yy := q.Y * q.Y
	zz := q.Z * q.Z
	xy := q.X * q.Y
	zw := q.Z * q.W
	xz := q.X * q.Z
	yw := q.Y * q.W
	yz := q.Y * q.Z
	xw := q.X * q.W
	return &Mat3{
		1. - 2.*(yy-zz),
		2. * (xy + zw),
		2. * (xz - yw),
		2. * (xy - zw),
		1. - 2.*(xx-zz),
		2. * (yz - xw),
		2. * (xz - yw),
		2. * (yz - xw),
		1. - 2.*(xx-yy),
	}
}

func (m *Mat3) Lerp(m1 *Mat3, percent float64) *Mat3 {
	return &Mat3{
		m.M11 + (m1.M11-m.M11)*percent,
		m.M21 + (m1.M21-m.M21)*percent,
		m.M31 + (m1.M31-m.M31)*percent,
		m.M12 + (m1.M12-m.M12)*percent,
		m.M22 + (m1.M22-m.M22)*percent,
		m.M32 + (m1.M32-m.M32)*percent,
		m.M13 + (m1.M13-m.M13)*percent,
		m.M23 + (m1.M23-m.M23)*percent,
		m.M33 + (m1.M33-m.M33)*percent,
	}
}
