package mathg

import "math"

/*
m11 m12
m21 m22
*/
type Mat2 struct {
	M11 float64
	M21 float64
	M12 float64
	M22 float64
}

func (m *Mat2) Zero() *Mat2 {
	return &Mat2{0., 0., 0., 0.}
}

func (m *Mat2) Identity() *Mat2 {
	return &Mat2{1., 0., 0., 1.}
}

func (m *Mat2) Determinant() float64 {
	return m.M11*m.M22 - m.M12*m.M21
}

func (m *Mat2) Negative() *Mat2 {
	return &Mat2{-m.M11, -m.M21, -m.M12, -m.M22}
}

func (m *Mat2) Transpose() *Mat2 {
	return &Mat2{m.M11, m.M12, m.M21, m.M22}
}

func (m *Mat2) Cofactor() *Mat2 {
	return &Mat2{m.M22, -m.M12, -m.M21, m.M11}
}

func (m *Mat2) Adjugate() *Mat2 {
	return &Mat2{m.M22, -m.M21, -m.M12, m.M11}
}

func (m *Mat2) Multiply(m1 *Mat2) *Mat2 {
	return &Mat2{
		m.M11*m1.M11 + m.M12*m1.M21,
		m.M21*m1.M11 + m.M22*m1.M21,
		m.M11*m1.M12 + m.M12*m1.M22,
		m.M21*m1.M12 + m.M22*m1.M22,
	}
}

func (m *Mat2) MultiplyScalar(scalar float64) *Mat2 {
	return &Mat2{m.M11 * scalar, m.M21 * scalar, m.M12 * scalar, m.M22 * scalar}
}

func (m *Mat2) Inverse() *Mat2 {
	det := m.Determinant()
	inverse := m.Cofactor()
	inverse = inverse.MultiplyScalar(1. / det)
	return inverse
}

func (m *Mat2) Scaling(v *Vec2) *Mat2 {
	return &Mat2{v.X, m.M21, m.M12, v.Y}
}

func (m *Mat2) Scale(v *Vec2) *Mat2 {
	return &Mat2{m.M11 * v.X, m.M21, m.M12, m.M22 * v.Y}
}

func (m *Mat2) RotationZ(angle float64) *Mat2 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return &Mat2{c, s, -s, c}
}

func (m *Mat2) Lerp(m1 *Mat2, percent float64) *Mat2 {
	return &Mat2{
		m.M11 + (m1.M11-m.M11)*percent,
		m.M21 + (m1.M21-m.M21)*percent,
		m.M12 + (m1.M12-m.M12)*percent,
		m.M22 + (m1.M22-m.M22)*percent,
	}
}
