package mathg

import "math"

/*
m11 m12 m13 m14
m21 m22 m23 m24
m31 m32 m33 m34
m41 m42 m43 m44
*/
type Mat4 struct {
	M11 float64
	M21 float64
	M31 float64
	M41 float64
	M12 float64
	M22 float64
	M32 float64
	M42 float64
	M13 float64
	M23 float64
	M33 float64
	M43 float64
	M14 float64
	M24 float64
	M34 float64
	M44 float64
}

func (m *Mat4) Zero() *Mat4 {
	return &Mat4{
		0., 0., 0., 0.,
		0., 0., 0., 0.,
		0., 0., 0., 0.,
		0., 0., 0., 0.,
	}
}

func (m *Mat4) Identity() *Mat4 {
	return &Mat4{
		1., 0., 0., 0.,
		0., 1., 0., 0.,
		0., 0., 1., 0.,
		0., 0., 0., 1.,
	}
}

func (m *Mat4) Clone() *Mat4 {
	return &Mat4{
		m.M11, m.M21, m.M31, m.M41,
		m.M12, m.M22, m.M32, m.M42,
		m.M13, m.M23, m.M33, m.M43,
		m.M14, m.M24, m.M34, m.M44,
	}
}

func (m *Mat4) Determinant() float64 {
	return m.M14*m.M23*m.M32*m.M41 -
		m.M13*m.M24*m.M32*m.M41 -
		m.M14*m.M22*m.M33*m.M41 +
		m.M12*m.M24*m.M33*m.M41 +
		m.M13*m.M22*m.M34*m.M41 -
		m.M12*m.M23*m.M34*m.M41 -
		m.M14*m.M23*m.M31*m.M42 +
		m.M13*m.M24*m.M31*m.M42 +
		m.M14*m.M21*m.M33*m.M42 -
		m.M11*m.M24*m.M33*m.M42 -
		m.M13*m.M21*m.M34*m.M42 +
		m.M11*m.M23*m.M34*m.M42 +
		m.M14*m.M22*m.M31*m.M43 -
		m.M12*m.M24*m.M31*m.M43 -
		m.M14*m.M21*m.M32*m.M43 +
		m.M11*m.M24*m.M32*m.M43 +
		m.M12*m.M21*m.M34*m.M43 -
		m.M11*m.M22*m.M34*m.M43 -
		m.M13*m.M22*m.M31*m.M44 +
		m.M12*m.M23*m.M31*m.M44 +
		m.M13*m.M21*m.M32*m.M44 -
		m.M11*m.M23*m.M32*m.M44 -
		m.M12*m.M21*m.M33*m.M44 +
		m.M11*m.M22*m.M33*m.M44
}

func (m *Mat4) Negative() *Mat4 {
	return &Mat4{
		-m.M11, -m.M21, -m.M31, -m.M41,
		-m.M12, -m.M22, -m.M32, -m.M42,
		-m.M13, -m.M23, -m.M33, -m.M43,
		-m.M14, -m.M24, -m.M34, -m.M44,
	}
}

func (m *Mat4) Transpose() *Mat4 {
	return &Mat4{
		m.M11, m.M12, m.M13, m.M14,
		m.M21, m.M22, m.M23, m.M24,
		m.M31, m.M32, m.M33, m.M34,
		m.M41, m.M42, m.M43, m.M44,
	}
}

func (m *Mat4) Cofactor() *Mat4 {
	cofactor := &Mat4{}
	minor := &Mat3{
		m.M22, m.M32, m.M42,
		m.M23, m.M33, m.M43,
		m.M24, m.M34, m.M44,
	}
	cofactor.M11 = minor.Determinant()
	minor = &Mat3{
		m.M12, m.M32, m.M42,
		m.M13, m.M33, m.M43,
		m.M14, m.M34, m.M44,
	}
	cofactor.M21 = -minor.Determinant()
	minor = &Mat3{
		m.M12, m.M22, m.M42,
		m.M13, m.M23, m.M43,
		m.M14, m.M24, m.M44,
	}
	cofactor.M31 = minor.Determinant()
	minor = &Mat3{
		m.M12, m.M22, m.M32,
		m.M13, m.M23, m.M33,
		m.M14, m.M24, m.M34,
	}
	cofactor.M41 = -minor.Determinant()
	minor = &Mat3{
		m.M21, m.M31, m.M41,
		m.M23, m.M33, m.M43,
		m.M24, m.M34, m.M44,
	}
	cofactor.M12 = -minor.Determinant()
	minor = &Mat3{
		m.M11, m.M31, m.M41,
		m.M13, m.M33, m.M43,
		m.M14, m.M34, m.M44,
	}
	cofactor.M22 = minor.Determinant()
	minor = &Mat3{
		m.M11, m.M21, m.M41,
		m.M13, m.M23, m.M43,
		m.M14, m.M24, m.M44,
	}
	cofactor.M32 = -minor.Determinant()
	minor = &Mat3{
		m.M11, m.M21, m.M31,
		m.M13, m.M23, m.M33,
		m.M14, m.M24, m.M34,
	}
	cofactor.M42 = minor.Determinant()
	minor = &Mat3{
		m.M21, m.M31, m.M41,
		m.M22, m.M32, m.M42,
		m.M24, m.M34, m.M44,
	}
	cofactor.M13 = minor.Determinant()
	minor = &Mat3{
		m.M11, m.M31, m.M41,
		m.M12, m.M32, m.M42,
		m.M14, m.M34, m.M44,
	}
	cofactor.M23 = -minor.Determinant()
	minor = &Mat3{
		m.M11, m.M21, m.M41,
		m.M12, m.M22, m.M42,
		m.M14, m.M24, m.M44,
	}
	cofactor.M33 = minor.Determinant()
	minor = &Mat3{
		m.M11, m.M21, m.M31,
		m.M12, m.M22, m.M32,
		m.M14, m.M24, m.M34,
	}
	cofactor.M43 = -minor.Determinant()
	minor = &Mat3{
		m.M21, m.M31, m.M41,
		m.M22, m.M32, m.M42,
		m.M23, m.M33, m.M43,
	}
	cofactor.M14 = -minor.Determinant()
	minor = &Mat3{
		m.M11, m.M31, m.M41,
		m.M12, m.M32, m.M42,
		m.M13, m.M33, m.M43,
	}
	cofactor.M24 = minor.Determinant()
	minor = &Mat3{
		m.M11, m.M21, m.M41,
		m.M12, m.M22, m.M42,
		m.M13, m.M23, m.M43,
	}
	cofactor.M34 = -minor.Determinant()
	minor = &Mat3{
		m.M11, m.M21, m.M31,
		m.M12, m.M22, m.M32,
		m.M13, m.M23, m.M33,
	}
	cofactor.M44 = minor.Determinant()
	return cofactor
}

func (m *Mat4) RotationX(angle float64) *Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return &Mat4{
		m.M11, m.M21, m.M31, m.M41,
		m.M12, c, s, m.M42,
		m.M13, -s, c, m.M43,
		m.M14, m.M24, m.M34, m.M44,
	}
}

func (m *Mat4) RotationY(angle float64) *Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return &Mat4{
		c, m.M21, -s, m.M41,
		m.M12, m.M22, m.M32, m.M42,
		s, m.M23, c, m.M43,
		m.M14, m.M24, m.M34, m.M44,
	}
}

func (m *Mat4) RotationZ(angle float64) *Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return &Mat4{
		c, s, m.M31, m.M41,
		-s, c, m.M32, m.M42,
		m.M13, m.M23, m.M33, m.M43,
		m.M14, m.M24, m.M34, m.M44,
	}
}

func (v *Vec3) RotationAxis(angle float64) *Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	one_c := 1. - c
	x, y, z := v.X, v.Y, v.Z
	xx := x * x
	xy := x * y
	xz := x * z
	yy := y * y
	yz := y * z
	zz := z * z
	l := xx + yy + zz
	sqrt := math.Sqrt(l)
	return &Mat4{
		(xx + (yy+zz)*c) / l,
		(xy*one_c + z*sqrt*s) / l,
		(xz*one_c - y*sqrt*s) / l,
		0.,
		(xy*one_c - z*sqrt*s) / l,
		(yy + (xx+zz)*c) / l,
		(yz*one_c + x*sqrt*s) / l,
		0.,
		(xz*one_c + y*sqrt*s) / l,
		(yz*one_c - x*sqrt*s) / l,
		(zz + (xx+yy)*c) / l,
		0.,
		0.,
		0.,
		0.,
		1.,
	}
}

func (q *Quaternion) RotationMatrix4() *Mat4 {
	xx := q.X * q.X
	yy := q.Y * q.Y
	zz := q.Z * q.Z
	xy := q.X * q.Y
	xz := q.X * q.Z
	xw := q.W * q.W
	yz := q.Y * q.Z
	yw := q.Y * q.W
	zw := q.Z * q.W
	return &Mat4{
		1. - 2.*(yy+zz),
		2. * (xy + zw),
		2. * (xz - yw),
		0.,
		2. * (xy - zw),
		1. - 2.*(xx+zz),
		2. * (yz + xw),
		0.,
		2. * (xz + yw),
		2. * (yz - xw),
		1. - 2.*(xx+yy),
		0.,
		0.,
		0.,
		0.,
		1.,
	}
}

func (m *Mat4) Translation(v *Vec3) *Mat4 {
	return &Mat4{
		m.M11, m.M21, m.M31, m.M41,
		m.M12, m.M22, m.M32, m.M42,
		m.M13, m.M23, m.M33, m.M43,
		v.X, v.Y, v.Z, m.M44,
	}
}

func (m *Mat4) Translate(v *Vec3) *Mat4 {
	return &Mat4{
		m.M11, m.M21, m.M31, m.M41,
		m.M12, m.M22, m.M32, m.M42,
		m.M13, m.M23, m.M33, m.M43,
		m.M14 + v.X, m.M24 + v.Y, m.M34 + v.Z, m.M44,
	}
}

func (m *Mat4) Scaling(v *Vec3) *Mat4 {
	return &Mat4{
		v.X, m.M21, m.M31, m.M41,
		m.M12, v.Y, m.M32, m.M42,
		m.M13, m.M23, v.Z, m.M43,
		m.M14, m.M24, m.M34, m.M44,
	}
}

func (m *Mat4) Scale(v *Vec3) *Mat4 {
	return &Mat4{
		m.M11 * v.X, m.M21, m.M31, m.M41,
		m.M12, m.M22 * v.Y, m.M32, m.M42,
		m.M13, m.M23, m.M33 * v.Z, m.M43,
		m.M14, m.M24, m.M34, m.M44,
	}
}

func (m *Mat4) Multiply(m1 *Mat4) *Mat4 {
	return &Mat4{
		m.M11*m1.M11 + m.M12*m1.M21 + m.M13*m1.M31 + m.M14*m.M41,
		m.M21*m1.M11 + m.M22*m1.M21 + m.M23*m1.M31 + m.M24*m.M41,
		m.M31*m1.M11 + m.M32*m1.M21 + m.M33*m1.M31 + m.M34*m.M41,
		m.M41*m1.M11 + m.M42*m1.M21 + m.M43*m1.M31 + m.M44*m.M41,
		m.M11*m1.M12 + m.M12*m1.M22 + m.M13*m1.M32 + m.M14*m.M42,
		m.M21*m1.M12 + m.M22*m1.M22 + m.M23*m1.M32 + m.M24*m.M42,
		m.M31*m1.M12 + m.M32*m1.M22 + m.M33*m1.M32 + m.M34*m.M42,
		m.M41*m1.M12 + m.M42*m1.M22 + m.M43*m1.M32 + m.M44*m.M42,
		m.M11*m1.M13 + m.M12*m1.M23 + m.M13*m1.M33 + m.M14*m.M43,
		m.M21*m1.M13 + m.M22*m1.M23 + m.M23*m1.M33 + m.M24*m.M43,
		m.M31*m1.M13 + m.M32*m1.M23 + m.M33*m1.M33 + m.M34*m.M43,
		m.M41*m1.M13 + m.M42*m1.M23 + m.M43*m1.M33 + m.M44*m.M43,
		m.M11*m1.M14 + m.M12*m1.M24 + m.M13*m1.M34 + m.M14*m.M44,
		m.M21*m1.M14 + m.M22*m1.M24 + m.M23*m1.M34 + m.M24*m.M44,
		m.M31*m1.M14 + m.M32*m1.M24 + m.M33*m1.M34 + m.M34*m.M44,
		m.M41*m1.M14 + m.M42*m1.M24 + m.M43*m1.M34 + m.M44*m.M44,
	}
}

func (m *Mat4) MultiplyScalar(scalar float64) *Mat4 {
	return &Mat4{
		m.M11 * scalar, m.M21 * scalar, m.M31 * scalar, m.M41 * scalar,
		m.M12 * scalar, m.M22 * scalar, m.M32 * scalar, m.M42 * scalar,
		m.M13 * scalar, m.M23 * scalar, m.M33 * scalar, m.M43 * scalar,
		m.M14 * scalar, m.M24 * scalar, m.M34 * scalar, m.M44 * scalar,
	}
}

func (m *Mat4) Inverse() *Mat4 {
	cof := m.Cofactor()
	adj := cof.Transpose()
	d := m.Determinant()
	return adj.MultiplyScalar(1. / d)
}

func (m *Mat4) Lerp(m1 *Mat4, percent float64) *Mat4 {
	return &Mat4{
		m.M11 + (m1.M11-m.M11)*percent,
		m.M21 + (m1.M21-m.M21)*percent,
		m.M31 + (m1.M31-m.M31)*percent,
		m.M41 + (m1.M41-m.M41)*percent,
		m.M12 + (m1.M12-m.M12)*percent,
		m.M22 + (m1.M22-m.M22)*percent,
		m.M32 + (m1.M32-m.M32)*percent,
		m.M42 + (m1.M42-m.M42)*percent,
		m.M13 + (m1.M13-m.M13)*percent,
		m.M23 + (m1.M23-m.M23)*percent,
		m.M33 + (m1.M33-m.M33)*percent,
		m.M43 + (m1.M43-m.M43)*percent,
		m.M14 + (m1.M14-m.M14)*percent,
		m.M24 + (m1.M24-m.M24)*percent,
		m.M34 + (m1.M34-m.M34)*percent,
		m.M44 + (m1.M44-m.M44)*percent,
	}
}

func (pos *Vec3) LookAt(target *Vec3, up *Vec3) *Mat4 {
	fwrd := target.Subtract(pos)
	fwrd = fwrd.Normalize()
	side := fwrd.Cross(up)
	side = side.Normalize()
	newUp := side.Cross(fwrd)
	return &Mat4{
		side.X, newUp.X, -fwrd.X, 0.,
		side.Y, newUp.Y, -fwrd.Y, 0.,
		side.Z, newUp.Z, -fwrd.Z, 0.,
		-side.Dot(pos), -newUp.Dot(pos), fwrd.Dot(pos), 1.,
	}
}

/*
left, right, bottom, top, near, far
*/
func Ortho(l, r, b, t, n, f float64) *Mat4 {
	return &Mat4{
		2. / (r - l), 0., 0., 0.,
		0., 2. / (t - b), 0., 0.,
		0., 0., -2. / (f - n), 0.,
		-((r + l) / (r - l)), -((t + b) / (t - b)), -((f + n) / (f - n)), 1.,
	}
}

func Perspective(fov_y, aspect, near, far float64) *Mat4 {
	tan_half_fov_y := 1. / math.Tan(fov_y*0.5)
	return &Mat4{
		1. / aspect * tan_half_fov_y, 0., 0., 0.,
		0., 1. / tan_half_fov_y, 0., 0.,
		0., 0., far / (near - far), -1.,
		0., 0., -(far * near) / (far - near), 0.,
	}
}

func PerspectiveFOV(fov, w, h, n, f float64) *Mat4 {
	h2 := math.Cos(fov*0.5) / math.Sin(fov*0.5)
	w2 := h2 * h / w
	return &Mat4{
		w2, 0., 0., 0.,
		0., h2, 0., 0.,
		0., 0., f / (n - f), 1.,
		0., 0., -(f * n) / (f - n), 0.,
	}
}

func PerspectiveInfinite(fov_y, aspect, near float64) *Mat4 {
	rng := math.Tan(fov_y*0.5) * near
	left := -rng * aspect
	right := rng * aspect
	top := rng
	bottom := -rng
	return &Mat4{
		2. * near / (right - left), 0., 0., 0.,
		0., 2. * near / (top - bottom), 0., 0.,
		0., 0., -1., -1.,
		0., 0., -2. * near, 0.,
	}
}
