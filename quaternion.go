package mathg

import "math"

type Quaternion Vec4

func (q *Quaternion) IsZero() bool {
	return math.Abs(q.X) < epsilon && math.Abs(q.Y) < epsilon && math.Abs(q.Z) < epsilon && math.Abs(q.W) < epsilon
}

func (q *Quaternion) IsEqual(q1 *Quaternion) bool {
	return math.Abs(q.X-q1.X) < epsilon && math.Abs(q.Y-q1.Y) < epsilon && math.Abs(q.Z-q1.Z) < epsilon && math.Abs(q.W-q1.W) < epsilon
}

func (q *Quaternion) Zero() {
	q.X = 0.
	q.Y = 0.
	q.Z = 0.
	q.W = 0.
}

func (q *Quaternion) Null() {
	q.X = 0.
	q.Y = 0.
	q.Z = 0.
	q.W = 1.
}

func (q *Quaternion) Multiply(q1 *Quaternion) *Quaternion {
	return &Quaternion{
		q.W*q1.X + q.X*q1.W + q.Y*q1.Z - q.Z*q1.Y,
		q.W*q1.Y + q.Y*q1.W + q.Z*q1.X - q.X*q1.Z,
		q.W*q1.Z + q.Z*q1.W + q.X*q1.Y - q.Y*q1.X,
		q.W*q1.W - q.X*q1.X - q.Y*q1.Y - q.Z*q1.Z,
	}
}

func (q *Quaternion) MultiplyScalar(scalar float64) *Quaternion {
	return &Quaternion{q.X * scalar, q.Y * scalar, q.Z * scalar, q.W * scalar}
}

func (q *Quaternion) Divide(q1 *Quaternion) *Quaternion {
	d := math.Pow(q1.X, 2) + math.Pow(q1.Y, 2) + math.Pow(q1.Z, 2) + math.Pow(q1.W, 2)
	return &Quaternion{
		(q1.X*q.X + q1.Y*q.Y + q1.Z*q.Z + q1.W*q.W) / d,
		(q1.X*q.Y + q1.Y*q.X + q1.Z*q.W + q1.W*q.Z) / d,
		(q1.X*q.Z + q1.Y*q.W + q1.Z*q.X + q1.W*q.Y) / d,
		(q1.X*q.W + q1.Y*q.Z + q1.Z*q.Y + q1.W*q.X) / d,
	}
}

func (q *Quaternion) DivideScalar(scalar float64) *Quaternion {
	return &Quaternion{q.X / scalar, q.Y / scalar, q.Z / scalar, q.W / scalar}
}

func (q *Quaternion) Magnitude() float64 {
	return math.Sqrt(math.Pow(q.X, 2) + math.Pow(q.Y, 2) + math.Pow(q.Z, 2) + math.Pow(q.W, 2))
}

func (q *Quaternion) LengthSquared() float64 {
	return math.Pow(q.X, 2) + math.Pow(q.Y, 2) + math.Pow(q.Z, 2) + math.Pow(q.W, 2)
}

func (q *Quaternion) Negative() *Quaternion {
	return &Quaternion{-1 * q.X, -1 * q.Y, -1 * q.Z, -1 * q.W}
}

func (q *Quaternion) Conjugate() *Quaternion {
	return &Quaternion{-1 * q.X, -1 * q.Y, -1 * q.Z, q.W}
}

func (q *Quaternion) Inverse() *Quaternion {
	l := 1. / (math.Pow(q.X, 2) + math.Pow(q.Y, 2) + math.Pow(q.Z, 2) + math.Pow(q.W, 2))
	c := q.Conjugate()
	return &Quaternion{c.X * l, c.Y * l, c.Z * l, c.W * l}
}

func (q *Quaternion) Normalize() *Quaternion {
	l := 1. / q.Magnitude()
	return &Quaternion{q.X * l, q.Y * l, q.Z * l, q.W * l}
}

func (q *Quaternion) Dot(q1 *Quaternion) float64 {
	return q.X*q1.X + q.Y*q1.Y + q.Z*q1.Z + q.W + q1.W
}

func (q *Quaternion) Power(exponent float64) *Quaternion {
	if math.Abs(q.W) < 1.-epsilon {
		alpha := math.Acos(q.W)
		new_alpha := alpha * exponent
		s := math.Sin(new_alpha) / math.Sin(alpha)
		return &Quaternion{q.X * s, q.Y * s, q.Z * s, math.Cos(q.W)}
	} else {
		return q
	}
}

func (v *Vec4) ToQuaternionFromAxisAngle(angle float64) *Quaternion {
	half := angle * 0.5
	s := math.Sin(half)
	return &Quaternion{v.X * s, v.Y * s, v.Z * s, math.Cos(half)}
}

func (v *Vec3) ToQuaternion(v1 *Vec3) *Quaternion {
	d := v.Dot(v1)
	als := v.LengthSquared()
	bls := v1.LengthSquared()
	c := v.Cross(v1)
	q := &Quaternion{c.X, c.Y, c.Z, d + math.Sqrt(als*bls)}
	return q.Normalize()
}

func (m *Mat4) ToQuaternion() *Quaternion {
	scale := m.M11 + m.M22 + m.M33
	if scale > 0. {
		sqrt := math.Sqrt(scale + 1.)
		half := 0.5 / sqrt
		return &Quaternion{
			(m.M23 - m.M32) * half,
			(m.M31 - m.M13) * half,
			(m.M12 - m.M21) * half,
			sqrt * 0.5,
		}
	} else if (m.M11 >= m.M22) && (m.M11 >= m.M33) {
		sqrt := math.Sqrt(1. + m.M11 - m.M22 - m.M33)
		half := 0.5 / sqrt
		return &Quaternion{
			0.5 * sqrt,
			(m.M12 + m.M21) * half,
			(m.M13 + m.M31) * half,
			(m.M23 - m.M32) * half,
		}
	} else if m.M22 > m.M33 {
		sqrt := math.Sqrt(1. + m.M22 - m.M11 - m.M33)
		half := 0.5 / sqrt
		return &Quaternion{
			(m.M21 + m.M12) * half,
			0.5 * sqrt,
			(m.M32 + m.M23) * half,
			(m.M31 - m.M13) * half,
		}
	} else {
		sqrt := math.Sqrt(1. + m.M33 - m.M11 - m.M22)
		half := 0.5 / sqrt
		return &Quaternion{
			(m.M31 + m.M13) * half,
			(m.M32 + m.M23) * half,
			(0.5 * sqrt),
			(m.M12 - m.M21) * half,
		}
	}
}

func (q *Quaternion) Lerp(q1 *Quaternion, percent float64) *Quaternion {
	return &Quaternion{
		q.X + (q1.X-q.X)*percent,
		q.Y + (q1.Y-q.Y)*percent,
		q.Z + (q1.Z-q.Z)*percent,
		q.W + (q1.W-q.W)*percent,
	}
}

func (q *Quaternion) Slerp(q1 *Quaternion, percent float64) *Quaternion {
	var tmp *Quaternion = q1
	var f0, f1 float64
	d := q.Dot(q1)
	if d < 0. {
		tmp = q1.Negative()
		d = -d
	}
	if d > 0.9995 {
		f0 = 1.0 - percent
		f1 = percent
	} else {
		theta := math.Acos(d)
		sin_theta := math.Sin(theta)
		f0 = math.Sin((1.0-percent)*theta) / sin_theta
		f1 = math.Sin(percent*theta) / sin_theta
	}
	return &Quaternion{
		q.X*f0 + tmp.X*f1,
		q.Y*f0 + tmp.Y*f1,
		q.Z*f0 + tmp.Z*f1,
		q.W*f0 + tmp.W*f1,
	}
}

func (q *Quaternion) Angle(q1 *Quaternion) float64 {
	s := 1. / math.Sqrt(q.LengthSquared()*q1.LengthSquared())
	return math.Acos(q.Dot(q1) * s)
}
